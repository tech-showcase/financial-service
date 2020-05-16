package middleware

import (
	"context"
	"errors"
	"github.com/tech-showcase/financial-service/global"
	"github.com/tech-showcase/financial-service/helper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"time"
)

func AuthorizationInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	}

	authHeader, ok := md["authorization"]
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	}

	err := authorizeJWT(authHeader[0])
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	h, err := handler(ctx, req)

	return h, err
}

func authorizeJWT(token string) error {
	authEndpoint, _ := helper.JoinURL(global.Configuration.Auth.ServerAddress, "/api/user")

	req, err := http.NewRequest("GET", authEndpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("authorization is failed")
	}

	return nil
}

func LoggingInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	start := time.Now()

	logger := helper.NewLogBlueprint()
	logger.Info("GRPC log interceptor is started", map[string]string{"method": info.FullMethod})

	defer func() {
		elapsedTime := time.Since(start)
		elapsedTimeStr := strconv.Itoa(int(elapsedTime))

		logger := helper.NewLogBlueprint()
		logger.Info("GRPC log interceptor is finished", map[string]string{"method": info.FullMethod, "elapsed_time": elapsedTimeStr})
	}()

	h, err := handler(ctx, req)

	return h, err
}
