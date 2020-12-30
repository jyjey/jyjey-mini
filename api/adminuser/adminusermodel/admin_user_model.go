package adminusermodel

import (
	"jyjey-mini/mysql"
)

//AdminUser 管理员用户实体
type AdminUser struct {
	ID         uint   `gorm:"primaryKey"`
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	Avatar     string `gorm:"column:avatar"`
	Status     uint8  `gorm:"column:status"`
	Level      uint8  `gorm:"column:level"`
	RoleID     uint   `gorm:"column:role_id"`
	NikeName   string `gorm:"column:nike_name"`
	CreateTime uint   `gorm:"column:create_time"`
	LoginTime  uint   `gorm:"column:login_time"`
	Mobile     string `gorm:"column:mobile"`
	LoginIP    string `gorm:"column:login_ip"`
}

//TableName 重命名表
func (AdminUser) TableName() string {
	return "admin_user"
}

//Insert 创建管理员用户
func (adminUser *AdminUser) Insert() (id uint, err error) {
	result := mysql.DB.Create(&adminUser)
	id = adminUser.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
