package v1

import (
	"github.com/gin-gonic/gin"

	"blog/pkg/app"
	error2 "blog/pkg/error"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// @Summary 获取标签详情
// @Produce json
// @Param id query int true "标签ID"
// @Success 200 {object} model.Tag "成功"
// @Failure 500 {object} error.Error "内部错误"
// @Router /api/v1/tags/{id} [get]
func (t Tag) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(error2.ServerError)
	return
}

func (t Tag) List(c *gin.Context)   {}
func (t Tag) Create(c *gin.Context) {}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}
