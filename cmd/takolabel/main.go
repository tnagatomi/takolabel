package main

import (
	"context"
	"flag"
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
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	createDryRun := createCmd.Bool("dry-run", false, "execute dry-run")

	switch os.Args[1] {
	case "create":
		createCmd.Parse(os.Args[2:])
		target := takolabel.GatherCreate()
		if *createDryRun {
			takolabel.DryRunCreate(target)
		} else {
			takolabel.ExecuteCreate(ctx, client, target)
		}
	case "delete":
		if confirm() {
			target := takolabel.GatherDelete()
			takolabel.ExecuteDelete(ctx, client, target)
		} else {
			os.Exit(0)
		}
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

func confirm() bool {
	var response string
	fmt.Printf("are you sure you want to do this? (y/n): ")
	_, err := fmt.Scan(&response)
	if err != nil {
		os.Exit(1)
	}
	if response == "y" {
		return true
	}
	return false
}
