package upload

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	conf "test-gin-mall/config"
	util "test-gin-mall/pkg/utils/log"
)

func ProductUploadToLocalStatic(file multipart.File, bossId uint, productName string) (filePath string, err error) {
	bId := strconv.Itoa(int(bossId))
	basePath := "." + conf.Config.PhotoPath.PhotoPath + "boss" + bId + "/"
	if !IsExistDir(basePath) {
		CreateDir(basePath)
	}
	productPath := fmt.Sprintf("%s%s.jpg", basePath, productName)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		util.LogrusObj.Error(err)
		return "", err
	}

	err = ioutil.WriteFile(productPath, content, 0666)
	if err != nil {
		util.LogrusObj.Error(err)
		return "", err
	}

	return fmt.Sprintf("boss%s/%s.jpg", bId, productName), err

}

func IsExistDir(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}

func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
