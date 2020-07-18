package middleware

import (
	"errors"
	"github.com/tech-showcase/financial-service/config"
	"github.com/tech-showcase/financial-service/helper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
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
	authEndpoint, _ := helper.JoinURL(config.Instance.Auth.ServerAddress, "/api/user")

	req, err := http.NewRequest("GET", authEndpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+token)

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
