package api

import (
	"github.com/gin-gonic/gin"
	"test-gin-mall/pkg/utils/log"
	"test-gin-mall/types"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.UserRegisterReq
		if err := context.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
		}
	}
}
