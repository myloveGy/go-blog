package service

import (
	"blog/app/model"
	"blog/app/repository"
	"blog/app/request"
	"blog/pkg/app"
)

type Article struct {
	ArticleId     int        `json:"article_id"`
	Title         string     `json:"title"`
	Desc          string     `json:"desc"`
	Content       string     `json:"content"`
	CoverImageUrl string     `json:"cover_image_url"`
	Status        int8       `json:"status"`
	Tag           *model.Tag `json:"tag"`
}

func (s *Service) ArticleFirst(id int) (*Article, error) {
	article, err := s.dao.ArticleFirst(id)
	if err != nil {
		return nil, err
	}

	articleTag, err := s.dao.ArticleTagFirstByArticleId(id)
	if err != nil {
		return nil, err
	}

	tag, err := s.dao.TagFirst(articleTag.TagId)
	if err != nil {
		return nil, err
	}

	return &Article{
		ArticleId:     article.ArticleId,
		Title:         article.Title,
		Desc:          article.Desc,
		Content:       article.Content,
		CoverImageUrl: article.CoverImageUrl,
		Status:        article.Status,
		Tag:           tag,
	}, nil
}

func (s *Service) ArticleList(param *request.ArticleListRequest, pager *app.Pager) ([]*Article, int, error) {
	// 统计数量
	articleCount, err := s.dao.ArticleCountListByTagId(param.TagId, param.Status)
	if err != nil {
		return nil, 0, err
	}

	// 查询护具
	articles, err := s.dao.ArticleListByTagId(param.TagId, param.Status, pager.Page, pager.PageSize)
	if err != nil {
		return nil, articleCount, err
	}

	var articleList []*Article
	for _, article := range articles {
		articleList = append(articleList, &Article{
			ArticleId:     article.ArticleId,
			Title:         article.ArticleTitle,
			Desc:          article.ArticleDesc,
			Content:       article.Content,
			CoverImageUrl: article.CoverImageUrl,
			Status:        article.Status,
			Tag: &model.Tag{
				TagId: article.TagId,
				Name:  article.TagName,
			},
		})
	}

	return articleList, articleCount, nil
}

func (s *Service) ArticleCreate(param *request.ArticleCreateRequest) (*model.Article, error) {
	article, err := s.dao.ArticleCreate(&repository.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		Status:        param.Status,
	})

	if err != nil {
		return nil, err
	}

	if _, err := s.dao.ArticleTagCreate(article.ArticleId, param.TagId); err != nil {
		return nil, err
	}

	return article, nil
}

func (s *Service) ArticleUpdate(id int, param *request.ArticleUpdateRequest) error {
	return s.dao.ArticleUpdate(&repository.Article{
		ArticleId:     id,
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		Status:        param.Status,
	})
}

func (s *Service) ArticleDelete(id int) error {
	if err := s.dao.ArticleDelete(id); err != nil {
		return err
	}

	return s.dao.ArticleTagDelete(id, 0)
}
