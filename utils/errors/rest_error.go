package errors

import (
	"net/http"
)

type RestErr struct {
	Message string `json: ""message`
	Status  int64  `json: "status"`
	Error   string `json: "error"`
}

func NewBadrequest(messsage string) *RestErr {
	return &RestErr{
		Message: messsage,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFounderror(messsage string) *RestErr {
	return &RestErr{
		Message: messsage,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(messsage string) *RestErr {
	return &RestErr{
		Message: messsage,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}