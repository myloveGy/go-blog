package response

import (
	"blog/app/model"
	"blog/pkg/time"
	"blog/profile"
)

type TagResponse struct {
	TagId      int    `json:"tag_id"`      // tag_id
	Name       string `json:"name"`        // 标签名称
	Status     uint8  `json:"status"`      // 状态 10 启用 5 停用
	StatusName string `json:"status_name"` // 状态说明
	CreatedAt  string `json:"created_at"`  // 创建时间
	UpdatedAt  string `json:"updated_at"`  // 修改时间
}

func NewTagResponse(tag *model.Tag) *TagResponse {
	return &TagResponse{
		TagId:      tag.TagId,
		Name:       tag.Name,
		Status:     tag.Status,
		StatusName: profile.StatusMapNames[tag.Status],
		CreatedAt:  time.FormatDateTime(tag.CreatedAt),
		UpdatedAt:  time.FormatDateTime(tag.UpdatedAt),
	}
}
