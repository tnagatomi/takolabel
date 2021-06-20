package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v33/github"
	"github.com/spf13/viper"
	"github.com/tommy6073/takolabel"
	"golang.org/x/oauth2"
	"os"
)

func main() {
	ctx := context.Background()

	client := getGitHubClient(ctx)

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "expected subcommands\n")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		target := takolabel.GatherCreate()
		takolabel.ExecuteCreate(ctx, client, target)
	}
}

func getGitHubClient(ctx context.Context) *github.Client {
	viper.SetConfigName("takolabel")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading config: %s\n", err)
		os.Exit(1)
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
			fmt.Fprintf(os.Stderr, "error setting ghe client: %s\n", err)
			os.Exit(1)
		}
	} else {
		client = github.NewClient(tc)
	}
	return client
}
