package main

import (
	"errors"
	"time"
)

const ISSUE_LIMIT = 100

type Entity struct {
	Name string `json:"login"`
}

type Pagination struct {
	Cursor string `json:"endCursor"`
}

type Label struct {
	Name string `json:"name"`
}

type TimelineItem struct {
	Type string `json:"__typename"`

	// IssueComment <https://developer.github.com/early-access/graphql/object/issuecomment/>
	Author Entity `json:"author"`
	Editor Entity `json:"editor"`
	Body string `json:"body"`
	CreatedAt time.Time `json:"createdAt"`

	// ClosedEvent <https://developer.github.com/early-access/graphql/object/closedevent/>
	Actor Entity `json:"actor"`
	// +CreatedAt

	// ReopenedEvent <https://developer.github.com/early-access/graphql/object/reopenedevent/>
	// +Actor
	// +CreatedAt

	// LabeledEvent <https://developer.github.com/early-access/graphql/object/labeledevent/>
	// +Actor
	// +CreatedAt
	Labels []Label `json:"label"`
}

type Timeline struct {
	Timeline []TimelineItem `json:"nodes"`
	NextPage Pagination `json:"pageInfo"`
}

type Issue struct {
	Author Entity `json:"author"`
	Editor Entity `json:"editor"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	State  string `json:"state"`
	Id     string `json:"id"`
	Number int    `json:"number"`
	Timeline Timeline  `json:"timeline"`
}


type Issues struct {
	Issues   []Issue         `json:"nodes"`
	NextPage Pagination `json:"pageInfo"`
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
    	pageInfo {
        endCursor
      }
      nodes {
        id
        body
        createdAt
        number
        title
        updatedAt
        state
        timeline(first:100) {
          pageInfo {
            endCursor
          }
          nodes {
            __typename
            ... on IssueComment {
              author {
                login
              }
              body
              createdAt
            }
	    ... on ClosedEvent {
              actor {
                login
              }
              createdAt
	    }
	    ... on ReopenedEvent { 
              actor {
                login
              }
              createdAt
	    }
	    ... on LabeledEvent {
              actor {
                login
              }
              createdAt
	    }
         }
        }
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

var followUpQuery = `
query ($owner: String!, $name: String!, $after: String!) {
  repository(owner: $owner, name: $name) {
    owner {
      login
    }
    name
    issues(first: 100, after: $after) {
    	pageInfo {
        endCursor
      }
      nodes {
        id
        body
        createdAt
        number
        title
        updatedAt
        state
        timeline(first:100) {
          pageInfo {
            endCursor
          }
          nodes {
            __typename
            ... on IssueComment {
              author {
                login
              }
              body
              createdAt
            }
	    ... on ClosedEvent {
              actor {
                login
              }
              createdAt
	    }
	    ... on ReopenedEvent { 
              actor {
                login
              }
              createdAt
	    }
	    ... on LabeledEvent {
              actor {
                login
              }
              createdAt
	    }
         }
        }
        editor {
          login
        }
        author {
          login
        }
      }
    }
  }
}`

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

	if err != nil {
		return Repository{}, err
	}

	if len(repo.Message) > 0 {
		return repo.Data.Repository, errors.New(repo.Message)
	}

	//TODO: Handle hitting the ratelimit gracefully
	if len(repo.Data.Repository.Issues.NextPage.Cursor) > 0 {
		after := repo.Data.Repository.Issues.NextPage.Cursor
		for {
			var repo2 Wrapper
			err := client.Execute(followUpQuery, map[string]interface{}{
				"owner": owner,
				"name":  name,
				"after": after,
			}, &repo2)
			if err != nil {
				return Repository{}, err
			}

			if len(repo2.Message) > 0 {
				return repo2.Data.Repository, errors.New(repo.Message)
			}

			repo.Data.Repository.Issues.Issues = append(repo.Data.Repository.Issues.Issues, repo2.Data.Repository.Issues.Issues...)

			if len(repo2.Data.Repository.Issues.NextPage.Cursor) > 0 {
				after = repo2.Data.Repository.Issues.NextPage.Cursor
			} else {
				break
			}

		}
	}

	return repo.Data.Repository, err
}
