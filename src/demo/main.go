package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"demo/server"
)

var (
	httpPort = flag.Int("http_port", 8080, "http port")
	grpcPort = flag.Int("grpc_port", 8080, "grpc port")

	enableHTTP = flag.Bool("enable_http", true, "run as http server")
)

func main() {
	fmt.Printf("Hello World from (Go version: %s)\n", runtime.Version())
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

		}
	}
}

func startHttpServer(ctx context.Context, cxl context.CancelFunc, srv *server.DemoService) error {

	return nil
}
