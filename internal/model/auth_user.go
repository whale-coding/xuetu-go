package model

// User 用户模型
type User struct {
	ID        uint   `gorm:"primarykey;autoIncrement" json:"id"` // 主键 ID自增
	Username  string `gorm:"column:user_name;size:50;not null;unique" json:"username"`
	Password  string `gorm:"size:100;not null" json:"-"` // 不返回给前端（json:"-"）
	Nickname  string `gorm:"column:nick_name;size:50" json:"nickname"`
	Avatar    string `gorm:"column:avatar;size:255" json:"avatar"`
	Email     string `gorm:"column:email;size:100" json:"email"`
	Phone     string `gorm:"column:phone;size:20" json:"phone"`
	Sex       int    `gorm:"column:sex;type:tinyint" json:"sex"` // 例如：0-男, 1-女
	Introduce string `gorm:"column:introduce;size:255" json:"introduce"`
	ExtJson   string `gorm:"column:ext_json;size:255" json:"extJson"`
	Status    int    `gorm:"column:status;type:tinyint;default:1" json:"status"` // 1-正常 0-禁用
	BaseModel
}

// TableName 显式指定表名，避免gorm自动使用复数表名
func (User) TableName() string {
	return "auth_user"
}
