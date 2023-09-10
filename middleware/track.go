package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"test-gin-mall/pkg/utils/track"
)

const SpanCTX = "span-ctx"

func Jaeger() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取请求头uber-trace-id
		traceId := context.GetHeader("uber-trace-id")
		var span opentracing.Span
		if traceId != "" {
			var err error
			span, err = track.GetParentSpan(context.FullPath(), traceId, context.Request.Header)
			if err != nil {
				return
			}
		} else {
			span = track.StartSpan(opentracing.GlobalTracer(), context.FullPath())
		}

		defer span.Finish()
		context.Set(SpanCTX, opentracing.ContextWithSpan(context, span))
		context.Next()
	}
}
