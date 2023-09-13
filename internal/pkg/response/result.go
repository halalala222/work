package response

import "work/internal/consts"

type Result struct {
	Code consts.ResponseCode `json:"code"`
	Msg  string              `json:"msg"`
	Data any                 `json:"data"`
}

func NewSuccessResult(data any) *Result {
	return &Result{
		Code: consts.Success,
		Msg:  SuccessMsg,
		Data: data,
	}
}

func NewErrorResult(code consts.ResponseCode, msg string) *Result {
	return &Result{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
