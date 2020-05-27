package eyowo

import (
	// "log"
	"net/http"
	"testing"
	"time"

	. "github.com/stretchr/testify/assert"
)

var testClient, _ = NewClient("ru6nmdqf9cqpyvz7b4ce2kj938w5gc3r", "zvze3bfmev5pxhexuzsjcrn6pjqwbspgnh43de9nkvkjeeq45qemudmzyvpanv5k", "2348000000000", PRODUCTION)

// [TODO] clarify test environment endpoint
func init() {
	// res, err := testClient.AuthenticateUser("sms")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println(res)
	// res, err = testClient.AuthenticateUser("sms", "111111")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println(res)
	// log.Println(testClient.GetAccessToken())
}

func TestNewClient(t *testing.T) {
	type args struct {
		key    string
		secret string
		mobile string
		env    environment
	}
	testcases := []struct {
		name string
		args args
		err  error
		want *Client
	}{
		{
			name: "Create a new eyowo client",
			args: args{
				key:    "edaf5c2fcf4dbd978b54272595f95ad7",
				secret: "3184190bad4e981d3da9b741525d6a378124a3640bb20134d61ac804ce7bd56f",
				mobile: "2348000000000",
				env:    PRODUCTION,
			},
			want: &Client{
				appKey:      "edaf5c2fcf4dbd978b54272595f95ad7",
				appSecret:   "3184190bad4e981d3da9b741525d6a378124a3640bb20134d61ac804ce7bd56f",
				mobile:      "2348000000000",
				environment: PRODUCTION,
				httpClient:  &http.Client{Timeout: time.Minute},
			},
		},
		{
			name: "Fail to create eyowo client without app key",
			args: args{
				secret: "zvze3bfmev5pxhexuzsjcrn6pjqwbspgnh43de9nkvkjeeq45qemudmzyvpanv5k",
				mobile: "2348000000000",
				env:    PRODUCTION,
			},
			err: InvalidAppKey,
		},
		{
			name: "Fail to create eyowo client without app secret",
			args: args{
				key:    "ru6nmdqf9cqpyvz7b4ce2kj938w5gc3r",
				mobile: "2348000000000",
				env:    PRODUCTION,
			},
			err: InvalidAppSecret,
		},
		{
			name: "Fail to create eyowo client with invalid environment",
			args: args{
				key:    "ru6nmdqf9cqpyvz7b4ce2kj938w5gc3r",
				secret: "zvze3bfmev5pxhexuzsjcrn6pjqwbspgnh43de9nkvkjeeq45qemudmzyvpanv5k",
				mobile: "2348000000000",
				env:    "Bad Env",
			},
			err: InvalidEnvironment,
		},
		{
			name: "Fail to create eyowo client with invalid mobile number",
			args: args{
				key:    "edaf5c2fcf4dbd978b54272595f95ad7",
				secret: "3184190bad4e981d3da9b741525d6a378124a3640bb20134d61ac804ce7bd56f",
				mobile: "2349090000111",
				env:    PRODUCTION,
			},
			err: InvalidMobile,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			tc, err := NewClient(testcase.args.key, testcase.args.secret, testcase.args.mobile, testcase.args.env)

			Equal(t, testcase.err, err)
			Equal(t, testcase.want, tc)
		})
	}
}

func TestClientHasValidToken(t *testing.T) {
	type args struct {
		client *Client
	}
	testcases := []struct {
		name string
		args args
		want bool
	}{}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			res := testcase.args.client.HasValidToken()

			Equal(t, testcase.want, res)
		})
	}
}

func TestClientGetAccessToken(t *testing.T) {}

func TestClientGetRefreshToken(t *testing.T) {}

func TestClientSetAccessToken(t *testing.T) {
	client, err := NewClient("edaf5c2fcf4dbd978b54272595f95ad7", "3184190bad4e981d3da9b741525d6a378124a3640bb20134d61ac804ce7bd56f", "2348000000000", PRODUCTION)
	Nil(t, err)
	initialToken := client.GetAccessToken()
	client.SetAccessToken("new-access-token")
	NotEqual(t, initialToken, client.GetAccessToken())
}

func TestClientSetRefreshToken(t *testing.T) {
	client, err := NewClient("edaf5c2fcf4dbd978b54272595f95ad7", "3184190bad4e981d3da9b741525d6a378124a3640bb20134d61ac804ce7bd56f", "2348000000000", PRODUCTION)
	Nil(t, err)
	initialToken := client.GetRefreshToken()
	client.SetRefreshToken("new-refresh-token")
	NotEqual(t, initialToken, client.GetRefreshToken())
}

func TestClientSetClientTimeout(t *testing.T) {

}

func TestClientBuyVTU(t *testing.T) {
	type args struct {
		mobile   string
		amount   uint
		provider provider
	}
	testcases := []struct {
		name string
		args args
		err  error
		want *Response
	}{}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			res, err := testClient.BuyVTU(testcase.args.mobile, testcase.args.amount, testcase.args.provider)

			Equal(t, testcase.err, err)
			Equal(t, testcase.want, res)
		})
	}
}

func TestClientGetBalance(t *testing.T) {
	type args struct {
		mobile string
	}
	testcases := []struct {
		name string
		args args
		err  error
		want *Response
	}{}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			res, err := testClient.GetBalance()

			Equal(t, testcase.err, err)
			Equal(t, testcase.want, res)
		})
	}
}

func TestClientValidateUser(t *testing.T) {
	type args struct {
		mobile string
	}
	testcases := []struct {
		name string
		args args
		err  error
		want *Response
	}{}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			res, err := testClient.ValidateUser(testcase.args.mobile)

			Equal(t, testcase.err, err)
			Equal(t, testcase.want, res)
		})
	}
}

func TestClientAuthenticateUser(t *testing.T) {
	type args struct {
		mobile   string
		factor   string
		passcode []string
	}
	testcases := []struct {
		name string
		args args
		err  error
		want *Response
	}{}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			res, err := testClient.AuthenticateUser(testcase.args.factor, testcase.args.passcode...)

			Equal(t, testcase.err, err)
			Equal(t, testcase.want, res)
		})
	}
}

func TestClientTransferToBank(t *testing.T) {
	type args struct {
		amount        uint
		accountName   string
		accountNumber string
		bankCode      string
	}
	testcases := []struct {
		name string
		args args
		err  error
		want *Response
	}{}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			res, err := testClient.TransferToBank(testcase.args.amount, testcase.args.accountName, testcase.args.accountNumber, testcase.args.bankCode)

			Equal(t, testcase.err, err)
			Equal(t, testcase.want, res)
		})
	}
}

func TestClientTransferToPhone(t *testing.T) {
	type args struct {
		amount uint
		mobile string
	}
	testcases := []struct {
		name string
		args args
		err  error
		want *Response
	}{}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			res, err := testClient.TransferToPhone(testcase.args.amount, testcase.args.mobile)

			Equal(t, testcase.err, err)
			Equal(t, testcase.want, res)
		})
	}
}

func TestClientRefreshAccessToken(t *testing.T) {
	type args struct {
		client *Client
	}
	testcases := []struct {
		name    string
		args    args
		err     error
		wantErr bool
	}{}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			initialToken := testcase.args.client.GetAccessToken()
			err := testcase.args.client.RefreshAccessToken()

			Equal(t, testcase.err, err)
			if !testcase.wantErr {
				NotEqual(t, initialToken, testcase.args.client.GetAccessToken())
			}
		})
	}
}
