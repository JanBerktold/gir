package main

import (
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add <owner/name>",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	RootCmd.AddCommand(AddCmd)
}
