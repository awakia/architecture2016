package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gengo/grpc-gateway/runtime"
	"github.com/golang/glog"
	"golang.org/x/net/context"

	gw "github.com/awakia/architecture2016/src/rpc"
)

var (
	endpoint = flag.String("endpoint", "localhost:9090", "endpoint of the service")
	port     = flag.Int("port", 8080, "The reverse proxy port")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	err := gw.RegisterUserServiceHandlerFromEndpoint(ctx, mux, *endpoint)
	if err != nil {
		return err
	}

	http.ListenAndServe(fmt.Sprintf(":%d", *port), mux)
	glog.Info("server running on port:%d", *port)
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
