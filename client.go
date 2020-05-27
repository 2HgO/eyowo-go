package eyowo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

var (
	defaultTimeout = time.Minute
)

type Client struct {
	// appKey is the application's app key
	appKey string
	// appSecret is the application's app secret
	appSecret string
	// accessToken is the access token used to authenticate requests by the client
	accessToken string
	// refreshToken is the refresh token used to refresh the client's access token
	refreshToken string
	// mobile is the mobile number for the client
	mobile string
	// environment is the environment the client is working in
	environment environment
	// httpClient is the underlying http client for the eyowo client
	httpClient *http.Client
	// lastRefresh specifies the time the client's access token was created
	lastRefresh *time.Time
	// expiresIn specifies the lifespan of the client's access token
	expiresIn time.Duration
}

// NewClient creates and returns a new eyowo API client
func NewClient(appKey, appSecret, mobile string, env environment) (*Client, error) {
	if strings.Trim(appKey, " ") == "" {
		return nil, InvalidAppKey
	}
	if strings.Trim(appSecret, " ") == "" {
		return nil, InvalidAppSecret
	}
	if env != SANDBOX && env != PRODUCTION {
		return nil, InvalidEnvironment
	}
	c := &Client{
		appKey:      appKey,
		appSecret:   appSecret,
		environment: env,
		mobile:      mobile,
		httpClient:  &http.Client{Timeout: defaultTimeout},
	}
	if res, err := c.ValidateUser(c.mobile); err != nil || !res.Success {
		return nil, InvalidMobile
	}
	return c, nil
}

// HasValidToken validates whether or not the client has a valid access token
func (c *Client) HasValidToken() bool {
	return c.lastRefresh != nil && c.lastRefresh.Add(c.expiresIn).After(time.Now())
}

// GetAccessToken returns the client's access token
func (c *Client) GetAccessToken() string {
	return c.accessToken
}

// GetMobile returns the client's mobile number
func (c *Client) GetMobile() string {
	return c.mobile
}

// GetRefreshToken returns the client's refresh token
func (c *Client) GetRefreshToken() string {
	return c.refreshToken
}

// SetAccessToken sets the access token for a client
func (c *Client) SetAccessToken(s string) {
	c.accessToken = s
}

// SetRefreshToken sets the refresh token for a client
func (c *Client) SetRefreshToken(s string) {
	c.refreshToken = s
}

// SetClientTimeout sets the timeout for requests by the client
// The default timeout value is 1 minute
func (c *Client) SetClientTimeout(t time.Duration) {
	c.httpClient.Timeout = t
}

// BuyVTU performs a Virtual Top-Up (VTU) for a mobile number
func (c *Client) BuyVTU(recipientMobileNumber string, amount uint, provider provider) (*Response, error) {
	payload := map[string]interface{}{
		"amount":   amount,
		"mobile":   recipientMobileNumber,
		"provider": provider,
	}
	return c.performRequest(payload, VTU_PURCHASE)
}

// GetBalance returns the account balance for an eyowo account
func (c *Client) GetBalance() (*Response, error) {
	payload := map[string]interface{}{
		"mobile": c.mobile,
	}
	return c.performRequest(payload, BALANCE)
}

// ValidateUser valdates whether or not a mobile number has an associated eyowo account
func (c *Client) ValidateUser(mobileNumber string) (*Response, error) {
	payload := map[string]interface{}{
		"mobile": c.mobile,
	}
	return c.performRequest(payload, VALIDATION)
}

// AuthenticateUser performs an authentication flow for a user
func (c *Client) AuthenticateUser(factor string, passcode ...string) (*Response, error) {
	payload := map[string]interface{}{
		"mobile": c.mobile,
		"factor": factor,
	}
	if len(passcode) != 0 {
		payload["passcode"] = passcode[0]
	}
	res, err := c.performRequest(payload, AUTHENTICATION)

	if len(passcode) > 0 {
		if err != nil {
			return nil, err
		}
		if _, ok := res.Data["accessToken"]; !ok {
			return nil, NoAccessToken
		}

		now := time.Now()
		expiresIn := res.Data["expiresIn"].(float64)

		c.accessToken = res.Data["accessToken"].(string)
		c.lastRefresh = &now
		c.expiresIn = time.Duration(expiresIn)
		c.refreshToken = res.Data["refreshToken"].(string)
	}
	return res, err
}

// TransferToBank transfers money from the client's user's account to the specified bank account
func (c *Client) TransferToBank(amount uint, accountName, accountNumber, bankCode string) (*Response, error) {
	payload := map[string]interface{}{
		"amount":        amount,
		"accountNumber": accountNumber,
		"accountName":   accountName,
		"bankCode":      bankCode,
	}
	return c.performRequest(payload, BANK_TRANSFER)
}

// TransferToBank transfers money from the client's user's account to the specified eyowo account
func (c *Client) TransferToPhone(amount uint, recipientMobileNumber string) (*Response, error) {
	payload := map[string]interface{}{
		"amount": amount,
		"mobile": recipientMobileNumber,
	}
	return c.performRequest(payload, PHONE_TRANSFER)
}

// RefreshAccessToken refreshes the client's access token using the refresh token
func (c *Client) RefreshAccessToken() error {
	if strings.Trim(c.refreshToken, " ") == "" {
		return NoRefeshToken
	}
	payload := map[string]interface{}{
		"refreshToken": c.refreshToken,
	}
	res, err := c.performRequest(payload, REFRESH)
	if err != nil {
		return err
	}

	if _, ok := res.Data["accessToken"]; !ok {
		return NoAccessToken
	}

	now := time.Now()
	expiresIn := res.Data["expiresIn"].(float64)

	c.accessToken = res.Data["accessToken"].(string)
	c.lastRefresh = &now
	c.expiresIn = time.Duration(expiresIn)
	return nil
}

// GetBanks fetches a list of nigerian banks and their corresponding bank code
func (c *Client) GetBanks() (*Response, error) {
	url := fmt.Sprintf("%s%s", c.environment, BANKS)

	ctx, cancel := context.WithTimeout(context.TODO(), c.httpClient.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	req.Header.Set(`Content-Type`, `application/json`)
	req.Header.Set(`X-App-Key`, c.appKey)
	req.Header.Set(`X-App-Wallet-Access-Token`, c.accessToken)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	eyowoRes := new(Response)
	err = json.NewDecoder(res.Body).Decode(eyowoRes)
	eyowoRes.Status = res.StatusCode
	return eyowoRes, err
}

// performRequest performs the http request to the eyowo developer environment for the client
func (c *Client) performRequest(payload map[string]interface{}, route route) (*Response, error) {
	url := fmt.Sprintf("%s%s", c.environment, route)

	data, _ := json.Marshal(payload)

	ctx, cancel := context.WithTimeout(context.TODO(), c.httpClient.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(data))
	req.Header.Set(`Content-Type`, `application/json`)
	req.Header.Set(`X-App-Key`, c.appKey)
	req.Header.Set(`X-App-Wallet-Access-Token`, c.accessToken)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	eyowoRes := new(Response)
	err = json.NewDecoder(res.Body).Decode(eyowoRes)
	eyowoRes.Status = res.StatusCode
	return eyowoRes, err
}
