package middleware

import (
	"net/http"
	"super_admin/http/buz_code"
	"super_admin/service"

	"github.com/gin-gonic/gin"
)

//登录态信息, 请求context公用
type AuthInfo struct {
	UID  string `json:"uid"`
	Role uint8  `json:"role"`
}

//从请求体内拿登录态信息
func GetAuthInfo(ctx *gin.Context) (info *AuthInfo, ok bool) {
	d, ok := ctx.Get("auth_info")
	info = d.(*AuthInfo)
	return
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//
		token := c.Request.Header.Get("o_access_token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": buz_code.CODE_AUTH_FAILED,
				"msg":  "缺少token",
			})
			c.Abort()
			return
		}
		uid, err := openAuth(token)
		if err != nil {
			//网络错误
			c.JSON(http.StatusOK, gin.H{
				"code": buz_code.CODE_SERVER_ERROR,
				"msg":  "服务器内部错误" + err.Error(),
			})
			c.Abort()
			return
		}
		//鉴权失败
		if uid == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": buz_code.CODE_AUTH_FAILED,
				"msg":  "未登录",
			})
			c.Abort()
		}
		info := &AuthInfo{
			UID: uid,
		}
		c.Set("auth_info", info)
		c.Next()
	}
}

//HeaderInjector injects the header to the request
func HeaderInjector() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
func SuperAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authInfo, ok := ctx.Get("auth_info")
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{
				"code": buz_code.CODE_AUTH_FAILED,
				"msg":  "未登录",
			})
			ctx.Abort()
			return
		}
		//超级管理员
		if authInfo.(*AuthInfo).Role == 1 {
			ctx.JSON(http.StatusOK, gin.H{
				"code": buz_code.CODE_UNAUTHORIZED,
				"msg":  "非超级管理员无权限访问",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

type BaseRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type RspAuthInfo struct {
	BaseRsp
	Data struct {
		UID       string `json:"uid"`       //b端用户id
		AppID     string `json:"app_id"`    //appID
		ExpiresAt string `json:"expire_at"` //过期时间

	} `json:"data"`
}

func openAuth(token string) (uid string, err error) {
	info, err := service.Check(token)
	uid = info.UID
	return
}
