package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	groupBy    string
	showHidden bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List cosmic services",
	Run:   runList,
}

func init() {
	listCmd.Flags().StringVarP(&groupBy, "group-by", "g", "", "Group by (type|category)")
	listCmd.Flags().BoolVar(&showHidden, "show-hidden", false, "Reveal celestial secrets")
	// The command registration happens in root.go for clarity.
	AddCommand(listCmd)
}

func runList(cmd *cobra.Command, args []string) {
	// Disable extra logging during command display.
	fmt.Println("NAME  TYPE  CATEGORY  HOST  PORT")
	// ... implementation details
}
