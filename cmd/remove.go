package cmd

import (
    "github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
    Use:   "remove",
    Short: "remove command",
    Run: func(cmd *cobra.Command, args []string) {
        // Implementation here
    },
}

func init() {
    rootCmd.AddCommand(removeCmd)
}
