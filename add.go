package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var AddCmd = &cobra.Command{
	Use:   "add <owner/name>",
	Short: "Cache an additional repository's issues",
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

		for _, repo := range data.Repositories {
			if repo.Owner.Name == targetOwner && repo.Name == targetName {
				fmt.Printf("Repository %q is already cached.\nTo update the repository use:\n\tgir update %s/%s\n", args[0], targetOwner, targetName)
				os.Exit(-1)
			}
		}

		repo, err := LoadRepo(targetOwner, targetName)
		if err != nil {
			fmt.Printf("Failed to cache repository %s/%s: %q\n", targetOwner, targetName, err.Error())
		}
		data.Repositories = append(data.Repositories, repo)

		if err := SaveData(data); err != nil {
			fmt.Printf("Failed to write to disk: %q\n", err.Error())
			os.Exit(-1)
		}
	},
}

func init() {
	RootCmd.AddCommand(AddCmd)
}
