package http

import (
	"super_admin/http/controller"
	"super_admin/http/middleware"

	"github.com/gin-gonic/gin"
)

func loadRouter() (router *gin.Engine) {
	gin.SetMode(gin.DebugMode)
	router = gin.New()
	//接口访问日志
	router.Use(middleware.RequestLogger())
	//业务错误日志(controller最终抛出)
	router.Use(middleware.ControllerErrorLogger())
	//routes
	router.GET("health", controller.Health)
	//登录模块
	auth := router.Group("/auth")
	{
		auth.POST("/signin/username", controller.STDwrapperJSON(controller.SignInUsername)) //用户名登录
	}
	//需鉴权的接口
	guarded := router.Group("")
	{
		guarded.Use(middleware.Auth())
		//用户
		user := guarded.Group("user")
		{
			user.POST("create", controller.STDwrapperJSON(controller.AddUser))
			user.POST("update", controller.STDwrapperJSON(controller.UpdateUser))
			user.POST("search", controller.STDwrapperJSON(controller.SearchUser))
		}
		//估值
		valuate := guarded.Group("valuate")
		{
			valuate.POST("get_details", controller.STDwrapperJSON(controller.GetValuateDetails))
		}
		//转发
		guarded.Any("enterprise/*url", controller.STDwrapperJSON(controller.ForwardCompanyService))
		guarded.Any("group/*url", controller.STDwrapperJSON(controller.FowardGroupService))
		guarded.Any("audit/*url", controller.STDwrapperJSON(controller.FowardAuditService))
	}
	return
}
