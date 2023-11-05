package consts

import "time"

const UserInitMoney = "10000"

const UserKeyLen = 6

const (
	UserDefaultAvatarLocal = "avatar.JPG"
	UserDefaultAvatarOss   = "http://q1.qlogo.cn/g?b=qq&nk=294350394&s=640"
)

const (
	AccessTokenExpireDuration        = 24 * time.Hour
	AccessRefreshTokenExpireDuration = 10 * 24 * time.Hour
)
