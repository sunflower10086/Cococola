package myerr

import (
	"errors"
	"fmt"
)

type MyError struct {
	ErrMsg string `json:"status_msg,omitempty"`
}

func (e MyError) Error() string {
	return fmt.Sprintf("err_msg=%s", e.ErrMsg)
}

// NewMyError 创建一个新的errno
func NewMyError(msg string) MyError {
	return MyError{ErrMsg: msg}
}

var (
	Success  = NewMyError("success")
	ParamErr = NewMyError("传入的参数错误")
)

// WithMessage 改变errno中的msg
func (e MyError) WithMessage(msg string) MyError {
	e.ErrMsg = msg
	return e
}

// ConvertErr 把Error类型变成ErrNo类型
func ConvertErr(err error) MyError {
	Err := MyError{}

	if errors.As(err, &Err) {
		return Err
	}

	s := NewMyError(err.Error())
	return s
}
