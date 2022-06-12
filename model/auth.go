package model

type AuthStatus struct {
	UID         string `json:"uid" example:"u_12345678901"`          //b端用户id
	AccessToken string `json:"access_token" example:"a_12345678901"` //a端用户id
}

//t_b_user
type User struct {
	UID      string `gorm:"column:uid" json:"uid"`
	Username string `gorm:"column:username" json:"username"`
	Name     string `gorm:"column:name" json:"name"`
	Password string `gorm:"column:pswd" json:"password"`
	Role     uint8  `gorm:"column:role" json:"role"`
	Phone    string `gorm:"column:phone" json:"phone"`
	State    uint8  `gorm:"column:state" json:"state"`
}

func (usr User) Table() string {
	return "t_user"
}
