package middleware

import (
	"github.com/gin-gonic/gin"
	"test-gin-mall/consts"
	"test-gin-mall/pkg/e"
	"test-gin-mall/pkg/utils/ctl"
	util "test-gin-mall/pkg/utils/jwt"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code int
		code = e.SUCCESS
		accessToken := context.GetHeader("access_header")
		refreshToken := context.GetHeader("refresh_token")
		if accessToken == "" {
			code = e.InvalidParams
			context.JSON(code, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "Token不能为空",
			})
			context.Abort()
			return
		}

		newAccessToken, newRefreshToken, err := util.ParseRefreshToken(accessToken, refreshToken)
		if err != nil {
			code = e.ErrorAuthCheckFail
			context.JSON(code, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "Token验证错误",
			})
			context.Abort()
			return
		}

		claims, err := util.ParseToken(accessToken)
		if err != nil {
			code = e.ErrorAuthCheckFail
			context.JSON(code, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "Token验证错误",
			})
			context.Abort()
			return
		}
		SetToken(context, newAccessToken, newRefreshToken)
		context.Request = context.Request.WithContext(ctl.NewContent(context.Request.Context(), &ctl.UserInfo{Id: claims.ID}))
		context.Next()

	}
}

func SetToken(c *gin.Context, accessToken, refreshToken string) {
	secure := IsHttps(c)
	c.Header(consts.AssessTokenHeader, accessToken)
	c.Header(consts.RefreshTokenHeader, refreshToken)
	c.SetCookie(consts.AssessTokenHeader, accessToken, consts.MaxAge, "/", "", secure, true)
	c.SetCookie(consts.RefreshTokenHeader, refreshToken, consts.MaxAge, "/", "", secure, true)
}

func IsHttps(c *gin.Context) bool {
	if c.GetHeader(consts.HeaderForwardProto) == "https" || c.Request.TLS != nil {
		return true
	}
	return false
}
