package web

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"status"`
}
