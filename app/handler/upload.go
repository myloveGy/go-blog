package handler

import (
	"github.com/gin-gonic/gin"

	"blog/app/service"
	"blog/global"
	"blog/pkg/app"
	"blog/pkg/convert"
	error2 "blog/pkg/error"
	"blog/pkg/upload"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	// 存在错误
	if err != nil {
		errRsp := error2.InvalidParams.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	// 验证上传了文件
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(error2.InvalidParams.WithDetails("没有上传文件或者文件类型type为空"))
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.ErrorFormat("svc.UploadFile err: %v", err)
		errRsp := error2.UploadFileFail.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	response.ToResponse(gin.H{
		"url":      fileInfo.AccessUrl,
		"filename": fileInfo.Name,
	})
}
