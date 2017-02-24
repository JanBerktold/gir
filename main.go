package main

import (
	"fmt"
	"os"
	"strings"
)

type Repo struct {
	Owner string
	Name  string
}

type Data struct {
	CurrentRepo  string
	Repositories []Repo
}

func LoadRepo(owner, name string) (Repo, error) {
	return Repo{}, nil
}

func LoadData() (Data, error) {
	return Data{}, nil
}

func SaveData(data Data) error {
	return nil
}

func ParseRepo(repo string) (owner, name string) {
	index := strings.Index(repo, "/")
	if index > -1 {
		return repo[0:index], repo[index+1:]
	}
	return "", ""
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
