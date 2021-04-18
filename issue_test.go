package takolabel

import (
	"context"
	"github.com/google/go-github/v33/github"
	"reflect"
	"takolabel/config"
	"testing"
)

type mockService struct {
}

func (ms *mockService) CreateLabel(_ context.Context, _ string, _ string, label *github.Label) (*github.Label, *github.Response, error) {
	return &github.Label{Name: label.Name, Description: label.Description, Color: label.Color}, nil, nil
}

func TestCreateLabels(t *testing.T) {
	ctx := context.Background()
	repositories := []config.Repository{
		{
			Org:  "some-org",
			Repo: "some-org-repo-1",
		},
		{
			Org:  "some-org",
			Repo: "some-org-repo-2",
		},
		{
			Org:  "another-org",
			Repo: "another-org-repo-1",
		},
	}
	labels := []config.Label{
		{
			Name:        "Label 1",
			Description: "This is the label one",
			Color:       "ff0000",
		},
		{
			Name:        "Label 2",
			Description: "This is the label two",
			Color:       "00ff00",
		},
		{
			Name:        "Label 3",
			Description: "This is the label three",
			Color:       "0000",
		},
	}
	got := CreateLabels(&IssuesClient{Ctx: ctx, IssuesService: &mockService{}}, repositories, labels)
	want := 9

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestCreateLabel(t *testing.T) {
	ctx := context.Background()
	ms := &mockService{}
	ic := &IssuesClient{Ctx: ctx, IssuesService: ms}
	createdLabel, err := ic.CreateLabel(
		config.Label{Name: "Label 1", Description: "This is the label one", Color: "ff0000"},
		config.Repository{Org: "org", Repo: "repository"},
	)

	if err != nil {
		t.Fatalf("CreateLabel: %v", err)
	}

	got := createdLabel
	want := &github.Label{
		Name:        github.String("Label 1"),
		Description: github.String("This is the label one"),
		Color:       github.String("ff0000"),
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
