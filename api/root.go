package api

import (
	"fmt"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/tech-showcase/financial-service/middleware"
	"github.com/tech-showcase/financial-service/presenter"
	digitalCurrencyProto "github.com/tech-showcase/financial-service/proto/digitalcurrency"
	"google.golang.org/grpc"
	"net"
)

func Activate(port int) {
	lisAddress := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", lisAddress)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(withInterceptor())
	digitalCurrencyProto.RegisterDigitalCurrencyServer(grpcServer, presenter.NewDigitalCurrencyServer())
	grpcServer.Serve(lis)
}

func withInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
		middleware.LoggingInterceptor,
		middleware.JWTAuthenticationInterceptor))
}
