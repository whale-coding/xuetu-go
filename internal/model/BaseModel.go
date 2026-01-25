package model

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel 通用字段基类
type BaseModel struct {
	CreatedBy   string    `gorm:"column:created_by;size:32" json:"createdBy"`
	CreatedTime time.Time `gorm:"column:created_time;type:datetime;not null" json:"createdTime"`
	UpdatedBy   string    `gorm:"column:update_by;size:32" json:"updatedBy"`
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime;not null" json:"updateTime"`
	IsDeleted   int       `gorm:"column:is_deleted;default:0;index" json:"isDeleted"` // 0:未删除, 1:已删除
}

//BeforeCreate 和 BeforeUpdate 是 GORM 的 回调钩子（Hooks），会在 Create() / Save() / Update() 时自动触发

// BeforeCreate GORM 钩子：自动设置创建/更新时间
func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	b.CreatedTime = now
	b.UpdateTime = now
	if b.CreatedBy == "" {
		b.CreatedBy = "system" // 或从上下文获取当前用户
	}
	if b.UpdatedBy == "" {
		b.UpdatedBy = "system" // 或从上下文获取当前用户
	}
	return nil
}

// BeforeUpdate GORM 钩子：自动更新时间
func (b *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	b.UpdateTime = time.Now()
	if b.UpdatedBy == "" {
		b.UpdatedBy = "system"
	}
	return nil
}
