package service

import (
	"errors"

	"blog/app/request"
)

func (s *Service) CheckApp(param *request.AppRequest) error {
	auth, err := s.dao.GetApp(param.AppId, param.AppSecret)
	if err != nil {
		return err
	}

	if auth.Id > 0 {
		return nil
	}

	return errors.New("auth info does not exist")
}
