package cmd

import (
	"context"

	"github.com/cloudflare/tableflip"
	"github.com/oklog/run"
	"github.com/spf13/cobra"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/pkg/platform"

	"github.com/scraly/hello-world/internal/version"
)

// -----------------------------------------------------------------------------

var serverCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "Starts the Schema gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Initialize config
		initConfig()

		// Start goroutine group
		err := platform.Run(ctx, &platform.Application{
			Debug:           conf.Debug.Enable,
			Name:            "schema-grpc",
			Version:         version.Version,
			Revision:        version.Revision,
			Instrumentation: conf.Instrumentation,
			Builder: func(upg *tableflip.Upgrader, group *run.Group) {
				// Starting banner
				log.For(ctx).Info("Starting schema gRPC server ...")

				// Allocate listener
				// ln, err := upg.Fds.Listen(conf.Server.Network, conf.Server.Listen)
				// if err != nil {
				// 	log.For(ctx).Fatal("Unable to start GRPC server", zap.Error(err))
				// }

				// TODO: Attach the dispatcher
				// server, err := grpc.New(ctx, conf)
				// if err != nil {
				// 	log.For(ctx).Fatal("Unable to start GRPC server", zap.Error(err))
				// }

				// // Add to goroutine group
				// group.Add(
				// 	func() error {
				// 		log.For(ctx).Info("GRPC server listening ...", zap.Stringer("address", ln.Addr()))
				// 		return server.Serve(ln)
				// 	},
				// 	func(e error) {
				// 		log.For(ctx).Info("Shutting GRPC server down")
				// 		server.GracefulStop()
				// 	},
				// )
			},
		})
		log.CheckErrCtx(ctx, "Unable to run application", err)
	},
}

// -----------------------------------------------------------------------------

func init() {
	serverCmd.Flags().StringVar(&cfgFile, "config", "", "config file")
}
