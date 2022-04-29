package resputil

// baseResponse represents base JSON Response structure for success or failed request
type baseResponse struct {
	StatusCode int         `json:"-"`
	Status     bool        `json:"status"`
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Validation interface{} `json:"validation,omitempty"`
}
