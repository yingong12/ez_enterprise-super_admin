package controller

import (
	"fmt"
	"super_admin/http/buz_code"
	"super_admin/http/request"
	"super_admin/http/response"
	"super_admin/library/env"
	"super_admin/service"

	"github.com/gin-gonic/gin"
)

//Create	登录态
//@Summary	登录态校验
//@Description	登录态校验
//@Tags	登录态校验
//@Produce	json
//@Param  b_access_token header string true "b端用户token"
//@Success 200 {object} model.AuthStatus
//@Router	/auth/check [get]
func Check(ctx *gin.Context) (res *STDResponse, err error) {
	res = &STDResponse{}
	token := ctx.GetHeader(env.GetStringVal("TOKEN_KEY"))
	//参数校验
	if token == "" {
		res.Code = buz_code.CODE_INVALID_ARGS
		res.Msg = "缺少token"
		return
	}
	//buz逻辑
	authInfo, err := service.Check(token)
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		res.Msg = err.Error()
		return
	}
	//为nil时没有登录
	if authInfo.UID == "" {
		res.Code = buz_code.CODE_AUTH_FAILED
		res.Msg = "未登录"
		return
	}
	res.Code = buz_code.CODE_OK
	res.Msg = "ok"
	res.Data = authInfo
	return
}

// /Create	注册登录
//@Summary	用户名登录
//@Description	用户名登录
//@Tags
//@Produce	json
//@Param xxx body request.SignInUsernameRequest  false "注释"
//@Success 200 {object} response.SignInUsernameRsp
//@Router	/signin/username [post]
func SignInUsername(ctx *gin.Context) (res *STDResponse, err error) {
	req := request.SignInUsernameRequest{}
	res = &STDResponse{}
	if err = ctx.BindJSON(&req); err != nil {
		res.Code = buz_code.CODE_INVALID_ARGS
		res.Msg = fmt.Sprintf("invalid params %s\n", err.Error())
		return
	}
	accessToken, uid, err := service.SignInUsername(req.Username, req.Password)
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		res.Msg = "SERVER ERROR"
		return
	}
	if accessToken == "" {
		res.Code = buz_code.CODE_USERNAME_PSWD_NOT_MATCH
		res.Msg = "用户名密码不匹配"
		return
	}
	//
	res.Code = buz_code.CODE_OK
	res.Msg = "ok"
	res.Data = response.SignInUsernameRsp{
		UID:         uid,
		AccessToken: accessToken,
	}
	return
}
