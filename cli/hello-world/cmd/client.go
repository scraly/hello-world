package cmd

import (
	helloworld "github.com/scraly/hello-world/pkg/protocol/helloworld/v1"

	"github.com/spf13/cobra"
)

// -----------------------------------------------------------------------------

var clientCmd = &cobra.Command{
	Use:     "client",
	Aliases: []string{"c", "cli"},
	Short:   "Query the gRPC server",
}

func init() {
	clientCmd.AddCommand(
		helloworld.GreeterClientCommand,
	)
}
