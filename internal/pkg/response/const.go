package response

import (
	"work/internal/consts"
)

const (
	UnKnownMsg = "unKnown"
	SuccessMsg = "success"
)

var (
	ServerError              = NewError(consts.ServerError, "internal server error")
	FormError                = NewError(consts.RequestError, "request form error")
	SuccessWithoutDataResult = NewSuccessResult(nil)
	UnknownErrorResult       = NewErrorResult(consts.UnknownError, UnKnownMsg)
)
