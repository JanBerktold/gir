package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var DropCmd = &cobra.Command{
	Use:   "drop owner/name",
	Short: "Removes a repository's cached data",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			fmt.Println("Invalid number of arguments.")
			os.Exit(-1)
		}

		targetOwner, targetName := ParseRepo(args[0])
		if len(targetOwner) == 0 {
			fmt.Printf("Invalid repo name %q\n", args[0])
			os.Exit(-1)
		}

		data, err := LoadData()
		if err != nil {
			fmt.Printf("Failed to load data with error: %q\n", err.Error())
			os.Exit(-1)
		}

		found := false
		for i, repo := range data.Repositories {
			if repo.Owner.Name == targetOwner && repo.Name == targetName {
				data.Repositories = append(data.Repositories[:i], data.Repositories[i+1:]...)
				found = true
			}
		}

		if found {
			if err := SaveData(data); err != nil {
				fmt.Printf("Failed to write to disk: %q\n", err.Error())
				os.Exit(-1)
			}
		} else {
			fmt.Printf("Repository %s/%s is not cached.\n", targetOwner, targetName)
			os.Exit(-1)
		}
	},
}

func init() {
	RootCmd.AddCommand(DropCmd)
}
