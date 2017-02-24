package main

import (
	"errors"
	"net/http"
	"os"
)

var (
	ErrNoGithubToken = errors.New("No GITHUB_TOKEN specified")
)

type Client struct {
	client *http.Client
	token  string
}

func (c *Client) Execute(str string, target interface{}) {
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
