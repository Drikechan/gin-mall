package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	conf "test-gin-mall/config"
	"test-gin-mall/pkg/e"
	"test-gin-mall/pkg/utils/ctl"
)

func ErrorResponse(ctx *gin.Context, err error) *ctl.TrackErrorResponse {
	if v, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range v {
			field := conf.T(fmt.Sprintf("Field.%s", fieldError.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", fieldError.Tag))
			return ctl.RespError(ctx, err, fmt.Sprintf("%s%s", field, tag))
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return ctl.RespError(ctx, err, "Json类型不匹配")
	}
	return ctl.RespError(ctx, err, err.Error(), e.ERROR)
}
