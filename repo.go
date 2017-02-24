package main

import "errors"

type Entity struct {
	Name string `json:"login"`
}

type Issue struct {
	Author Entity `json:"author"`
	Editor Entity `json:"editor"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	State  string `json:"state"`
	id     string `json:"id"`
	Number int    `json:"number"`
}

type Issues struct {
	Issues []Issue `json:"nodes"`
}

type Repository struct {
	Owner  Entity `json:"owner"`
	Name   string `json:"name"`
	Issues Issues `json:"issues"`
}

var query = `
query ($owner: String!, $name: String!) {
  repository(owner: $owner, name: $name) {
    owner {
      login
    }
    name
    issues(first: 100) {
      nodes {
        id
        body
        createdAt
        number
        title
        updatedAt
        state
        editor {
          login
        }
        author {
          login
        }
      }
    }
  }
}
`

type Wrapper struct {
	Message string `json:"message"`
	Data    struct {
		Repository Repository `json:"repository"`
	} `json:"data"`
}

func LoadRepo(owner, name string) (Repository, error) {
	client, err := NewClient()
	if err != nil {
		return Repository{}, err
	}

	var repo Wrapper
	err = client.Execute(query, map[string]interface{}{
		"owner": owner,
		"name":  name,
	}, &repo)

	if len(repo.Message) > 0 {
		return repo.Data.Repository, errors.New(repo.Message)
	}

	return repo.Data.Repository, err
}
