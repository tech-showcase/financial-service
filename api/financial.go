package api

import (
	"context"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/soheilhy/cmux"
	"github.com/tech-showcase/financial-service/middleware"
	"github.com/tech-showcase/financial-service/presenter"
	digitalCurrencyProto "github.com/tech-showcase/financial-service/proto/digitalcurrency"
	"google.golang.org/grpc"
	"net/http"
)

func RegisterFinancialGRPCAPI(m cmux.CMux) {
	grpcServer := grpc.NewServer(withInterceptor())
	digitalCurrencyProto.RegisterDigitalCurrencyServer(grpcServer, presenter.NewDigitalCurrencyServer())

	grpcListener := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	go grpcServer.Serve(grpcListener)
}

func withInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
		middleware.LoggingInterceptor,
		middleware.JWTAuthenticationInterceptor))
}

func RegisterFinancialHTTPAPI(m cmux.CMux) {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	digitalCurrencyProto.RegisterDigitalCurrencyHandlerServer(ctx, mux, presenter.NewDigitalCurrencyServer())

	httpServer := &http.Server{
		Handler: mux,
	}
	httpListener := m.Match(cmux.HTTP1Fast())
	go httpServer.Serve(httpListener)
}
