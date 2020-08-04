package model

import "time"

type ArticleTag struct {
	Id        int       `json:"id" db:"id" gorm:"primary_key"` // id
	ArticleId int       `json:"article_id" db:"article_id"`    // 文章ID
	TagId     int       `json:"tag_id" db:"tag_id"`            // 标签ID
	CreatedAt time.Time `json:"created_at" db:"created_at"`    // 创建时间
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`    // 修改时间
}

func (*ArticleTag) TableName() string {
	return "article_tag"
}

func (*ArticleTag) PK() string {
	return "id"
}
