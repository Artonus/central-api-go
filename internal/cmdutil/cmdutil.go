package cmdutil

import (
	"context"
	_ "database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.uber.org/zap"
	"log"
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
func CreateDbConnection() *sqlx.DB {

	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, 5432, username, pass, "central-api")

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}
	return db
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
	session := driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: ""})

	defer session.Close(ctx)
	// Run index queries via implicit auto-commit transaction
	indexes := []string{
		"create index on :Location(id)",
		"create index on :Location(name)",
	}
	for _, index := range indexes {
		_, err = session.Run(ctx, index, nil)
		if err != nil {
			panic(err)
		}
	}

	return driver
}
