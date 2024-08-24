package main

import (
	"github.com/spf13/cobra"

	"github.com/nint8835/monticola/cmd/common"
	"github.com/nint8835/monticola/pkg/agent/api/server"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the agent",
	Run: func(cmd *cobra.Command, args []string) {
		serverInst, err := server.New(agentConfig)
		common.CheckError(err, "Error creating agent server")

		err = serverInst.Start()
		common.CheckError(err, "Error starting agent server")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
