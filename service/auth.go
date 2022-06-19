package service

import (
	"super_admin/library/env"
	"super_admin/model"
	"super_admin/providers"
	"super_admin/repository"
	"super_admin/utils"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func Check(accessToken string) (authInfo *model.AuthStatus, err error) {
	authInfo, err = repository.GetAuthStatus(accessToken)
	//空key 未登录
	if err == redis.Nil {
		err = nil
	}
	return
}

func SignUpUsername(username, pswd string, role uint8) (accessToken, uid string, err error) {
	tx := providers.DBAccount.Begin()
	defer func() {
		//网络错误或者已被注册
		if err != nil || uid == "" {
			tx.Rollback()
			return
		}
		//没问题后注册用户并登录
		accessToken, err = setLoginStatus(uid, role)
		tx.Commit()
	}()
	//没被注册才继续
	if err = repository.GetUserByKey("username", username); err != gorm.ErrRecordNotFound {
		return
	}
	uid, err = repository.InsertUser(username, "", pswd, role)
	return
}

//TODO:这里要做同时valid token数量的限制
func SignInUsername(username, pswd string) (accessToken, uid string, err error) {
	m := map[string]interface{}{
		"username": username,
		"pswd":     pswd,
	}
	usr, err := repository.GetUserByKeys(m)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
		return
	}
	uid = usr.UID
	role := usr.Role
	accessToken, err = setLoginStatus(uid, role)
	return

}

//设置登录态
func setLoginStatus(uid string, role uint8) (accessToken string, err error) {
	accessToken = utils.GenerateOAccessToken()
	err = repository.SetLoginStatus(uid, role, accessToken)
	return
}

//校验验证码
func checkVerifyCode(verifyCode, phone string) (ok bool, err error) {
	serverCode, err := repository.GetSMSEntry(env.GetStringVal("KEY_PREFIX_SMS") + phone)
	//验证码校验不通过
	if err != nil {
		if err == redis.Nil {
			//错误原因设为验证码错误
			err = nil
		}
		return
	}
	ok = serverCode == verifyCode
	return
}
