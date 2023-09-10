package log

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"test-gin-mall/pkg/utils/es"
	"time"
)

var LogrusObj *logrus.Logger

func InitLog() {
	if LogrusObj != nil {
		src, _ := setOutputFile()
		LogrusObj.Out = src

		return
	}

	logger := logrus.New()
	src, _ := setOutputFile()

	logger.Out = src

	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2016-01-02 15:12:00",
	})

	hook := es.EsHookLog()
	logger.AddHook(hook)
	LogrusObj = logger
}

func setOutputFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""

	// 获取当前路径，拼接指定路径
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs"
	}

	// 获取当前路径文件信息
	_, err := os.Stat(logFilePath)

	// 判断是否存在
	if os.IsNotExist(err) {
		// 创建文件夹
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}

	logFileName := now.Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}

	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return src, err
}
