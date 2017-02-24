package main

import (
	"github.com/spf13/cobra"
	"strings"
	"strconv"
	"fmt"
	"os"
)

const MAX_LENGTH = 80

func PrepareBody(str string) string {
	resp := ""
	for _, line := range strings.Split(str, "\n") {
		for len(line) > MAX_LENGTH {
			resp += fmt.Sprintf("\t%s\n", line[:MAX_LENGTH-1])
			line = line[MAX_LENGTH-1:]
		}
		resp += fmt.Sprintf("\t%s\n", line)
	}
	return resp[0:len(resp)-1]
}

var ShowCmd = &cobra.Command{
	Use:   "show [<owner/name>] <id>",
	Short: "Gives detailed information about an issue including comments",
	Run: func(cmd *cobra.Command, args []string) {

		data, err := LoadData()
		if err != nil {
			fmt.Printf("Failed to load data with error: %q\n", err.Error())
			os.Exit(-1)
		}

		var targetOwner, targetName string
		var issueIdentifier string
		if len(args) == 1 {
			if data.CurrentRepo == nil {
				fmt.Println("Neither arg nor current repo specified\n")
				os.Exit(-1)
			} else {
				targetOwner = data.CurrentRepo.Owner
				targetName = data.CurrentRepo.Name
				issueIdentifier = args[0]
			}
		} else if len(args) == 2 {
			targetOwner, targetName = ParseRepo(args[0])
			issueIdentifier = args[1]
			if len(targetOwner) == 0 {
				fmt.Printf("Invalid repo name %q given\n", args[0])
				os.Exit(-1)
			}
		}

		number, err := strconv.Atoi(issueIdentifier)
		if err != nil {
			fmt.Printf("%q is not a valid issue number.\n", issueIdentifier)
			os.Exit(-1)
		}

		for _, repo := range data.Repositories {
			if repo.Owner.Name == targetOwner && repo.Name == targetName {
				for _, issue := range repo.Issues.Issues {
					if issue.Number == number {
						fmt.Printf("%s: %s #%d (%s) \n\n", issue.Author.Name, issue.Title, issue.Number, issue.State)
						fmt.Println(PrepareBody(issue.Body))
						for _, item := range issue.Timeline.Timeline {
							switch item.Type {
							case "IssueComment":
								fmt.Printf("%s commented on %s:\n", item.Author.Name, item.CreatedAt.Format("Jan 2 3:04 PM"))
								fmt.Println(PrepareBody(item.Body))
							case "ClosedEvent":
								fmt.Printf("%s closed this issue on %s\n", item.Actor.Name, item.CreatedAt.Format("Jan 2 3:04 PM"))
							case "ReopenedEvent":
								fmt.Printf("%s reopened this issue on %s\n", item.Actor.Name, item.CreatedAt.Format("Jan 2 3:04 PM"))
							default:
								fmt.Println(item.Type)
							}
						}
						return
					}
				}
				fmt.Printf("Issue %d does not exist in repo %s/%s (perhaps run gir update?)\n",
						targetOwner, targetName, number)
				return
			}
		}

		fmt.Printf("Repository %s/%s is not cached.\n", targetOwner, targetName)
	},
}

func init() {
	RootCmd.AddCommand(ShowCmd)
}
