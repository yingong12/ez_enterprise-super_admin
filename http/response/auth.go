package response

type SignInUsernameRsp struct {
	AccessToken string `json:"o_access_token" example:"o_u_uasdasd"` //b端用户token
	UID         string `json:"uid" exmample:"zhuyan911"`             //用户ID
}

type SignUpRsp struct {
	AccessToken string `json:"o_access_token" example:"o_u_uasdasd"` //b端用户token
	UID         string `json:"uid" exmample:"zhuyan911"`             //用户ID
}
