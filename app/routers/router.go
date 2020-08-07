package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	v1 "blog/app/handler/v1"
	"blog/app/middleware"

	_ "blog/docs"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := r.Group("api/v1")
	{
		tag := v1.NewTag()

		// 标签处理
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.GET("/tags/:id", tag.Get)
		apiV1.PATCH("/tags/:id/state", tag.Update)
		apiV1.GET("/tags", tag.List)

		article := v1.NewArticle()
		// 文章处理
		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.PATCH("/articles/:id/state", article.Update)
		apiV1.GET("/articles", article.List)
	}

	return r
}
