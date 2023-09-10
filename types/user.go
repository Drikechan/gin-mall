package types

type UserRegisterReq struct {
	NickName string `from:"nick_name" json:"nick_name"`
	UserName string `from:"user_name" json:"user_name"`
	Password string `from:"password" json:"password"`
	Key      string `from:"key" json:"key"`
}
