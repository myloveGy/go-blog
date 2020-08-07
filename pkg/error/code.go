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
	ErrorTagListFail   = NewError(50101, "获取标签列表失败")
	ErrorTagCreateFail = NewError(50102, "创建标签失败")
	ErrorTagUpdateFail = NewError(50103, "修改标签失败")
	ErrorTagDeleteFail = NewError(50104, "删除标签失败")
	ErrorTagCountFail  = NewError(50105, "统计标签失败")
)
