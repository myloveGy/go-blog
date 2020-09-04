package request

type AppRequest struct {
	AppId     string `form:"app_id" binding:"required,min=10"`
	AppSecret string `form:"app_secret" binding:"required,min=10"`
}
