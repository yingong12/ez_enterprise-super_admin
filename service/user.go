package service

import "super_admin/repository"

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
