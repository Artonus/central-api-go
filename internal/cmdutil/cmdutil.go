package cmdutil

import (
	"github.com/redis/rueidis"
	"go.uber.org/zap"
	"os"
)

func CreateLogger(serviceName string) *zap.Logger {
	env := os.Getenv("ENV")

	logger, _ := zap.NewProduction(zap.Fields(
		zap.String("env", env),
		zap.String("service", serviceName)))

	if env == "" || env == "DEBUG" {
		logger, _ = zap.NewDevelopment()
	}

	return logger
}

func CreateRedisConnection() rueidis.Client {
	//TODO: add env variables or secret manager to create connection string
	pass := os.Getenv("REDIS_PASSWORD")
	hostname := os.Getenv("REDIS_HOSTNAME")
	c, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{hostname},
		Username:    "",
		Password:    pass,
	})
	if err != nil {
		panic(err)
	}
	return c
}
