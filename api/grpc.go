package api

import (
	"fmt"
	"github.com/tech-showcase/financial-service/controller"
	dcProto "github.com/tech-showcase/financial-service/proto/digitalcurrency"
	"google.golang.org/grpc"
	"net"
)

func ActivateGRPC(port int) {
	portStr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", portStr)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	dcProto.RegisterDigitalCurrencyServer(grpcServer, controller.NewDigitalCurrencyServer())
	grpcServer.Serve(lis)
}
