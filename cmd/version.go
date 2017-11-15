package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the markdown-toc version",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("markdown-toc %s\n", version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
