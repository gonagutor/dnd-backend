package responses

// @Description Error code and readable error message
type FailureResponse struct {
	Error string `json:"error"`
	Message string `json:"message"`
}