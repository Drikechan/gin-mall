package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"test-gin-mall/api"
	"test-gin-mall/docs"
	"test-gin-mall/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// 创建了一个基于 Cookie 的会话存储
	store := cookie.NewStore([]byte("something-very-secret"))

	// 将会话管理中间件添加到 Gin 路由器
	r.Use(sessions.Sessions("my-session", store))
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.Use(middleware.Cors(), middleware.Jaeger())
	v1 := r.Group("/api/v1")

	{
		v1.POST("/user/register", api.UserRegisterHandler())
		v1.POST("/user/login", api.UserLoginHandler())
		v1.GET("/product/list", api.ListProductsHandler())

		authed := v1.Group("/")
		authed.Use(middleware.AuthMiddleWare())
		{
			authed.POST("product/create", api.CreateProductHandler())
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
