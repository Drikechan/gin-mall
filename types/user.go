package types

type UserRegisterReq struct {
	NickName string `from:"nick_name" json:"nickName"`
	UserName string `from:"user_name" json:"userName"`
	Password string `from:"password" json:"password"`
	Key      string `from:"key" json:"key"`
}

type UserInfoResp struct {
	ID       uint   `json:"id"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
	NickName string `json:"nickName"`
	Status   string `json:"status"`
	Type     string `json:"type"`
	Email    string `json:"email"`
	CreateAt string `json:"createAt"`
}

type UserTokenResp struct {
	AccessToken  string      `json:"accessToken"`
	RefreshToken string      `json:"refreshToken"`
	User         interface{} `json:"user"`
}
