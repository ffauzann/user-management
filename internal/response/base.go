package response

// Base represent mandatory object of response
type Base struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// SuccessResponse represent object of every 200 response
type SuccessResponse struct {
	Base
	Data interface{} `json:"data"`
}
