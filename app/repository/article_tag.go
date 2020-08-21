package repository

import (
	"blog/app/model"
)

func (d *Dao) ArticleTagFirstByArticleId(articleId int) (*model.ArticleTag, error) {
	articleTag := model.ArticleTag{ArticleId: articleId}
	return articleTag.FirstByArticleId(d.engine)
}

func (d *Dao) ArticleTagFirstByTagId(tagId int) (*model.ArticleTag, error) {
	articleTag := model.ArticleTag{TagId: tagId}
	return articleTag.FirstByTagId(d.engine)
}

func (d *Dao) ArticleTagListByArticleIds(articleIds []int) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{}
	return articleTag.ListByArticleIds(d.engine, articleIds)
}

func (d *Dao) ArticleTagCreate(articleId, tagId int) (*model.ArticleTag, error) {
	articleTag := model.ArticleTag{ArticleId: articleId, TagId: tagId}
	return articleTag.Create(d.engine)
}

func (d *Dao) ArticleTagUpdate(articleId, tagId int) error {
	articleTag := model.ArticleTag{ArticleId: articleId, TagId: tagId}
	return articleTag.Update(d.engine, nil)
}

func (d *Dao) ArticleTagDelete(articleId, tagId int) error {
	articleTag := model.ArticleTag{ArticleId: articleId, TagId: tagId}
	return articleTag.Delete(d.engine)
}
