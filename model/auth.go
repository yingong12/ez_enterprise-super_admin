package model

type AuthStatus struct {
	UID       string `json:"uid" example:"u_12345678901"` //b端用户id
	AppID     string `json:"app_id"`
	ExpiresAt string `json:"expire_at" example:"2022-05-16 23:00:00"` //过期时间

}

//t_b_user
type User struct {
	AppID    string `gorm:"column:app_id"`
	UID      string `gorm:"column:uid"`
	Password string `gorm:"column:pswd"`
	Username string `gorm:"column:username"`
	Phone    string `gorm:"column:phone"`
	Email    string `gorm:"column:email"`
	Area     string `gorm:"column:area"`
	NickName string `gorm:"column:nickname"`
}

func (usr User) Table() string {
	return "t_b_user"
}
