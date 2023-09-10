package es

import (
	"fmt"
	"github.com/CocaineCong/eslogrus"
	"github.com/elastic/go-elasticsearch"
	"github.com/sirupsen/logrus"
	"log"
	conf "test-gin-mall/config"
)

var EsClient *elasticsearch.Client

func InitEs() {
	esConfig := conf.Config.Es
	esConn := fmt.Sprintf("http://%s:%s", esConfig.EsHost, esConfig.EsPort)
	config := elasticsearch.Config{
		Addresses: []string{esConn},
	}

	client, err := elasticsearch.NewClient(config)
	if err != nil {
		return
	}

	EsClient = client

}

func EsHookLog() *eslogrus.ElasticHook {
	esConfig := conf.Config.Es
	hook, err := eslogrus.NewElasticHook(EsClient, esConfig.EsHost, logrus.DebugLevel, esConfig.EsIndex)

	if err != nil {
		log.Panic(err)
	}
	return hook
}
