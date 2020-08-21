package repository

import (
	"blog/app/model"
	"blog/pkg/app"
)

type Article struct {
	ArticleId     int    `json:"article_id"`
	TagId         int    `json:"tag_id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	Status        int8   `json:"status"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

func (d *Dao) ArticleCreate(param *Article) (*model.Article, error) {
	article := model.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		Status:        param.Status,
	}

	return article.Create(d.engine)
}

func (d *Dao) ArticleUpdate(param *Article) error {
	article := model.Article{ArticleId: param.ArticleId}
	values := map[string]interface{}{
		"status": param.Status,
	}

	if param.Title != "" {
		values["title"] = param.Title
	}

	if param.Desc != "" {
		values["desc"] = param.Desc
	}

	if param.CoverImageUrl != "" {
		values["cover_image_url"] = param.CoverImageUrl
	}

	if param.Content != "" {
		values["content"] = param.Content
	}

	return article.Update(d.engine, values)
}

func (d *Dao) ArticleFirst(id int) (*model.Article, error) {
	article := model.Article{ArticleId: id}
	return article.First(d.engine)
}

func (d *Dao) ArticleDelete(id int) error {
	article := model.Article{ArticleId: id}
	return article.Delete(d.engine)
}

func (d *Dao) ArticleCountListByTagId(id int, status int8) (int, error) {
	article := model.Article{Status: status}
	return article.CountByTagId(d.engine, id)
}

func (d *Dao) ArticleListByTagId(id int, status int8, page, pageSize int) ([]*model.ArticleRow, error) {
	article := model.Article{Status: status}
	return article.ListByTagId(d.engine, id, app.GetPageOffset(page, pageSize), pageSize)
}
