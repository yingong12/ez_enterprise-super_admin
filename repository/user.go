package repository

import (
	"log"
	"super_admin/model"
	"super_admin/providers"
)

//TODO:这里改为事务
func UpdateUser(UID, Username, Password, Phone string, Role, State uint8) (rows int64, err error) {
	env := model.User{
		Username: Username,
		Password: Password,
		Phone:    Phone,
		Role:     Role,
	}

	log.Println(UID, Username, Password, Phone, Role, State)
	tx := providers.DBAccount.Table(env.Table())
	tx.
		Where("uid", UID).
		Updates(env)

	tx.Where("uid", UID).
		Update("state", State)
	rows = tx.RowsAffected
	err = tx.Error
	return
}
