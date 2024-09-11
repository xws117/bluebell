package controllder

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code    MyCode      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseErrorWithMsg(context *gin.Context, code MyCode, errMsg string) {
	rd := &ResponseData{
		Code: code, Message: errMsg, Data: nil,
	}
	context.JSON(http.StatusOK, rd)
}

func ResponseSuccess(context *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code: CodeSuccess, Message: CodeSuccess.Msg(), Data: data,
	}
	context.JSON(http.StatusOK, rd)
}
