package dao

import "test-gin-mall/repository/db/model"

func migrate() (err error) {
	err = DB.AutoMigrate(&model.User{})

	return
}
