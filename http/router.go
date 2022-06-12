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
		/**
		1. 登录态校验  token-> code 0:uid,app_id  code 1: 过期   code 2:wrong token code 3:token missing
		2. 注册： username,pswd,phone(带验证码) -> token
		3. 手机登录： phone#, veriCode -> token
		*/
		auth.GET("/check", controller.STDwrapperJSON(controller.Check)) //校验登录态

		auth.POST("/signin/username", controller.STDwrapperJSON(controller.SignInUsername)) //用户名注册
	}
	//用户
	user := router.Group("user")
	{
		user.POST("create", controller.STDwrapperJSON(controller.AddUser))
		user.POST("update", controller.STDwrapperJSON(controller.UpdateUser))
		user.POST("search", controller.STDwrapperJSON(controller.SearchUser))
	}
	return
}
