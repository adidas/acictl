package github

import (
	"context"
	"fmt"

	"github.com/adidas/acictl/utils"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func New(token string) Git {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: token,
		},
	)
	tokenClient := oauth2.NewClient(ctx, ts)

	return Git{
		Client: github.NewClient(tokenClient),
		Ctx:    ctx,
	}
}

func (g Git) CreateRepo(org, instance string) {
	repo := &github.Repository{
		Name: github.String(
			fmt.Sprintf(
				utils.Repo().Instance,
				instance,
			),
		),
		Private: github.Bool(false),
	}

	g.Client.Repositories.Create(g.Ctx, "acictl", repo)
}
