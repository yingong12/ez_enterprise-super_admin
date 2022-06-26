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
		//O端用户
		user := guarded.Group("user")
		{
			user.Use(middleware.SuperAdmin()) //TODO:超级管理员才能用
			user.POST("create", controller.STDwrapperJSON(controller.AddUser))
			user.POST("update", controller.STDwrapperJSON(controller.UpdateUser))
			user.POST("search", controller.STDwrapperJSON(controller.SearchUser))
		}
		//估值
		valuate := guarded.Group("valuate")
		{
			valuate.Any("", controller.ForwardEnterpriseRequest)
			valuate.Any("*url", controller.ForwardEnterpriseRequest)
		}
		//企业
		enterprise := guarded.Group("enterprise")
		{
			enterprise.POST("", controller.ForwardEnterpriseRequest)
			enterprise.Any("*url", controller.ForwardEnterpriseRequest)
		}
		//机构
		group := guarded.Group("group")
		{
			group.Any("", controller.FowardGroupRequest)
			group.Any("*url", controller.FowardGroupRequest)
		}
		//审核
		audit := guarded.Group("audit")
		{
			audit.Any("", controller.ForwardEnterpriseRequest)
			audit.Any("*url", controller.ForwardEnterpriseRequest)
		}
	}
	//区分网关和业务侧404
	router.NoRoute(func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("Gateway 404"))
	})
	return
}
