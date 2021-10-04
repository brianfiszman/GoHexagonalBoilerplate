package dtos

import "net/http"

const (
	STATUS_OK = http.StatusOK
	HELLO_MSG = "Hello World!"
)

type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
