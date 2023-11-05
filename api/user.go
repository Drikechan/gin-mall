package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"test-gin-mall/consts"
	"test-gin-mall/pkg/utils/ctl"
	"test-gin-mall/pkg/utils/log"
	"test-gin-mall/services"
	"test-gin-mall/types"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.UserRegisterReq
		if err := context.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		if req.Key == "" || len(req.Key) != consts.UserKeyLen {
			err := errors.New("key长度必须为6")
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}

		l := services.GetUserSrv()
		resp, err := l.UserRegister(context, &req)

		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
	}
}

func UserLoginHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.UserRegisterReq
		if err := context.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusBadRequest, ErrorResponse(context, err))
			return
		}

		l := services.GetUserSrv()
		resp, err := l.UserLogin(context.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusInternalServerError, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))

	}
}
