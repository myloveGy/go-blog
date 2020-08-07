package service

import (
	"context"

	"blog/app/repository"
	"blog/global"
)

type Service struct {
	ctx context.Context
	dao *repository.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = repository.New(global.DBEngine)
	return svc
}
