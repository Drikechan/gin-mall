package types

type UserRegisterReq struct {
	NickName string `from:"nick_name" json:"nickName"`
	UserName string `from:"user_name" json:"userName"`
	Password string `from:"password" json:"password"`
	Key      string `from:"key" json:"key"`
}
