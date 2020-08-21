package request

type ArticleFirstRequest struct {
	Status int8 `form:"status,default=10" binding:"oneof=5 10"`
}

type ArticleListRequest struct {
	Status int8 `form:"status,default=10" binding:"oneof=5 10"`
	TagId  int  `form:"tag_id" binding:"required,min=1"`
}

type ArticleCreateRequest struct {
	Status        int8   `form:"status,default=10" binding:"oneof=5 10"`
	TagId         int    `form:"tag_id" binding:"required,min=1"`
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" binding:"required,min=2,max=255"`
	Content       string `form:"content" binding:"required,min=2"`
	CoverImageUrl string `form:"cover_image_url'"`
}

type ArticleUpdateRequest struct {
	Status        int8   `form:"status,default=10" binding:"oneof=5 10"`
	TagId         int    `form:"tag_id" binding:"required,min=1"`
	Title         string `form:"title" binding:"min=2,max=100"`
	Desc          string `form:"desc" binding:"min=2,max=255"`
	Content       string `form:"content" binding:"min=2"`
	CoverImageUrl string `form:"cover_image_url'"`
}
