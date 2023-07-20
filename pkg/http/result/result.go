package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CodeStatus int

const (
	SuccessCode  CodeStatus = 0
	ParamErrCode            = 1000
)

type Result struct {
	Code CodeStatus `json:"code"`
	Date any        `json:"data"`
}

func Success(data any) *Result {
	return &Result{Code: SuccessCode, Date: data}
}

func Fail(code CodeStatus, Err error) *Result {
	return &Result{Code: code, Date: code}
}

func SendResult(c *gin.Context, r *Result) {
	c.JSON(http.StatusOK, r)
}
