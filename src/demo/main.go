package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/valyala/fasthttp"
	"logger"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"demo/server"
	"github.com/buaazp/fasthttprouter"
)

const ()

var (
	httpPort = flag.Int("http_port", 8080, "http port")
	grpcPort = flag.Int("grpc_port", 8080, "grpc port")

	enableHTTP = flag.Bool("enable_http", true, "run as http server")
)

func main() {
	fmt.Printf("Demo Service from (Go version: %s)\n", runtime.Version())
	ctx, cxl := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	service, err := server.New(ctx)
	if err != nil {
		fmt.Printf("Server create error")
		return
	}

	if *enableHTTP {
		if err := startHttpServer(ctx, cxl, service); err != nil {
			logger.Error("failed to start http server, error=%v", err)
			return
		}
	}

	select {
	case s := <-c:
		logger.Infof("signal %v received", s)
		cxl()
	case <-ctx.Done():
	}
	logger.Info("Server exited main")
}

func startHttpServer(ctx context.Context, cxl context.CancelFunc, srv *server.DemoService) error {
	router := fasthttprouter.New()

	router.GET("/ready", srv.HealthHTTP)
	router.POST("/get_feature", srv.GetFeatureHTTP)

	s := &fasthttp.Server{
		Handler: router.Handler,
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *httpPort))
	if err != nil {
		logger.Errorf("HTTP server fail to listen on %v error=%v", *httpPort, err)
	}

	go func() {
		if err := s.Serve(l); err != nil {
			logger.Errorf("HTTP server error=%v", err)
		}
		cxl()
		logger.Info("HTTP server stopped")
	}()
	go func() {
		<-ctx.Done()
		logger.Info("stopping HTTP server...")
		s.Shutdown()
	}()

	return nil
}
