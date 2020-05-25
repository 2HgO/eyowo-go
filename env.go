package eyowo

type environment string

const (
	// PRODUCTION is the eyowo developer API production environment
	PRODUCTION environment = `https://api.console.eyowo.com`
	// SANDBOX is the eyowo developer API test/sandbox environment
	SANDBOX environment = `https://api.sandbox.developer.eyowo.com`
)
