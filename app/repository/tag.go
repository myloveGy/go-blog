package repository

import (
	"blog/app/model"
	"blog/pkg/app"
)

func (d *Dao) TagFirst(tagId int) (*model.Tag, error) {
	tag := model.Tag{TagId: tagId}
	return &tag, tag.First(d.engine)
}

func (d *Dao) TagCount(name string, status uint8) (int, error) {
	tag := model.Tag{Name: name, Status: status}
	return tag.Count(d.engine)
}

func (d *Dao) TagGetList(name string, status uint8, page, size int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, Status: status}
	pageOffset := app.GetPageOffset(page, size)
	return tag.List(d.engine, pageOffset, size)
}

func (d *Dao) TagCreate(name string, status uint8) (*model.Tag, error) {
	tag := model.Tag{Name: name, Status: status}
	return &tag, tag.Create(d.engine)
}

func (d *Dao) TagUpdate(tagId int, name string, status uint8) error {
	tag := model.Tag{TagId: tagId, Name: name, Status: status}
	return tag.Update(d.engine, tag)
}

func (d *Dao) TagDelete(tagId int) error {
	tag := model.Tag{TagId: tagId}
	return tag.Delete(d.engine)
}
