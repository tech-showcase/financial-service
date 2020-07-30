package api

import (
	"context"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/tech-showcase/financial-service/middleware"
	"github.com/tech-showcase/financial-service/presenter"
	digitalCurrencyProto "github.com/tech-showcase/financial-service/proto/digitalcurrency"
	"google.golang.org/grpc"
)

func RegisterFinancialGRPCAPI(grpcServer *grpc.Server) {
	digitalCurrencyProto.RegisterDigitalCurrencyServer(grpcServer, presenter.NewDigitalCurrencyServer())
}

func withInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
		middleware.LoggingInterceptor,
		middleware.JWTAuthenticationInterceptor))
}

func RegisterFinancialHTTPAPI(mux *runtime.ServeMux) {
	ctx := context.Background()
	digitalCurrencyProto.RegisterDigitalCurrencyHandlerServer(ctx, mux, presenter.NewDigitalCurrencyServer())
}
