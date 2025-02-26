package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:     "svcmgr",
	Short:   "Stellar Service Manager",
	Long:    `Stellar Service Manager enables management of cosmic services in a clean and unified manner.`,
	Version: "1.0.0-stellar Build: 20240115-lunar | 2025-02-27 00:39:01 SAST",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func AddCommand(cmds ...*cobra.Command) {
	for _, c := range cmds {
		rootCmd.AddCommand(c)
	}
}
