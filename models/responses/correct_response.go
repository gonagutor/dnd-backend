package responses

// @Description Action code and readable message. Data is optional
type CorrectResponse struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data" swaggerignore:"true"`
}