package api

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func Activate(port int) {
	lisAddress := fmt.Sprintf(":%d", port)
	lis, err := net.Listen("tcp", lisAddress)
	if err != nil {
		panic(err)
	}

	cMux := cmux.New(lis)
	ActivateGRPC(cMux)
	ActivateHTTP(cMux)

	cMux.Serve()
}

func ActivateGRPC(cMux cmux.CMux) {
	grpcServer := grpc.NewServer(withInterceptor())
	RegisterFinancialGRPCAPI(grpcServer)

	grpcListener := cMux.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	go grpcServer.Serve(grpcListener)
}

func ActivateHTTP(cMux cmux.CMux) {
	runtimeMux := runtime.NewServeMux()
	RegisterFinancialHTTPAPI(runtimeMux)

	httpMux := http.NewServeMux()
	httpMux.Handle("/", runtimeMux)
	RegisterOAuth2HTTPAPI(httpMux)

	httpListener := cMux.Match(cmux.HTTP1())
	httpServer := &http.Server{
		Handler: httpMux,
	}
	go httpServer.Serve(httpListener)
}
