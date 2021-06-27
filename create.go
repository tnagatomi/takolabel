package takolabel

import (
	"context"
	"fmt"
	"github.com/google/go-github/v33/github"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type Create struct {
	Target CreateTarget
}

type CreateTargetConfig struct {
	Repositories []string
	Labels       []Label
}

type CreateTarget struct {
	Repositories []Repository
	Labels       []Label
}

func (c *Create) Parse(bytes []byte) error {
	targetConfig := CreateTargetConfig{}
	err := yaml.Unmarshal(bytes, &targetConfig)
	if err != nil {
		return err
	}

	target := CreateTarget{Labels: targetConfig.Labels}
	for _, repository := range targetConfig.Repositories {
		s := strings.Split(repository, "/")
		if len(s) != 2 {
			return fmt.Errorf("repository %s is not properly formatted in setting yaml file", repository)
		}
		target.Repositories = append(target.Repositories, Repository{s[0], s[1]})
	}

	c.Target = target

	return nil
}

func (c *Create) Gather() error {
	content, err := os.ReadFile("takolabel_create.yaml")
	if err != nil {
		return err
	}

	err = c.Parse(content)
	if err != nil {
		return err
	}

	return nil
}

func (c *Create) DryRun() {
	for _, repository := range c.Target.Repositories {
		for _, label := range c.Target.Labels {
			fmt.Printf("Would create label \"%s\" for repository \"%s\"\n", label.Name, repository.Owner+"/"+repository.Repo)
		}
	}
}

func (c *Create) Execute(ctx context.Context, client *github.Client) {
	for _, repository := range c.Target.Repositories {
		for _, label := range c.Target.Labels {
			_, err := CreateLabel(ctx, client.Issues, label, repository.Owner, repository.Repo)
			if err != nil {
				fmt.Printf("error creating label \"%s\" for repository \"%s\": %s\n", label.Name, repository.Owner+"/"+repository.Repo, err)
			} else {
				fmt.Printf("created label \"%s\" for repository \"%s\"\n", label.Name, repository.Owner+"/"+repository.Repo)
			}
		}
	}
}
