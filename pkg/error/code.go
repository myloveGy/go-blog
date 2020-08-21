package error

var (
	Success                   = NewError(10000, "成功")
	InvalidParams             = NewError(40001, "参数错误")
	NotFound                  = NewError(40004, "Not Found")
	UnauthorizedAuthNotExists = NewError(40030, "鉴权失败: 找不到对应的AppKey")
	UnauthorizedTokenError    = NewError(40031, "鉴权失败: Token 错误")
	UnauthorizedTokenTimeout  = NewError(40032, "鉴权失败: Token 超时")
	UnauthorizedTokenGenerate = NewError(40033, "鉴权失败: Token 生成失败")
	ServerError               = NewError(40005, "服务内部错误")
	TooManyRequests           = NewError(40007, "请求过多")

	// 标签处理失败
	TagListFail   = NewError(50101, "获取标签列表失败")
	TagCreateFail = NewError(50102, "创建标签失败")
	TagUpdateFail = NewError(50103, "修改标签失败")
	TagDeleteFail = NewError(50104, "删除标签失败")
	TagCountFail  = NewError(50105, "统计标签失败")
	TagFirstError = NewError(50106, "标签不存在")

	// 文章错误处理
	ArticleFail       = NewError(60101, "文章不存在")
	ArticleCreateFail = NewError(60102, "创建文章失败")
	ArticleUpdateFail = NewError(60103, "修改文章失败")
	ArticleDeleteFail = NewError(60104, "删除文章失败")
	ArticleListFail   = NewError(60105, "获取文章列表失败")
)
