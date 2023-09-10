package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"strings"
	conf "test-gin-mall/config"
	"time"
)

var DB *gorm.DB

func InitMysql() {
	mysqlConfig := conf.Config.Mysql["default"]
	pathRead := strings.Join([]string{mysqlConfig.UserName, ":", mysqlConfig.Password, "@tcp(", mysqlConfig.DbHost, ":", mysqlConfig.DbPort, ")/", mysqlConfig.DbName, "?charset=" + mysqlConfig.Charset + "&parseTime=true"}, "")
	pathWrite := strings.Join([]string{mysqlConfig.UserName, ":", mysqlConfig.Password, "@tcp(", mysqlConfig.DbHost, ":", mysqlConfig.DbPort, ")/", mysqlConfig.DbName, "?charset=" + mysqlConfig.Charset + "&parseTime=true"}, "")

	var ormLogger logger.Interface

	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	//
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       pathRead, // DSN data source name
		DefaultStringSize:         256,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{ //你在 GORM 中定义一个模型结构体时，GORM 会根据结构体的名称来自动映射到数据库表的名称。
			SingularTable: true,
		},
	})

	fmt.Println(db, err)

	if err != nil {
		fmt.Println("数据库打开错误")
		panic(err)
	}
	//
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	DB = db

	_ = DB.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{mysql.Open(pathRead)},                         // 读操作
		Sources:  []gorm.Dialector{mysql.Open(pathWrite), mysql.Open(pathWrite)}, // 写操作
		Policy:   dbresolver.RandomPolicy{},                                      // 负载均衡
	}))

	DB = DB.Set("gorm:table_options", "charset=utf8mb4")

	err = migrate()
	if err != nil {
		fmt.Println("创建表失败")
		return
	}
}
