package router

import (
	"jyjey-mini/api/adminuser/adminusercontroller"
	"jyjey-mini/handler"
	"jyjey-mini/middlewares"

	"github.com/gin-gonic/gin"
)

//InitJYJEYRouter 初始化路由
func InitJYJEYRouter() {
	router := gin.Default()
	// 日志中间件
	router.Use(middlewares.LoggerToFile())
	//注意 Recover 要尽量放在第一个被加载
	//如不是的话，在recover前的中间件或路由，将不能被拦截到
	//程序的原理是：
	//1.请求进来，执行recover
	//2.程序异常，抛出panic
	//3.panic被 recover捕获，返回异常信息，并Abort,终止这次请求
	router.Use(handler.Recover)
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(middlewares.Cors())
	v1 := router.Group("v1")
	{
		adminusercontroller.InitAdminUserRouter(v1)
	}

	router.Run(":8080")
}
