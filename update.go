package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var UpdateCmd = &cobra.Command{
	Use:   "update [<owner/name>]",
	Short: "Updates cached issues",
	Run: func(cmd *cobra.Command, args []string) {

		// parse whether the user specified a target repo
		var targetOwner, targetName string
		if len(args) > 0 {
			targetOwner, targetName = ParseRepo(args[0])
			if len(targetOwner) == 0 {
				fmt.Printf("Invalid repo name %q\n", args[0])
				os.Exit(-1)
			}
		}

		data, err := LoadData()
		if err != nil {
			fmt.Printf("Failed to load data with error: %q\n", err.Error())
			os.Exit(-1)
		}

		if len(targetOwner) == 0 {
			for i, repo := range data.Repositories {
				newRepo, err := LoadRepo(repo.Owner.Name, repo.Name)
				if err != nil {
					fmt.Printf("Failed to load repo %s/%s: %q\n", repo.Owner, repo.Name, err.Error())
				} else {
					data.Repositories[i] = newRepo
				}
			}
		} else {
			for i, repo := range data.Repositories {
				if repo.Owner.Name == targetOwner && repo.Name == targetName {
					newRepo, err := LoadRepo(repo.Owner.Name, repo.Name)
					if err != nil {
						fmt.Printf("Failed to load repo %s/%s: %q\n", repo.Owner, repo.Name, err.Error())
						os.Exit(-1)
					}
					data.Repositories[i] = newRepo
					goto finished
				}
			}
		}

	finished:
		if err := SaveData(data); err != nil {
			fmt.Printf("Failed to write to disk: %q\n", err.Error())
			os.Exit(-1)
		}
	},
}

func init() {
	RootCmd.AddCommand(UpdateCmd)
}
