package main

import (
	"github.com/spf13/cobra"

	"github.com/nint8835/monticola/cmd/common"
	"github.com/nint8835/monticola/pkg/server"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the server",
	Run: func(cmd *cobra.Command, args []string) {
		serverInst, err := server.New(serverConfig)
		common.CheckError(err, "Error creating server")

		err = serverInst.Start()
		common.CheckError(err, "Error starting server")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
