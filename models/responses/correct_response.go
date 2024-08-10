package responses

type Pagination struct {
	Page int `json:"page" example:"1"`
	MaxPages int `json:"maxPage" example:"4"`
	PageSize int `json:"pageSize" example:"25"`
}

// @Description Action code and readable message. Data and Pagination is optional
type CorrectResponse struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Data Pagination `json:"data" swaggerignore:"true"`
	Pagination interface{} `json:"pagination" swaggerignore:"true"`
}