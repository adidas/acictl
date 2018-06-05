package github

import (
	"context"

	"github.com/google/go-github/github"
)

type Client interface {
	CreateRepo()
}

type Git struct {
	Client *github.Client
	Ctx    context.Context
}
