package ctl

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
	"test-gin-mall/consts"
	"test-gin-mall/pkg/e"
)

type Response struct {
	Msg     string      `json:"msg"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
	TrackId string      `json:"track_id"`
}

type TrackErrorResponse struct {
	Response
	TrackId string `json:"track_id"`
}

func RespSuccess(ctx *gin.Context, data interface{}, code ...int) *Response {
	trackId, _ := getTrackIdFormatCtx(ctx)
	status := e.SUCCESS
	if code != nil {
		status = code[0]
	}
	if data == nil {
		data = "操作成功"
	}

	return &Response{
		Data:    data,
		Status:  status,
		TrackId: trackId,
		Msg:     e.GetMsg(status),
	}
}

func RespError(ctx *gin.Context, err error, data interface{}, code ...int) *TrackErrorResponse {
	trackId, _ := getTrackIdFormatCtx(ctx)
	status := e.ERROR
	if code != nil {
		status = code[0]
	}
	if data == nil {
		data = "操作失败"
	}
	return &TrackErrorResponse{
		Response: Response{
			Msg:     e.GetMsg(status),
			Data:    data,
			Status:  status,
			TrackId: trackId,
			Error:   err.Error(),
		},
		TrackId: trackId,
	}
}

func getTrackIdFormatCtx(ctx *gin.Context) (trackId string, err error) {
	spanCtxInterface, _ := ctx.Get(consts.SpanCTX)
	str := fmt.Sprintf("%v", spanCtxInterface)
	re := regexp.MustCompile(`([0-9a-fA-F]{16})`)

	match := re.FindStringSubmatch(str)
	if len(match) > 0 {
		return match[1], nil
	}
	return "", errors.New("获取trackId错误")
}
