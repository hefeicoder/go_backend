package server

import (
	"context"

	dsp "demo/proto/demo_service_proto"
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
	for _, v := range req.GetKeys() {
		sum += int64(v)
	}
	return &dsp.DemoResponse{
		Answer: sum,
	}, nil
}
