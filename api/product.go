package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test-gin-mall/consts"
	"test-gin-mall/pkg/utils/ctl"
	"test-gin-mall/pkg/utils/log"
	"test-gin-mall/services"
	"test-gin-mall/types"
)

func CreateProductHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.CreateProductResp
		if err := context.ShouldBind(&req); err != nil {
			log.LogrusObj.Error(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}

		form, _ := context.MultipartForm()
		files := form.File["images"]
		fmt.Println(files)
		//l := services.GetProductSrv()

	}
}

func ListProductsHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req types.ProductListReq
		if err := context.ShouldBind(&req); err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusOK, ErrorResponse(context, err))
			return
		}

		if req.PageSize == 0 {
			req.PageSize = consts.BaseProductListPageSize
		}
		l := services.GetProductSrv()
		resp, err := l.ListProduct(context.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			context.JSON(http.StatusInternalServerError, ErrorResponse(context, err))
			return
		}
		context.JSON(http.StatusOK, ctl.RespSuccess(context, resp))
	}
}
