package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	TagId     int       `json:"tag_id" db:"tag_id" gorm:"primary_key"`                             // tag_id
	Name      string    `json:"name" db:"name"`                                                    // 标签名称
	Status    uint8     `json:"status" db:"status"`                                                // 状态 10 启用 5 停用
	CreatedAt time.Time `json:"created_at" db:"created_at" gorm:"type:datetime;column:created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" gorm:"type:datetime;column:updated_at"` // 修改时间
}

func (*Tag) TableName() string {
	return "tag"
}

func (*Tag) PK() string {
	return "tag_id"
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}

	db.Where("status = ?", t.Status)
	if err := db.Model(&t).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (t Tag) List(db *gorm.DB, offset, size int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if offset >= 0 && size > 0 {
		db = db.Offset(offset).Limit(size)
	}

	if t.Name != "" {
		db.Where("name like ?", t.Name)
	}

	db = db.Where("status = ?", t.Status)
	if err = db.Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (t *Tag) Create(db *gorm.DB) error {
	return db.Create(t).Error
}

func (t *Tag) Update(db *gorm.DB, values interface{}) error {
	return db.Model(t).Where("tag_id = ?", t.TagId).Updates(values).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("tag_id = ?", t.TagId).Delete(&t).Error
}
