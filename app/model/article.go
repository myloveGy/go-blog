package model

import "time"

type Article struct {
	ArticleId     int       `json:"id" db:"article_id" gorm:"primary_key"` // 主键ID
	UserId        int       `json:"user_id" db:"user_id"`                  // 用户ID
	Title         string    `json:"title" db:"title"`                      // 文章标题
	Desc          string    `json:"desc" db:"desc"`                        // 文章简述
	CoverImageUrl string    `json:"cover_image_url" db:"cover_image_url"`  // 封面图片地址
	Content       string    `json:"content" db:"content"`                  // 文章内容
	Status        int8      `json:"status" db:"status"`                    // 状态 10 启用 5 停用
	IsDel         int8      `json:"is_del" db:"is_del"`                    // 删除状态 10 未删除 5 已删除
	DeletedAt     time.Time `json:"deleted_at" db:"deleted_at"`            // 删除时间
	CreatedAt     time.Time `json:"created_at" db:"created_at"`            // 创建时间
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`            // 修改时间
}

func (*Article) TableName() string {
	return "article"
}

func (*Article) PK() string {
	return "article_id"
}
