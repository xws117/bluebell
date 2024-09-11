package controllder

import (
	"github.com/gin-gonic/gin"
	"github.com/xws117/bluebell/dao/mysql"
	"github.com/xws117/bluebell/model"
)

func Loginfunc() gin.HandlerFunc {
	return func(context *gin.Context) {
		var u model.User
		if err := context.ShouldBindBodyWithJSON(&u); err != nil {
			ResponseErrorWithMsg(context, CodeInvalidParams, err.Error())
			return
		}
		if err := mysql.Login(&u); err != nil {
			//ResponseError(c, CodeInvalidPassword)
			return
		}
	}
}
