package repository

import (
	"blog/app/model"
)

func (d *Dao) GetApp(appId, appSecret string) (*model.App, error) {
	auth := &model.App{AppId: appId, AppSecret: appSecret}
	return auth.Get(d.engine)
}
