package cmd

import (
	"context"

	"github.com/datumforge/datum/pkg/otelx"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/datumforge/datum-cloud/internal/httpserve/config"
	"github.com/datumforge/datum-cloud/internal/httpserve/server"
	"github.com/datumforge/datum-cloud/internal/httpserve/serveropts"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the Datum Cloud API Server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return serve(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.PersistentFlags().String("config", "./config/.config.yaml", "config file location")
}

func serve(ctx context.Context) error {
	serverOpts := []serveropts.ServerOption{}
	serverOpts = append(serverOpts,
		serveropts.WithConfigProvider(&config.ConfigProviderWithRefresh{}),
		serveropts.WithLogger(logger),
		serveropts.WithDatumClient(),
		serveropts.WithHTTPS(),
		serveropts.WithMiddleware(),
		serveropts.WithRateLimiter(),
	)

	so := serveropts.NewServerOptions(serverOpts, k.String("config"))

	err := otelx.NewTracer(so.Config.Settings.Tracer, appName, logger)
	if err != nil {
		logger.Fatalw("failed to initialize tracer", "error", err)
	}

	srv := server.NewServer(so.Config, so.Config.Logger)

	if err := srv.StartEchoServer(ctx); err != nil {
		logger.Error("failed to run server", zap.Error(err))
	}

	return nil
}
