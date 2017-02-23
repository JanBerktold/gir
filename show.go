package main

import (
	"github.com/spf13/cobra"
)

var ShowCmd = &cobra.Command{
	Use:   "show [<owner/name>] <id>",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	RootCmd.AddCommand(ShowCmd)
}
