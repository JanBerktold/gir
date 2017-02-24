package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

var (
	ErrNoGithubToken = errors.New("No GITHUB_TOKEN specified")
	GraphURL         = "https://api.github.com/graphql"
)

type Client struct {
	client *http.Client
	token  string
}

type Query struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func (c *Client) Execute(str string, args map[string]interface{}, target interface{}) error {
	jsonQuery, err := json.Marshal(Query{
		Query:     str,
		Variables: args,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", GraphURL, bytes.NewBuffer(jsonQuery))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "bearer "+c.token)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func NewClient() (*Client, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if len(token) == 0 {
		return nil, ErrNoGithubToken
	}
	return &Client{
		token:  token,
		client: &http.Client{},
	}, nil
}
