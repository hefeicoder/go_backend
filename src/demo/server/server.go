package server

import (
	"context"
	"fmt"

	dsp "demo/proto/demo_service_proto"
	"github.com/valyala/fasthttp"
	"google.golang.org/protobuf/proto"
	"logger"
)

const (
	http_response_content_type = "application/octet-stream"
)

// DemoService
type DemoService struct {
	// Service state variables
}

// create a new DemoService
func New(ctx context.Context) (*DemoService, error) {
	return &DemoService{}, nil
}

// GetFeature
func (*DemoService) GetFeature(ctx context.Context, req *dsp.DemoRequest) (*dsp.DemoResponse, error) {
	var sum int64 = 0
	logger.Infof("GetFeature called: %v", req.String())

	for _, v := range req.GetKeys() {
		sum += int64(v)
	}

	return &dsp.DemoResponse{
		Answer: sum,
	}, nil
}

func (demo *DemoService) GetFeatureHTTP(httpctx *fasthttp.RequestCtx) {
	req := &dsp.DemoRequest{}

	if err := proto.Unmarshal(httpctx.PostBody(), req); err != nil {
		handleError(httpctx, fasthttp.StatusBadRequest, "bad request: error=%v", err)
		return
	}

	resp, err := demo.GetFeature(context.Background(), req)

	if err != nil {
		handleError(httpctx, fasthttp.StatusInternalServerError, "failed to process request: error=%v", err)
	}

	payload, _ := proto.Marshal(resp)
	httpctx.SetStatusCode(fasthttp.StatusOK)
	httpctx.SetContentType(http_response_content_type)
	httpctx.Write(payload)
}

func (*DemoService) HealthHTTP(httpctx *fasthttp.RequestCtx) {
	httpctx.SetStatusCode(fasthttp.StatusOK)
	httpctx.SetBodyString("OK")
}

func handleError(httpctx *fasthttp.RequestCtx, code int, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args)
	httpctx.Error(msg, code)
	logger.Errorf(format, msg)
}
