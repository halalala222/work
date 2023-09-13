package consts

import "net/http"

type ResponseCode uint8

const (
	UnknownError ResponseCode = iota
	Success
	ServerError
	Forbidden
	Unauthorized
	TaskOwnershipError
	RequestError
)

var ErrorCodeHTTPStatus = map[ResponseCode]int{
	UnknownError:       http.StatusBadRequest,
	Success:            http.StatusOK,
	ServerError:        http.StatusInternalServerError,
	Forbidden:          http.StatusForbidden,
	Unauthorized:       http.StatusUnauthorized,
	TaskOwnershipError: http.StatusBadRequest,
	RequestError:       http.StatusBadRequest,
}
