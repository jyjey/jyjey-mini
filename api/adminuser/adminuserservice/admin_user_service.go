package adminuserservice

import (
	"jyjey-mini/api/adminuser/adminusermodel"
	"jyjey-mini/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AdminUser 管理员用户交互实体
type AdminUser struct {
	ID         uint   `json:"id"`
	Username   string `json:"username" validate:"required,min=5,max=35"`
	Password   string `json:"password" validate:"required,min=6,max=20"`
	Avatar     string `json:"avatar"`
	Status     uint8  `json:"status"`
	Level      uint8  `json:"level"`
	RoleID     uint   `json:"role_id"`
	NikeName   string `json:"nike_name" validate:"required,min=0,max=12"`
	CreateTime uint   `json:"create_time"`
	LoginTime  uint   `json:"login_time"`
	Mobile     string `json:"mobile" validate:"required"`
	LoginIP    string `json:"login_ip"`
}

//CreateAdminUser 创建管理员用户
func (adminUser *AdminUser) CreateAdminUser(c *gin.Context) {
	var adminUserModel adminusermodel.AdminUser
	msg, err := util.CheckStruct(adminUser)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "系统异常",
		})
		return
	}
	adminUserModel.Username = adminUser.Username
	//密码加密处理
	adminUserModel.Password, err = util.EncodePwd(adminUser.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": msg,
		})
		return
	}
	adminUserModel.Avatar = adminUser.Avatar
	adminUserModel.Status = adminUser.Status
	adminUserModel.Level = adminUser.Level
	adminUserModel.RoleID = adminUser.RoleID
	adminUserModel.NikeName = adminUser.NikeName
	adminUserModel.CreateTime = adminUser.CreateTime
	adminUserModel.LoginTime = adminUser.LoginTime
	adminUserModel.Mobile = adminUser.Mobile
	adminUserModel.LoginIP = adminUser.LoginIP
	id, err := adminUserModel.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Insert() error!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "success",
		"data":    id,
	})
	return
}
