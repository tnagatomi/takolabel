package takolabel

import (
	"context"
	"github.com/google/go-github/v33/github"
	"github.com/tommy6073/takolabel/config"
)

func CreateLabel(ctx context.Context, issuesService *github.IssuesService, label config.Label, owner string, repo string) (*github.Label, error) {
	githubLabel := &github.Label{
		Name:        github.String(label.Name),
		Description: github.String(label.Description),
		Color:       github.String(label.Color),
	}
	createdLabel, _, err := issuesService.CreateLabel(ctx, owner, repo, githubLabel)
	return createdLabel, err
}

func DeleteLabel(ctx context.Context, issuesService *github.IssuesService, label string, owner string, repo string) error {
	_, err := issuesService.DeleteLabel(ctx, owner, repo, label)
	return err
}
