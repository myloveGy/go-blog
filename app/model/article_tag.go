package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

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

func (a *ArticleTag) FirstByTagId(db *gorm.DB) (*ArticleTag, error) {
	if err := db.First(a).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return a, nil
}

func (a *ArticleTag) FirstByArticleId(db *gorm.DB) (*ArticleTag, error) {
	if err := db.First(a).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return a, nil
}

func (a *ArticleTag) ListByArticleIds(db *gorm.DB, articleIds []int) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	if err := db.Where("article_id IN (?)", articleIds).Find(&articleTags).Error; err != nil {
		return nil, err
	}

	return articleTags, nil
}

func (a *ArticleTag) ListByArticleId(db *gorm.DB) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	if err := db.Where("tag_id = ?", a.TagId).Find(&articleTags).Error; err != nil {
		return nil, err
	}

	return articleTags, nil
}

func (a *ArticleTag) Create(db *gorm.DB) (*ArticleTag, error) {
	if err := db.Create(a).Error; err != nil {
		return nil, err
	}

	return a, nil
}

func (a *ArticleTag) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&a).Update(values).Error
}

func (a *ArticleTag) First(db *gorm.DB) (*ArticleTag, error) {
	if err := db.First(a).Error; err != nil {
		return nil, err
	}

	return a, nil
}

func (a *ArticleTag) Delete(db *gorm.DB) error {
	return db.Delete(a).Error
}
