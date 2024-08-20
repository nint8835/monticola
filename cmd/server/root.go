package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"github.com/nint8835/monticola/pkg/config"
)

var rootCmd = &cobra.Command{
	Use: "monticola-server",
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("Hello world!")
	},
}

var serverConfig config.ServerConfig

func init() {
	cobra.OnInitialize(func() {
		configInst, err := config.LoadServerConfig()
		if err != nil {
			slog.Error("Error loading server config", "err", err)
			os.Exit(1)
		}

		serverConfig = *configInst
	})
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
