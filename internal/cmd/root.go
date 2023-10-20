package cmd

import (
	"context"
	"github.com/Artonus/central-api-go/internal/api"
	"github.com/Artonus/central-api-go/internal/cmdutil"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"net/http"
)

func Execute(ctx context.Context) int {
	_ = godotenv.Load()

	apiCmd := apiCmd(ctx)

	go func() {
		_ = http.ListenAndServe("localhost:6060", nil)
	}()

	if err := apiCmd.Execute(); err != nil {
		return 1
	}

	return 0
}

func apiCmd(ctx context.Context) *cobra.Command {
	var port int
	cmd := &cobra.Command{
		Use:   "api",
		Args:  cobra.ExactArgs(0),
		Short: "Runs API",
		RunE: func(cmd *cobra.Command, args []string) error {
			port = 8080
			logger := cmdutil.CreateLogger("Api")
			defer func() { _ = logger.Sync() }()

			graph := cmdutil.CreateGraphConnection()
			ctx := context.Background()
			defer graph.Close(ctx)

			db := cmdutil.CreateDbConnection()

			api := api.CreateApi(ctx, logger, graph, db)
			srv := api.Server(port)

			go func() { _ = srv.ListenAndServe() }()

			logger.Info("started api", zap.Int("port", port))

			<-ctx.Done()

			_ = srv.Shutdown(ctx)

			return nil
		},
	}
	return cmd
}
