package cmd

// TODO: fix
// helloworld "github.com/scraly/hello-world/pkg/protocol/helloworld"

import "github.com/spf13/cobra"

// -----------------------------------------------------------------------------

var clientCmd = &cobra.Command{
	Use:     "client",
	Aliases: []string{"c", "cli"},
	Short:   "Query the gRPC server",
}

func init() {
	clientCmd.AddCommand(
	// helloworld.DbAPIClientCommand,
	)
}
