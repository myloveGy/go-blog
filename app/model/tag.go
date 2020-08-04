package model

import "time"

type Tag struct {
	TagId     int       `json:"tag_id" db:"tag_id" gorm:"primary_key"` // tag_id
	Name      string    `json:"name" db:"name"`                        // 标签名称
	Status    int8      `json:"status" db:"status"`                    // 状态 10 启用 5 停用
	CreatedAt time.Time `json:"created_at" db:"created_at"`            // 创建时间
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`            // 修改时间
}

func (*Tag) TableName() string {
	return "tag"
}

func (*Tag) PK() string {
	return "tag_id"
}
