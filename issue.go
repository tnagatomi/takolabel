package takolabel

import (
	"context"
	"fmt"
	"github.com/google/go-github/v33/github"
	"takolabel/config"
)

func CreateLabels(issuesClient *IssuesClient, repositories []config.Repository, labels []config.Label) int {
	createdLabelsCount := 0
	for _, repository := range repositories {
		for _, label := range labels {
			_, err := issuesClient.CreateLabel(label, repository)
			if err != nil {
				fmt.Printf("error creating label \"%s\" for repository \"%s\": %s\n", label.Name, repository.Org+"/"+repository.Repo, err)
			} else {
				fmt.Printf("created label \"%s\" for repository \"%s\"\n", label.Name, repository.Org+"/"+repository.Repo)
				createdLabelsCount++
			}
		}
	}
	return createdLabelsCount
}

type IssuesClient struct {
	Ctx           context.Context
	IssuesService IssuesService
}

func (ic *IssuesClient) CreateLabel(label config.Label, repository config.Repository) (*github.Label, error) {
	githubLabel := &github.Label{
		Name:        github.String(label.Name),
		Description: github.String(label.Description),
		Color:       github.String(label.Color),
	}
	createdLabel, _, err := ic.IssuesService.CreateLabel(ic.Ctx, repository.Org, repository.Repo, githubLabel)
	return createdLabel, err
}

type IssuesService interface {
	CreateLabel(ctx context.Context, owner string, repository string, label *github.Label) (*github.Label, *github.Response, error)
}
