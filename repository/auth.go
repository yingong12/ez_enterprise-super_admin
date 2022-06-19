package repository

import (
	"super_admin/library/env"
	"super_admin/model"
	"super_admin/providers"
	"time"

	"super_admin/utils"

	"gorm.io/gorm/clause"
)

func SetLoginStatus(uid string, role uint8, accessToken string) (err error) {
	prefixedToken := env.GetStringVal("KEY_PREFIX_O_TOKEN") + accessToken
	if cmd := providers.RedisClient.HMSet(prefixedToken, map[string]interface{}{
		"uid":  uid,
		"role": role,
	}); cmd.Err() != nil {
		return cmd.Err()
	}
	cmd := providers.RedisClient.Expire(accessToken, time.Hour*2)
	return cmd.Err()
}
func GetAuthStatus(token string) (as *model.AuthStatus, err error) {
	key := env.GetStringVal("KEY_PREFIX_O_TOKEN") + token
	res, err := providers.RedisClient.HGetAll(key).Result()
	if err != nil {
		return
	}
	as = &model.AuthStatus{}
	as.UID = res["uid"]
	as.AccessToken = token
	return
}
func DelAuthStatus(token string) (count int64, err error) {
	key := env.GetStringVal("KEY_PREFIX_O_TOKEN") + token
	count, err = providers.RedisClient.Del(key).Result()
	return
}
func GetUserByKeys(m map[string]interface{}) (usr model.User, err error) {
	usr = model.User{}
	tx := providers.DBAccount.Table(usr.Table())
	//AND连接
	for k, v := range m {
		tx = tx.Where(k, v)
	}
	tx.First(&usr)
	err = tx.Error
	return
}
func GetUserByKey(key string, val string) (err error) {
	usr := model.User{}
	tx := providers.DBAccount.Table(usr.Table()).
		//排它锁
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where(key, val).
		First(&usr)
	err = tx.Error
	return
}
func InsertUser(username, phone, pswd string, role uint8) (uid string, err error) {
	uid = utils.GenerateUID()
	usr := model.User{
		Username: username,
		Phone:    phone,
		Password: pswd,
		UID:      uid,
		Role:     role,
	}
	tb := usr.Table()
	tx := providers.DBAccount.
		Table(tb).
		Create(usr)
	err = tx.Error
	return
}
