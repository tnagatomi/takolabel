package takolabel

import (
	"context"
	"fmt"
	"github.com/google/go-github/v33/github"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type Delete struct {
	Target DeleteTarget
}

type DeleteTargetConfig struct {
	Repositories []string
	Labels       []string
}

type DeleteTarget struct {
	Repositories []Repository
	Labels       []string
}

func (d *Delete) Parse(bytes []byte) error {
	targetConfig := DeleteTargetConfig{}
	err := yaml.Unmarshal(bytes, &targetConfig)
	if err != nil {
		return err
	}

	target := DeleteTarget{Labels: targetConfig.Labels}
	for _, repository := range targetConfig.Repositories {
		s := strings.Split(repository, "/")
		if len(s) != 2 {
			return fmt.Errorf("repository %s is not properly formatted in setting yaml file", repository)
		}
		target.Repositories = append(target.Repositories, Repository{s[0], s[1]})
	}

	d.Target = target

	return nil
}

func (d *Delete) Gather() error {
	content, err := os.ReadFile("takolabel_delete.yaml")
	if err != nil {
		return err
	}

	err = d.Parse(content)
	if err != nil {
		return err
	}

	return nil
}

func (d *Delete) DryRun() {
	for _, repository := range d.Target.Repositories {
		for _, label := range d.Target.Labels {
			fmt.Printf("would delete label \"%s\" for repository \"%s\"\n", label, repository.Owner+"/"+repository.Repo)
		}
	}
}

func (d *Delete) Execute(ctx context.Context, client *github.Client) {
	for _, repository := range d.Target.Repositories {
		for _, label := range d.Target.Labels {
			err := DeleteLabel(ctx, client.Issues, label, repository.Owner, repository.Repo)
			if err != nil {
				fmt.Printf("error deleting label \"%s\" for repository \"%s\": %s\n", label, repository.Owner+"/"+repository.Repo, err)
			} else {
				fmt.Printf("deleted label \"%s\" for repository \"%s\"\n", label, repository.Owner+"/"+repository.Repo)
			}
		}
	}
}
