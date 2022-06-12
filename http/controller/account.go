package controller

import (
	"super_admin/http/buz_code"
	"super_admin/http/request"
	"super_admin/http/response"
	"super_admin/service"

	"github.com/gin-gonic/gin"
)

//BindPhone 绑定手机号
func BindPhone(ctx *gin.Context) {
	uid, _ := ctx.Params.Get("uid")
	ctx.Writer.Write([]byte(uid))
}

//Create	注册登录
//@Summary	用户名注册
//@Description	用户名注册
//@Tags
//@Produce	json
//@Param xxx body request.SignUpUsernameRequest false "注释"
//@Success 200 {object} response.SignUpRsp
//@Router	/signup/username [post]
func AddUser(ctx *gin.Context) (res STDResponse, err error) {
	/*
		username+pswd
	*/
	req := request.SignUpUsernameRequest{}
	if err = ctx.BindJSON(&req); err != nil {
		res.Code = buz_code.CODE_INVALID_ARGS
		res.Msg = err.Error()
		return
	}
	//默认角色为0
	if len(req.Roles) == 0 {
		req.Roles = []uint8{0}
	}
	accessToken, uid, err := service.SignUpUsername(req.Username, req.Password, req.Roles[0])
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		res.Msg = "SERVER ERROR"
		return
	}
	if accessToken == "" {
		res.Code = buz_code.CODE_USER_ALREADY_EXISTS
		res.Msg = "用户名已存在"
		return
	}
	//
	res.Code = buz_code.CODE_OK
	res.Msg = "ok"
	res.Data = response.SignUpRsp{
		UID:         uid,
		AccessToken: accessToken,
	}
	return
}

func UpdateUser(ctx *gin.Context) (res STDResponse, err error) {
	req := request.UpdateUser{}
	if err = ctx.BindJSON(&req); err != nil {
		res.Code = buz_code.CODE_INVALID_ARGS
		res.Msg = err.Error()
		return
	}
	role := uint8(127)
	if len(req.Roles) > 0 {
		role = req.Roles[0]
	}
	_, err = service.UpdateUser(req.UID, req.Username, req.Password, req.Phone, role, req.State)
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		res.Msg = err.Error()
		return
	}
	return
}

func SearchUser(ctx *gin.Context) (res STDResponse, err error) {
	req := request.SearchUser{}
	if err = ctx.BindJSON(&req); err != nil {
		res.Code = buz_code.CODE_INVALID_ARGS
		res.Msg = err.Error()
		return
	}
	list, err := service.SearchUser(req.Filters, req.Page, req.PageSize)
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		res.Msg = err.Error()
		return
	}
	res.Data = map[string]interface{}{
		"list":  list,
		"total": 0,
	}
	return
}
