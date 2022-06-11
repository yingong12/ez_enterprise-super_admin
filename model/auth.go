package model

type AuthStatus struct {
	UID         string `json:"uid" example:"u_12345678901"`          //b端用户id
	AccessToken string `json:"access_token" example:"a_12345678901"` //a端用户id
}

//t_b_user
type User struct {
	UID      string `gorm:"column:uid"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:pswd"`
	Role     uint8  `gorm:"column:role"`
	Phone    string `gorm:"column:phone"`
	State    uint8  `gorm:"column:state"`
}

func (usr User) Table() string {
	return "t_user"
}
