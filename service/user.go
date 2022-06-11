package service

import (
	"super_admin/http/request"
	"super_admin/model"
	"super_admin/providers"
	"super_admin/repository"
)

func UpdateUser(UID, Username, Password, Phone string, Role, State uint8) (rows int64, err error) {
	/*
		更新状态， 清除登录态
	*/
	rows, err = repository.UpdateUser(UID, Username, Password, Phone, Role, State)
	if rows == 0 {
		return
	}
	//TODO:这里怎么能拿到token?
	// repository.DelAuthStatus()
	//clear login status
	return
}

var dict = []string{
	"username", "name", "uid",
}

func SearchUser(Filters []request.Filter, page, pageSize int) (res []model.User, err error) {
	//
	whereMap := map[string]interface{}{}
	for _, filter := range Filters {
		if filter.Type < 0 || filter.Type >= len(dict) {
			continue
		}
		whereMap[dict[filter.Type]] = filter.Value
	}
	tx := providers.DBAccount.Table(model.User{}.Table())

	for k, v := range whereMap {
		tx.Where(k, v)
	}
	tx.Find(&res)
	err = tx.Error
	return
}
