package helper

import (
	"encoding/json"
	"go.uber.org/zap"
)

type (
	logger struct {
		logger *zap.Logger
	}
	Logger interface {
		Info(message string, additionalData map[string]string)
		Debug(message string, additionalData map[string]string)
		Warn(message string, additionalData map[string]string)
		Error(message string, additionalData map[string]string)
	}
)

func NewLogger() Logger {
	instance := logger{}

	logger, err := zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		return nil
	}
	instance.logger = logger

	return &instance
}

func (instance *logger) Info(message string, additionalData map[string]string) {
	fields := instance.Construct(additionalData)

	instance.logger.Info(
		message,
		fields...,
	)
}

func (instance *logger) Debug(message string, additionalData map[string]string) {
	fields := instance.Construct(additionalData)

	instance.logger.Debug(
		message,
		fields...,
	)
}

func (instance *logger) Warn(message string, additionalData map[string]string) {
	fields := instance.Construct(additionalData)

	instance.logger.Warn(
		message,
		fields...,
	)
}

func (instance *logger) Error(message string, additionalData map[string]string) {
	fields := instance.Construct(additionalData)

	instance.logger.Error(
		message,
		fields...,
	)
}

func (instance *logger) Construct(additionalData map[string]string) []zap.Field {
	fields := make([]zap.Field, 0)

	additionalDataJson, _ := json.Marshal(additionalData)
	fields = append(fields, zap.String("data", string(additionalDataJson)))

	return fields
}
