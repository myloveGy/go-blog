package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"blog/app/request"
	"blog/app/response"
	"blog/app/service"
	"blog/global"
	"blog/pkg/app"
	error2 "blog/pkg/error"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) find(c *gin.Context) (*response.TagResponse, *error2.Error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, error2.InvalidParams
	}

	tag, err := service.New(c.Request.Context()).TagFirst(id)
	if err != nil {
		return nil, error2.TagFirstError
	}

	return tag, nil
}

// @Summary 获取标签详情
// @Produce json
// @Param id query int true "标签ID"
// @Success 200 {object} response.TagResponse "成功"
// @Failure 500 {object} error.Error "内部错误"
// @Router /api/v1/tags/{id} [get]
func (t Tag) Get(c *gin.Context) {
	rep := app.NewResponse(c)
	tag, err := t.find(c)
	if err != nil {
		rep.ToErrorResponse(err)
		return
	}

	rep.ToResponse(tag)
	return
}

// @Summary 获取标签列表
// @Produce json
// @Param status query int false "标签状态(10 启用 5 停用)"
// @Param name   query string false "标签名称"
// @Success 200 {object} model.Tag "成功"
// @Failure 500 {object} error.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	var params request.TagListRequest
	rep := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if valid {
		global.Logger.ErrorFormat(" app.BindAndValid errs: %v", errs)
		rep.ToErrorResponse(error2.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}

	total, err := svc.TagCount(&request.TagListRequest{
		Name:   params.Name,
		Status: params.Status,
	})
	if err != nil {
		global.Logger.ErrorFormat("svc.TagCount err: %v", err)
		rep.ToErrorResponse(error2.TagCountFail)
		return
	}

	tags, err := svc.TagGetList(&params, &pager)
	if err != nil {
		global.Logger.ErrorFormat("svc.TagGetList err: %v", err)
		rep.ToErrorResponse(error2.TagListFail)
		return
	}

	rep.ToResponseList(tags, total)
	return
}

// @Summary 添加标签
// @Produce json
// @Param status query int false "标签状态(10 启用 5 停用)"
// @Param name   query string true "标签名称"
// @Success 200 {object} model.Tag "成功"
// @Failure 500 {object} error.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	var params request.TagCreateRequest
	rep := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if valid {
		global.Logger.ErrorFormat(" app.BindAndValid errs: %v", errs)
		rep.ToErrorResponse(error2.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	tag, err := svc.TagCreate(&params)
	if err != nil {
		global.Logger.ErrorFormat("svc.TagCreate err: %v", err)
		rep.ToErrorResponse(error2.TagCreateFail)
		return
	}

	rep.ToResponse(tag)
	return
}

// @Summary 修改标签
// @Produce json
// @Param status query int false "标签状态(10 启用 5 停用)"
// @Param name   query string true "标签名称"
// @Param name   query string true "标签名称"
// @Success 200 {object} model.Tag "成功"
// @Failure 500 {object} error.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	var params request.TagCreateRequest
	rep := app.NewResponse(c)

	// 验证数据
	valid, errs := app.BindAndValid(c, &params)
	if valid {
		global.Logger.ErrorFormat("Update app.BindAndValid errs: %v", errs)
		rep.ToErrorResponse(error2.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 查询数据是否存在
	tmpTag, err := t.find(c)
	if err != nil {
		rep.ToErrorResponse(err)
		return
	}

	// 修改数据
	svc := service.New(c.Request.Context())
	if err1 := svc.TagUpdate(tmpTag.TagId, &params); err1 != nil {
		rep.ToErrorResponse(error2.TagUpdateFail)
		return
	}

	rep.ToResponse(gin.H{"ok": true})
	return
}

// @Summary 添加标签
// @Produce json
// @Param status query int false "标签状态(10 启用 5 停用)"
// @Param name   query string true "标签名称"
// @Success 200 {object} model.Tag "成功"
// @Failure 500 {object} error.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	// 查询数据是否存在
	rep := app.NewResponse(c)
	tmpTag, err := t.find(c)
	if err != nil {
		rep.ToErrorResponse(err)
		return
	}

	// 删除数据
	svc := service.New(c.Request.Context())
	if err1 := svc.TagDelete(tmpTag.TagId); err1 != nil {
		rep.ToErrorResponse(error2.TagDeleteFail)
		return
	}

	rep.ToResponse(gin.H{"ok": true})
	return
}
