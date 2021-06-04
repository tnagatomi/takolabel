package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v33/github"
	"github.com/spf13/viper"
	"github.com/tommy6073/takolabel"
	"golang.org/x/oauth2"
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

	target := takolabel.GatherCreate()
	takolabel.ExecuteCreate(ctx, client, target)
}
