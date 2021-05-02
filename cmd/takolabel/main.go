package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v33/github"
	"github.com/spf13/viper"
	"github.com/tommy6073/takolabel"
	"github.com/tommy6073/takolabel/config"
	"golang.org/x/oauth2"
	"strings"
)

func main() {
	ctx := context.Background()

	viper.SetConfigName("takolabel")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config: %s", err))
	}

	githubToken := viper.GetString("GITHUB_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	baseUrl := viper.GetString("GITHUB_SERVER_URL")
	var client *github.Client
	if baseUrl != "" {
		client, err = github.NewEnterpriseClient(baseUrl, baseUrl, tc)
		if err != nil {
			panic(fmt.Errorf("error setting ghe client: %s", err))
		}
	} else {
		client = github.NewClient(tc)
	}

	viper.SetConfigName("takolabel_create")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config: %s", err))
	}

	var labels []config.Label
	err = viper.UnmarshalKey("labels", &labels)
	if err != nil {
		panic(fmt.Errorf("error reading config: %s", err))
	}

	repositories := viper.GetStringSlice("repositories")

	for _, repository := range repositories {
		s := strings.Split(repository, "/")
		if len(s) != 2 {
			panic(fmt.Errorf("repository %s is not properly formatted in setting yaml file", repository))
		}
		owner, repo := s[0], s[1]
		for _, label := range labels {
			_, err := takolabel.CreateLabel(ctx, client.Issues, label, owner, repo)
			if err != nil {
				fmt.Printf("error creating label \"%s\" for repository \"%s\": %s\n", label.Name, owner+"/"+repo, err)
			} else {
				fmt.Printf("created label \"%s\" for repository \"%s\"\n", label.Name, owner+"/"+repo)
			}
		}
	}
}
