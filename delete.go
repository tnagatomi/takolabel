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

func DryRunDelete(target DeleteTarget) {
	for _, repository := range target.Repositories {
		s := strings.Split(repository, "/")
		if len(s) != 2 {
			fmt.Fprintf(os.Stderr, "repository %s is not properly formatted in setting yaml file\n", repository)
			os.Exit(1)
		}
		owner, repo := s[0], s[1]
		for _, label := range target.Labels {
			fmt.Printf("would delete label \"%s\" for repository \"%s\"\n", label, owner+"/"+repo)
		}
	}
}

func ExecuteDelete(ctx context.Context, client *github.Client, target DeleteTarget) {
	for _, repository := range target.Repositories {
		s := strings.Split(repository, "/")
		if len(s) != 2 {
			fmt.Fprintf(os.Stderr, "repository %s is not properly formatted in setting yaml file\n", repository)
			os.Exit(1)
		}
		owner, repo := s[0], s[1]
		for _, label := range target.Labels {
			err := DeleteLabel(ctx, client.Issues, label, owner, repo)
			if err != nil {
				fmt.Printf("error deleting label \"%s\" for repository \"%s\": %s\n", label, owner+"/"+repo, err)
			} else {
				fmt.Printf("deleted label \"%s\" for repository \"%s\"\n", label, owner+"/"+repo)
			}
		}
	}
}
