package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"blog/pkg/app"
	error2 "blog/pkg/error"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var token string

		// 先从请求参数中获取token, 没有那么使用 Header 中的信息
		if s, exist := context.GetQuery("token"); exist {
			token = s
		} else {
			token = context.GetHeader("Token")
		}

		// token 无效
		if token == "" {
			app.NewResponse(context).ToErrorResponse(error2.InvalidParams.WithDetails("签名Token信息为空"))
			context.Abort()
			return
		}

		// 解析token 失败
		if _, err := app.ParseToken(token); err != nil {
			var errCode *error2.Error
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				errCode = error2.UnauthorizedTokenTimeout
			default:
				errCode = error2.UnauthorizedTokenError
			}

			app.NewResponse(context).ToErrorResponse(errCode)
			context.Abort()
			return
		}

		context.Next()
	}
}
