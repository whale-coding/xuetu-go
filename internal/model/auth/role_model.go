package auth

import "xuetu-project/internal/model"

// Role 角色模型
type Role struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleName string `gorm:"column:role_name;size:32" json:"roleName"`
	RoleKey  string `gorm:"column:role_key;size:64" json:"roleKey"`
	model.BaseModel
}
