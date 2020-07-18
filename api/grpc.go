package api

import (
	"fmt"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/tech-showcase/financial-service/middleware"
	"github.com/tech-showcase/financial-service/presenter"
	dcProto "github.com/tech-showcase/financial-service/proto/digitalcurrency"
	"google.golang.org/grpc"
	"net"
)

func ActivateGRPC(port int) {
	lisAddress := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", lisAddress)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(withInterceptor())
	dcProto.RegisterDigitalCurrencyServer(grpcServer, presenter.NewDigitalCurrencyServer())
	grpcServer.Serve(lis)
}

func withInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
		middleware.LoggingInterceptor,
		middleware.AuthorizationInterceptor))
}
