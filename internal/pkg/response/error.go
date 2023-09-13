package response

import (
	"errors"
	"work/internal/consts"
)

var _ error = &ErrorResult{}

type ErrorResult struct {
	Msg  string
	Code consts.ResponseCode
}

func (r *ErrorResult) Error() string {
	return r.Msg
}

func NewError(code consts.ResponseCode, msg string) error {
	if _, ok := consts.ErrorCodeHTTPStatus[code]; !ok {
		code = consts.UnknownError
	}
	if len(msg) == 0 {
		msg = UnKnownMsg
	}
	return &ErrorResult{
		Code: code,
		Msg:  msg,
	}
}

func GetErrorHTTPStatus(err error) int {
	var (
		e          *ErrorResult
		ok         bool
		httpStatus int
	)

	if ok = errors.As(err, &e); !ok {
		return consts.ErrorCodeHTTPStatus[consts.UnknownError]
	}
	if httpStatus, ok = consts.ErrorCodeHTTPStatus[e.Code]; !ok {
		return consts.ErrorCodeHTTPStatus[consts.UnknownError]
	}
	return httpStatus
}

func GetErrorResult(err error) *Result {
	var (
		e  *ErrorResult
		ok bool
	)
	if ok = errors.As(err, &e); !ok {
		return UnknownErrorResult
	}
	return NewErrorResult(e.Code, e.Msg)
}
