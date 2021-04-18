package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v33/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"takolabel/util"
)

func main() {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config: %s", err))
	}

	viper.SetConfigName("labels")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config: %s", err))
	}

	githubToken := viper.GetString("GITHUB_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)
	baseUrl := viper.GetString("BASE_URL")
	client, err := github.NewEnterpriseClient(baseUrl, baseUrl, tc)
	if err != nil {
		panic(fmt.Errorf("error setting ghe client: %s", err))
	}

	var labels []util.Label
	err = viper.UnmarshalKey("labels", &labels)
	if err != nil {
		panic(fmt.Errorf("error reading config: %s", err))
	}

	var repositories []util.Repository
	err = viper.UnmarshalKey("repositories", &repositories)
	if err != nil {
		panic(fmt.Errorf("error reading config: %s", err))
	}

	for _, repository := range repositories {
		for _, label := range labels {
			githubLabel := &github.Label{
				Name:        github.String(label.Name),
				Description: github.String(label.Description),
				Color:       github.String(label.Color),
			}
			_, _, err = client.Issues.CreateLabel(ctx, repository.Org, repository.Repo, githubLabel)
			if err != nil {
				fmt.Printf("error creating label \"%s\" for repository \"%s\": %s\n", label.Name, repository.Org+"/"+repository.Repo, err)
			} else {
				fmt.Printf("created label \"%s\" for repository \"%s\"\n", label.Name, repository.Org+"/"+repository.Repo)
			}
		}
	}
}
