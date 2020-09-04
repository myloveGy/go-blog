package request

type TagListRequest struct {
	Name   string `form:"name" binding:"max=100"`
	Status uint8  `form:"status,default=10" binding:"oneof=5 10"`
}

type TagCreateRequest struct {
	Name   string `form:"name" binding:"required,min=2,max=100"`
	Status uint8  `form:"status,default=10" binding:"oneof=5 10"`
}

type TagDeleteRequest struct {
	TagId int `form:"tag_id" binding:"required,gte=1"`
}
