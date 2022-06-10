package response

type SignInUsernameRsp struct {
	AccessToken string `json:"b_access_token" example:"b_u_uasdasd"` //b端用户token
	UID         string `json:"uid" exmample:"zhuyan911"`             //用户ID
}

type SignUpRsp struct {
	AccessToken string `json:"b_access_token" example:"b_u_uasdasd"` //b端用户token
	UID         string `json:"uid" exmample:"zhuyan911"`             //用户ID
}
