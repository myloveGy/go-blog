package v1

import (
	"github.com/gin-gonic/gin"

	"blog/app/request"
	response2 "blog/app/response"
	"blog/app/service"
	"blog/global"
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

func (t Tag) List(c *gin.Context) {
	var params request.TagListRequest
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if valid {
		global.Logger.ErrorFormat(" app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(error2.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}

	total, err := svc.TagCount(&request.TagCountRequest{
		Name:   params.Name,
		Status: params.Status,
	})
	if err != nil {
		global.Logger.ErrorFormat("svc.TagCount err: %v", err)
		response.ToErrorResponse(error2.ErrorTagCountFail)
		return
	}

	tags, err := svc.TagGetList(&params, &pager)
	if err != nil {
		global.Logger.ErrorFormat("svc.TagGetList err: %v", err)
		response.ToErrorResponse(error2.ErrorTagListFail)
		return
	}

	response.ToResponseList(tags, total)
	return
}

func (t Tag) Create(c *gin.Context) {
	var params request.TagCreateRequest
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if valid {
		global.Logger.ErrorFormat(" app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(error2.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	tag, err := svc.TagCreate(&params)
	if err != nil {
		global.Logger.ErrorFormat("svc.TagCreate err: %v", err)
		response.ToErrorResponse(error2.ErrorTagCreateFail)
		return
	}

	response.ToResponse(response2.NewTagResponse(tag))
	return
}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}
