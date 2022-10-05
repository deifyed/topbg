package cmd

import (
	"github.com/deifyed/topbg/cmd/list"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	RunE:  list.RunE(fs),
}

func init() {
	rootCmd.AddCommand(listCmd)
}
