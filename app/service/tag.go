package service

import (
	"blog/app/model"
	"blog/app/request"
	"blog/pkg/app"
)

func (s *Service) TagCount(countRequest *request.TagCountRequest) (int, error) {
	return s.dao.TagCount(countRequest.Name, countRequest.Status)
}

func (s *Service) TagGetList(listRequest *request.TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return s.dao.TagGetList(listRequest.Name, listRequest.Status, pager.Page, pager.PageSize)
}

func (s *Service) TagCreate(createRequest *request.TagCreateRequest) (*model.Tag, error) {
	return s.dao.TagCreate(createRequest.Name, createRequest.Status)
}

func (s *Service) TagUpdate(updateRequest *request.TagUpdateRequest) error {
	return s.dao.TagUpdate(updateRequest.TagId, updateRequest.Name, updateRequest.Status)
}

func (s *Service) TagDelete(deleteRequest *request.TagDeleteRequest) error {
	return s.dao.TagDelete(deleteRequest.TagId)
}
