package adminusercontroller

import (
	"jyjey-mini/api/adminuser/adminuserservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

//InitAdminUserRouter 加载后台用户路由
func InitAdminUserRouter(router *gin.RouterGroup) {
	v1 := router.Group("adminUser")
	{
		//新增管理员用户
		v1.POST("/createAdminUser", func(c *gin.Context) {
			var adminUserService adminuserservice.AdminUser

			err := c.ShouldBindJSON(&adminUserService)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			adminUserService.CreateAdminUser(c)
		})
	}
}
