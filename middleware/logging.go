package middleware

import (
	"context"
	"github.com/tech-showcase/financial-service/helper"
	"google.golang.org/grpc"
	"strconv"
	"time"
)

func LoggingInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	start := time.Now()

	logger := helper.NewLogger()
	logger.Info("GRPC request processing is started", map[string]string{"method": info.FullMethod})

	defer func() {
		elapsedTime := time.Since(start)
		elapsedTimeStr := strconv.Itoa(int(elapsedTime))

		logger := helper.NewLogger()
		logger.Info("GRPC request processing is finished", map[string]string{"method": info.FullMethod, "elapsed_time": elapsedTimeStr})
	}()

	h, err := handler(ctx, req)

	return h, err
}
