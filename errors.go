package eyowo

import "errors"

var (
	NoAccessToken      = errors.New("No access token set")
	NoRefeshToken      = errors.New("No refresh token set")
	InvalidAppKey      = errors.New("Invalid App Key")
	InvalidAppSecret   = errors.New("Invalid App Secret")
	InvalidMobile      = errors.New("Invalid Mobile Number")
	InvalidEnvironment = errors.New("Invalid App Environment")
)
