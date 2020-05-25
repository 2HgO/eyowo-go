package eyowo

type route string

const (
	// BALANCE is the eyowo developer API route to retrieve an account's balance
	BALANCE route = `/v1/users/balance`
	// AUTHENTICATION is the eyowo developer API route to perform a user authentication flow
	AUTHENTICATION route = `/v1/users/auth`
	// PHONE_TRANSFER is the eyowo developer API route to perform an eyowo transfer
	PHONE_TRANSFER route = `/v1/users/transfers/phone`
	// BANK_TRANSFER is the eyowo developer API route to perform a bank transfer
	BANK_TRANSFER route = `/v1/users/transfers/bank`
	// VTU_PURCHASE is the eyowo developer API route to perform a Virtual Top-Up
	VTU_PURCHASE route = `/v1/users/payments/bills/vtu`
	// VALIDATION is the eyowo developer API route to validate a user
	VALIDATION route = `/v1/users/auth/validate`
	// REFRESH is the eyowo developer API route to refresh a user's access token
	REFRESH route = `/v1/users/accessToken`
)
