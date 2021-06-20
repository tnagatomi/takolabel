package takolabel

import (
	"context"
	"fmt"
	"github.com/google/go-github/v33/github"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func GatherDelete() DeleteTarget {
	viper.SetConfigName("takolabel_delete")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.MergeInConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading config: %s\n", err)
		os.Exit(1)
	}

	repositories := viper.GetStringSlice("repositories")
	labels := viper.GetStringSlice("labels")
	return DeleteTarget{repositories, labels}
}

func ExecuteDelete(ctx context.Context, client *github.Client, target DeleteTarget) {
	for _, repository := range target.repositories {
		s := strings.Split(repository, "/")
		if len(s) != 2 {
			fmt.Fprintf(os.Stderr, "repository %s is not properly formatted in setting yaml file\n", repository)
			os.Exit(1)
		}
		owner, repo := s[0], s[1]
		for _, label := range target.labels {
			err := DeleteLabel(ctx, client.Issues, label, owner, repo)
			if err != nil {
				fmt.Printf("error deleting label \"%s\" for repository \"%s\": %s\n", label, owner+"/"+repo, err)
			} else {
				fmt.Printf("deleted label \"%s\" for repository \"%s\"\n", label, owner+"/"+repo)
			}
		}
	}
}
