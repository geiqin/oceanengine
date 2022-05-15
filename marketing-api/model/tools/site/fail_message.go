package site

import "github.com/geiqin/oceanengine/marketing-api/model"

// FailMessage 错误信息
type FailMessage struct {
	// Message 错误信息
	Message string `json:"message,omitempty"`
	// SiteID 更新失败的site_id
	SiteID model.FlexUint64 `json:"site_id,omitempty"`
}
