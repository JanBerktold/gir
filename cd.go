package main

import (
	"github.com/spf13/cobra"
)

var CDCmd = &cobra.Command{
	Use:   "cd <owner/name>",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	RootCmd.AddCommand(CDCmd)
}
