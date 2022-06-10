package request

type SignUpUsernameRequest struct {
	BaseRequest
	Username string `json:"username" example:"zhuyan"`         //用户名
	Password string `json:"password" exmample:"123456@Zhuyan"` //密码，需要包含大小写数字和特殊字符
}

type SignInUsernameRequest struct {
	BaseRequest
	Username string `json:"username" example:"zhuyan"`         //用户名
	Password string `json:"password" exmample:"123456@Zhuyan"` //密码，需要包含大小写数字和特殊字符
}

type SignUpSMSRequest struct {
	BaseRequest
	Phone      string `json:"phone" example:"18391025131"`       //用户名
	Password   string `json:"password" exmample:"123456@Zhuyan"` //密码，需要包含大小写数字和特殊字符
	VerifyCode string `json:"verify_code" exmample:"637522"`     //验证码
}
type SignInSMSRequest struct {
	BaseRequest
	Phone      string `json:"phone" example:"18391025131"`       //用户名
	Password   string `json:"password" exmample:"123456@Zhuyan"` //密码，需要包含大小写数字和特殊字符
	VerifyCode string `json:"verify_code" exmample:"637522"`     //验证码
}
