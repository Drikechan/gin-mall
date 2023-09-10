package main

import (
	conf "test-gin-mall/config"
	"test-gin-mall/pkg/utils/es"
	util "test-gin-mall/pkg/utils/log"
	"test-gin-mall/repository/db/dao"
	"test-gin-mall/routes"
)

func main() {
	loading()
	r := routes.NewRouter()
	r.Run(conf.Config.System.Port)
}

func loading() {
	conf.InitConfig()
	dao.InitMysql()
	util.InitLog()
	es.InitEs()
}
