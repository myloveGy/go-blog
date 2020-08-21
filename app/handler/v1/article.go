package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"blog/app/request"
	"blog/app/service"
	"blog/global"
	"blog/pkg/app"
	error2 "blog/pkg/error"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) find(c *gin.Context) (*service.Article, *error2.Error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, error2.InvalidParams
	}

	svc := service.New(c.Request.Context())
	article, err := svc.ArticleFirst(id)
	if err != nil {
		return nil, error2.ArticleFail
	}

	return article, nil
}

// @Summary 获取文章详情
// @Produce json
// @Param article_id query int true "文章ID"
// @Success 200 {object} response.TagResponse "成功"
// @Failure 500 {object} error.Error "内部错误"
// @Router /api/v1/article/{id} [get]
func (a Article) Get(c *gin.Context) {
	rep := app.NewResponse(c)
	article, err := a.find(c)
	if err != nil {
		rep.ToErrorResponse(err)
		return
	}

	rep.ToResponse(article)
}

// @Summary 添加文章信息
// @Produce json
// @Param article_id query int true "文章ID"
// @Success 200 {object} response.TagResponse "成功"
// @Failure 500 {object} error.Error "内部错误"
// @Router /api/v1/article [get]
func (a Article) List(c *gin.Context) {
	// 定义变量
	var (
		param request.ArticleListRequest
		rep   = app.NewResponse(c)
	)

	// 验证绑定数据
	valida, errs := app.BindAndValid(c, &param)
	if valida {
		global.Logger.ErrorFormat("Article List app.BindAndValid errs: %v", errs)
		rep.ToErrorResponse(error2.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 查询数据
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	articles, totalRows, err := svc.ArticleList(&param, &pager)
	if err != nil {
		rep.ToErrorResponse(error2.ArticleListFail)
		return
	}

	// 响应数据
	rep.ToResponseList(articles, totalRows)
	return
}

// @Summary 删除文章信息
// @Produce json
// @Param article_id query int true "文章ID"
// @Success 200 {object} response.TagResponse "成功"
// @Failure 500 {object} error.Error "内部错误"
// @Router /api/v1/article/{id} [post]
func (a Article) Create(c *gin.Context) {
	// 定义变量
	var (
		param request.ArticleCreateRequest
		rep   = app.NewResponse(c)
	)

	// 验证绑定数据
	valida, errs := app.BindAndValid(c, &param)
	if valida {
		global.Logger.ErrorFormat("Article Create app.BindAndValid errs: %v", errs)
		rep.ToErrorResponse(error2.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 新增数据
	svc := service.New(c.Request.Context())
	article, err := svc.ArticleCreate(&param)
	if err != nil {
		rep.ToErrorResponse(error2.ArticleCreateFail)
		return
	}

	// 响应数据
	rep.ToResponse(article)
	return
}

// @Summary 修改文章信息
// @Produce json
// @Param article_id query int true "文章ID"
// @Success 200 {object} response.TagResponse "成功"
// @Failure 500 {object} error.Error "内部错误"
// @Router /api/v1/article/{id} [put]
func (a Article) Update(c *gin.Context) {
	// 定义变量
	var (
		param request.ArticleUpdateRequest
		rep   = app.NewResponse(c)
	)

	// 验证绑定数据
	valida, errs := app.BindAndValid(c, &param)
	if valida {
		global.Logger.ErrorFormat("Article Update app.BindAndValid errs: %v", errs)
		rep.ToErrorResponse(error2.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 查询数据
	article, err := a.find(c)
	if err != nil {
		rep.ToErrorResponse(err)
		return
	}

	// 修改数据
	svc := service.New(c.Request.Context())
	if err2 := svc.ArticleUpdate(article.ArticleId, &param); err2 != nil {
		rep.ToErrorResponse(error2.ArticleUpdateFail)
		return
	}

	// 返回
	rep.ToResponse(gin.H{"ok": true})
	return
}

// @Summary 删除文章信息
// @Produce json
// @Param article_id query int true "文章ID"
// @Success 200 {object} response.TagResponse "成功"
// @Failure 500 {object} error.Error "内部错误"
// @Router /api/v1/article/{id} [delete]
func (a Article) Delete(c *gin.Context) {
	rep := app.NewResponse(c)
	article, err := a.find(c)
	if err != nil {
		rep.ToErrorResponse(err)
		return
	}

	svc := service.New(c.Request.Context())
	if err2 := svc.ArticleDelete(article.ArticleId); err2 != nil {
		rep.ToErrorResponse(error2.ArticleDeleteFail)
		return
	}

	rep.ToResponse(gin.H{"ok": true})
	return
}
