package request

type UpdateUser struct {
	UID      string  `json:"uid" example:"u_12345678901"`
	Username string  `json:"username" example:"bajie"`
	Password string  `json:"pwsd" example:"asda"`
	Phone    string  `json:"phone" example:""`
	State    uint8   `json:"state" example:""`
	Roles    []uint8 `json:"roles" binding:"dive,required,min=1,max=5" example:"1,2,3"` //角色列表
}

type SignUpUsernameRequest struct {
	BaseRequest
	Username string  `json:"username" example:"zhuyan"`                                 //用户名
	Password string  `json:"password" exmample:"123456@Zhuyan"`                         //密码，需要包含大小写数字和特殊字符
	Roles    []uint8 `json:"roles" binding:"dive,required,min=1,max=5" example:"1,2,3"` //角色列表
}
