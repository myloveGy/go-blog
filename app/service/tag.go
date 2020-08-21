package service

import (
	"blog/app/request"
	"blog/app/response"
	"blog/pkg/app"
)

func (s Service) TagFirst(id int) (*response.TagResponse, error) {
	tag, err := s.dao.TagFirst(id)
	return response.NewTagResponse(tag), err
}

func (s *Service) TagCount(countRequest *request.TagListRequest) (int, error) {
	return s.dao.TagCount(countRequest.Name, countRequest.Status)
}

func (s *Service) TagGetList(listRequest *request.TagListRequest, pager *app.Pager) ([]*response.TagResponse, error) {
	tagList, err := s.dao.TagGetList(listRequest.Name, listRequest.Status, pager.Page, pager.PageSize)
	if err != nil {
		return nil, err
	}

	// 格式化返回
	var tagResponseList = make([]*response.TagResponse, 0)
	for _, v := range tagList {
		tagResponseList = append(tagResponseList, response.NewTagResponse(v))
	}

	return tagResponseList, nil
}

func (s *Service) TagCreate(createRequest *request.TagCreateRequest) (*response.TagResponse, error) {
	tag, err := s.dao.TagCreate(createRequest.Name, createRequest.Status)
	return response.NewTagResponse(tag), err
}

func (s *Service) TagUpdate(id int, updateRequest *request.TagCreateRequest) error {
	return s.dao.TagUpdate(id, updateRequest.Name, updateRequest.Status)
}

func (s *Service) TagDelete(id int) error {
	return s.dao.TagDelete(id)
}
