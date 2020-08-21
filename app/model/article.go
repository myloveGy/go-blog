package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	ArticleId     int       `json:"article_id" db:"article_id" gorm:"primary_key"` // 主键ID
	UserId        int       `json:"user_id" db:"user_id"`                          // 用户ID
	Title         string    `json:"title" db:"title"`                              // 文章标题
	Desc          string    `json:"desc" db:"desc"`                                // 文章简述
	CoverImageUrl string    `json:"cover_image_url" db:"cover_image_url"`          // 封面图片地址
	Content       string    `json:"content" db:"content"`                          // 文章内容
	Status        int8      `json:"status" db:"status"`                            // 状态 10 启用 5 停用
	IsDel         int8      `json:"is_del" db:"is_del"`                            // 删除状态 10 未删除 5 已删除
	DeletedAt     time.Time `json:"deleted_at" db:"deleted_at"`                    // 删除时间
	CreatedAt     time.Time `json:"created_at" db:"created_at"`                    // 创建时间
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`                    // 修改时间
}

type ArticleRow struct {
	ArticleId     int
	TagId         int
	TagName       string
	ArticleTitle  string
	ArticleDesc   string
	CoverImageUrl string
	Content       string
	Status        int8
}

func (*Article) TableName() string {
	return "article"
}

func (*Article) PK() string {
	return "article_id"
}

func (a *Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(a).Error; err != nil {
		return nil, err
	}

	return a, nil
}

func (a *Article) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&a).Update(values).Error
}

func (a *Article) First(db *gorm.DB) (*Article, error) {
	if err := db.First(a).Error; err != nil {
		return nil, err
	}

	return a, nil
}

func (a *Article) Delete(db *gorm.DB) error {
	return db.Delete(a).Error
}

func (a *Article) ListByTagId(db *gorm.DB, tagId, pageOffset, pageSize int) ([]*ArticleRow, error) {
	fields := []string{
		"article.article_id",
		"article.title AS article_title",
		"article.desc AS article_desc",
		"article.cover_image_url",
		"article.content",
		"tag.tag_id",
		"tag.name as tag_name",
		"article.status",
	}

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	var (
		at  = ArticleTag{}
		tag = Tag{}
	)
	rows, err := db.Select(fields).Table(at.TableName()).
		Joins("LEFT JOIN "+tag.TableName()+" ON tag.tag_id = article_tag.tag_id").
		Joins("JOIN "+a.TableName()+" ON article.article_id = article_tag.article_id").
		Where("article_tag.tag_id = ? AND article.status", tagId, a.Status).
		Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var articles []*ArticleRow
	for rows.Next() {
		r := &ArticleRow{}
		if err := rows.Scan(
			&r.ArticleId,
			&r.ArticleTitle,
			&r.ArticleDesc,
			&r.CoverImageUrl,
			&r.Content,
			&r.TagId,
			&r.TagName,
			&r.Status,
		); err != nil {
			return nil, err
		}

		articles = append(articles, r)
	}

	return articles, nil
}

func (a *Article) CountByTagId(db *gorm.DB, tagId int) (int, error) {
	var (
		at    = ArticleTag{}
		tag   = Tag{}
		count int
	)
	if err := db.Table(at.TableName()).
		Joins("LEFT JOIN "+tag.TableName()+" ON tag.tag_id = article_tag.tag_id").
		Joins("JOIN "+a.TableName()+" ON article.article_id = article_tag.article_id").
		Where("article_tag.tag_id = ? AND article.status = ?", tagId, a.Status).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
