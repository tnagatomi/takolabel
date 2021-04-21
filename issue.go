package takolabel

import (
	"context"
	"github.com/google/go-github/v33/github"
	"github.com/tommy6073/takolabel/config"
)

func CreateLabel(ctx context.Context, issuesService IssuesService, label config.Label, repository config.Repository) (*github.Label, error) {
	githubLabel := &github.Label{
		Name:        github.String(label.Name),
		Description: github.String(label.Description),
		Color:       github.String(label.Color),
	}
	createdLabel, _, err := issuesService.CreateLabel(ctx, repository.Org, repository.Repo, githubLabel)
	return createdLabel, err
}

type IssuesService interface {
	CreateLabel(ctx context.Context, owner string, repository string, label *github.Label) (*github.Label, *github.Response, error)
}
