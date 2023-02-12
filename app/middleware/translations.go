package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	zhTwTranslations "github.com/go-playground/validator/v10/translations/zh_tw"
)

func NewUniversalTranslator() *ut.UniversalTranslator {
	uni := ut.New(zh.New(), en.New(), zh.New(), zh_Hant_TW.New())
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if trans, found := uni.GetTranslator("en"); found {
			_ = enTranslations.RegisterDefaultTranslations(v, trans)
		}

		if trans, found := uni.GetTranslator("zh"); found {
			_ = zhTranslations.RegisterDefaultTranslations(v, trans)
		}

		if trans, found := uni.GetTranslator("zh_Hant_TW"); found {
			_ = zhTwTranslations.RegisterDefaultTranslations(v, trans)
		}
	}

	return uni
}

func Translations() gin.HandlerFunc {
	uni := NewUniversalTranslator()
	return func(c *gin.Context) {
		locale := c.GetHeader("locale")
		if locale == "zh-tw" || locale == "zh-TW" || locale == "tw" {
			locale = "zh_Hant_TW"
		}

		trans, _ := uni.GetTranslator(locale)
		c.Set("trans", trans)

		c.Next()
	}
}
