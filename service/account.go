package service

import (
	"super_admin/model"
	"super_admin/providers"
)

func GetUserAppID(uid string) (appID string, err error) {
	usr := model.User{}
	tx := providers.DBAccount.Table(usr.Table())
	tx.
		Select("app_id").
		Where("uid", uid).
		Find(&usr)
	appID = usr.AppID
	err = tx.Error
	return
}
