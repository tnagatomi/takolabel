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

type CreateTarget struct {
	Repositories []string
	Labels       []Label
}

func (c *Create) Parse(bytes []byte) error {
	err := yaml.Unmarshal(bytes, &c.Target)
	if err != nil {
		return err
	}
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
		s := strings.Split(repository, "/")
		if len(s) != 2 {
			fmt.Fprintf(os.Stderr, "repository %s is not properly formatted in setting yaml file\n", repository)
			os.Exit(1)
		}
		owner, repo := s[0], s[1]
		for _, label := range c.Target.Labels {
			fmt.Printf("Would create label \"%s\" for repository \"%s\"\n", label.Name, owner+"/"+repo)
		}
	}
}

func (c *Create) Execute(ctx context.Context, client *github.Client) {
	for _, repository := range c.Target.Repositories {
		s := strings.Split(repository, "/")
		if len(s) != 2 {
			fmt.Fprintf(os.Stderr, "repository %s is not properly formatted in setting yaml file\n", repository)
			os.Exit(1)
		}
		owner, repo := s[0], s[1]
		for _, label := range c.Target.Labels {
			_, err := CreateLabel(ctx, client.Issues, label, owner, repo)
			if err != nil {
				fmt.Printf("error creating label \"%s\" for repository \"%s\": %s\n", label.Name, owner+"/"+repo, err)
			} else {
				fmt.Printf("created label \"%s\" for repository \"%s\"\n", label.Name, owner+"/"+repo)
			}
		}
	}
}
