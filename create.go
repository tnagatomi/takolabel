package takolabel

import (
	"context"
	"fmt"
	"github.com/google/go-github/v33/github"
	"github.com/spf13/viper"
	"github.com/tommy6073/takolabel/config"
	"strings"
)

func GatherCreate() Target {
	viper.SetConfigName("takolabel_create")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config: %s", err))
	}

	var labels []config.Label
	err = viper.UnmarshalKey("labels", &labels)
	if err != nil {
		panic(fmt.Errorf("error reading config: %s", err))
	}

	repositories := viper.GetStringSlice("repositories")
	return Target{repositories, labels}
}

func ExecuteCreate(ctx context.Context, client *github.Client, target Target) {
	for _, repository := range target.repositories {
		s := strings.Split(repository, "/")
		if len(s) != 2 {
			panic(fmt.Errorf("repository %s is not properly formatted in setting yaml file", repository))
		}
		owner, repo := s[0], s[1]
		for _, label := range target.labels {
			_, err := CreateLabel(ctx, client.Issues, label, owner, repo)
			if err != nil {
				fmt.Printf("error creating label \"%s\" for repository \"%s\": %s\n", label.Name, owner+"/"+repo, err)
			} else {
				fmt.Printf("created label \"%s\" for repository \"%s\"\n", label.Name, owner+"/"+repo)
			}
		}
	}
}
