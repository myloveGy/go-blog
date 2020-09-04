package handler

import (
	"github.com/gin-gonic/gin"

	"blog/app/request"
	"blog/app/service"
	"blog/global"
	"blog/pkg/app"
	error2 "blog/pkg/error"
)

func GetApp(c *gin.Context) {
	param := request.AppRequest{}
	response := app.NewResponse(c)

	// 验证数据
	valid, errs := app.BindAndValid(c, &param)
	if valid {
		global.Logger.ErrorFormat("get app app.BindAndValid errs: %v", errs)
		errRsp := error2.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	if err := svc.CheckApp(&param); err != nil {
		global.Logger.ErrorFormat("get app svc.CheckApp err: %v", err)
		response.ToErrorResponse(error2.UnauthorizedAuthNotExists)
		return
	}

	token, err := app.GenerateToken(param.AppId, param.AppSecret)
	if err != nil {
		global.Logger.ErrorFormat("get app app.GenerateToken err: %v", err)
		response.ToErrorResponse(error2.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
