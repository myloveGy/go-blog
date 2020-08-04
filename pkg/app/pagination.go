package app

import (
	"github.com/gin-gonic/gin"

	"blog/global"
	"blog/pkg/convert"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	size := convert.StrTo(c.Query("page_size")).MustInt()
	if size <= 0 {
		return global.AppSetting.DefaultPageSize
	}

	if size > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}

	return size
}

func GetPageOffset(page, size int) int {
	if page > 0 {
		return (page - 1) * size
	}

	return 0
}
