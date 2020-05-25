package eyowo

// Response specifies the structure of a response from the eyowo developer API
type Response struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Status  int                    `json:"status,omitempty"`
}
