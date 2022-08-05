package model

type Response struct {
	Message string      `json:"message"`
	Error   error       `json:"error"`
	Data    interface{} `json:"data"`
}
