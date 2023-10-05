package cmdutil

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
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

func CreateGraphConnection() neo4j.DriverWithContext {

	address := os.Getenv("GRAPH_ADDRESS")
	username := os.Getenv("GRAPH_USERNAME")
	pass := os.Getenv("GRAPH_PASSWORD")
	driver, err := neo4j.NewDriverWithContext(address, neo4j.BasicAuth(username, pass, ""))

	ctx := context.Background()
	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Viola! Connected to Memgraph!")
	}

	return driver
}
