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

type DeleteTarget struct {
	Repositories []string
	Labels       []string
}

func (d *Delete) Parse(bytes []byte) error {
	err := yaml.Unmarshal(bytes, &d.Target)
	if err != nil {
		return err
	}
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
		s := strings.Split(repository, "/")
		if len(s) != 2 {
			fmt.Fprintf(os.Stderr, "repository %s is not properly formatted in setting yaml file\n", repository)
			os.Exit(1)
		}
		owner, repo := s[0], s[1]
		for _, label := range d.Target.Labels {
			fmt.Printf("would delete label \"%s\" for repository \"%s\"\n", label, owner+"/"+repo)
		}
	}
}

func (d *Delete) Execute(ctx context.Context, client *github.Client) {
	for _, repository := range d.Target.Repositories {
		s := strings.Split(repository, "/")
		if len(s) != 2 {
			fmt.Fprintf(os.Stderr, "repository %s is not properly formatted in setting yaml file\n", repository)
			os.Exit(1)
		}
		owner, repo := s[0], s[1]
		for _, label := range d.Target.Labels {
			err := DeleteLabel(ctx, client.Issues, label, owner, repo)
			if err != nil {
				fmt.Printf("error deleting label \"%s\" for repository \"%s\": %s\n", label, owner+"/"+repo, err)
			} else {
				fmt.Printf("deleted label \"%s\" for repository \"%s\"\n", label, owner+"/"+repo)
			}
		}
	}
}
