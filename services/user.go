package services

import (
	"context"
	"errors"
	"sync"
	conf "test-gin-mall/config"
	"test-gin-mall/consts"
	"test-gin-mall/pkg/utils/log"
	"test-gin-mall/repository/db/dao"
	"test-gin-mall/repository/db/model"
	"test-gin-mall/types"
)

var UserSrvIns *UserSrv

type UserSrv struct {
}

var UserSrvOnce sync.Once

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (s *UserSrv) UserRegister(ctx context.Context, req *types.UserRegisterReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(req.UserName)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	if exist {
		err = errors.New("用户已经存在")
		return
	}

	user := &model.User{
		UserName: req.UserName,
		NickName: req.NickName,
		Status:   model.Active,
		Money:    consts.UserInitMoney,
	}

	if err = user.SetPassword(req.Password); err != nil {
		log.LogrusObj.Error(err)
		return
	}

	money, err := user.EncryptMoney(req.Key)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	user.Money = money
	user.Avatar = consts.UserDefaultAvatarLocal
	if conf.Config.System.UploadModel == consts.UploadModalOss {
		user.Avatar = consts.UserDefaultAvatarLocal
	}
	err = userDao.CreateUser(user)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	return
}
