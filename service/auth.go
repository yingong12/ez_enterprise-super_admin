package service

import (
	"log"
	"super_admin/http/buz_code"
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
		log.Println(accessToken, err)
		err = nil
	}
	return
}

func SignUpUsername(username, pswd string) (accessToken, uid string, err error) {
	tx := providers.DBAccount.Begin()
	defer func() {
		//网络错误或者已被注册
		if err != nil || uid == "" {
			tx.Rollback()
			return
		}
		//没问题后注册用户并登录
		accessToken, err = setLoginStatus(uid, "")
		tx.Commit()
	}()
	//没被注册才继续
	if err = repository.GetUserByKey("username", username); err != gorm.ErrRecordNotFound {
		return
	}
	uid, err = repository.InsertUser(username, "", pswd)
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
	appID := usr.AppID
	accessToken, err = setLoginStatus(uid, appID)
	return

}

//设置登录态
func setLoginStatus(uid, appID string) (accessToken string, err error) {
	accessToken = utils.GenerateAccessToken()
	err = repository.SetLoginStatus(uid, appID, accessToken)
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
func SinUpSMS(phone, password, verifyCode string) (accessToken, uid string, err error, buzCode buz_code.Code) {
	//校验验证码
	ok, err := checkVerifyCode(verifyCode, phone)
	if err != nil {
		return
	}
	if !ok {
		buzCode = buz_code.CODE_BAD_SMS_CODE
		return
	}
	//
	tx := providers.DBAccount.Begin()
	defer func() {
		//网络错误或者已被注册
		if err != nil {
			tx.Rollback()
			return
		}
		//已注册
		if uid == "" {
			buzCode = buz_code.CODE_USER_ALREADY_EXISTS
			tx.Rollback()
			return
		}
		//没问题后注册用户并登录
		accessToken, err = setLoginStatus(uid, "")
		tx.Commit()
	}()
	//没被注册才继续
	if err = repository.GetUserByKey("phone", phone); err != gorm.ErrRecordNotFound {
		return
	}
	uid, err = repository.InsertUser("", phone, password)
	return
}

func SignInSMS(phone, password, verifyCode string) (accessToken, uid string, err error, buzCode buz_code.Code) {
	//校验验证码
	ok, err := checkVerifyCode(verifyCode, phone)
	if err != nil {
		return
	}
	if !ok {
		buzCode = buz_code.CODE_BAD_SMS_CODE
		return
	}
	//业务逻辑
	m := map[string]interface{}{
		"phone": phone,
		"pswd":  password,
	}
	usr, err := repository.GetUserByKeys(m)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			buzCode = buz_code.CODE_USERNAME_PSWD_NOT_MATCH
			err = nil
		}
		return
	}
	uid = usr.UID
	accessToken, err = setLoginStatus(uid, usr.AppID)
	return
}
