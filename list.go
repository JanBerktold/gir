package main

import (
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list [<owner/name>]",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	RootCmd.AddCommand(ListCmd)
}
