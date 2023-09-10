package track

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

func StartSpan(tracer opentracing.Tracer, name string) opentracing.Span {
	// 创建最顶级span
	span := tracer.StartSpan(name)
	return span
}

func GetParentSpan(spanName string, traceId string, header http.Header) (opentracing.Span, error) {
	carrier := opentracing.HTTPHeadersCarrier{}
	carrier.Set("uber-trace-id", traceId)

	tracer := opentracing.GlobalTracer()
	//从 HTTP 头中提取追踪信息，以获取先前创建的父级 span 的上下文
	wireContext, err := tracer.Extract(
		opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))

	parentSpan := opentracing.StartSpan(spanName, ext.RPCServerOption(wireContext))

	if err != nil {
		return nil, err
	}

	return parentSpan, err

}
