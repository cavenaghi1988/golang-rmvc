package util

// ApplicationError is return error json format
type ApplicationError struct {
	Meassage string `json:"message"`
	Status   int    `json:"status"`
	Code     string `json:"code"`
}
