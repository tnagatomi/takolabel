package takolabel

import (
	"context"
	"github.com/google/go-github/v33/github"
	"github.com/tommy6073/takolabel/config"
	"reflect"
	"testing"
)

type mockService struct {
}

func (ms *mockService) CreateLabel(_ context.Context, _ string, _ string, label *github.Label) (*github.Label, *github.Response, error) {
	return &github.Label{Name: label.Name, Description: label.Description, Color: label.Color}, nil, nil
}

func TestCreateLabel(t *testing.T) {
	ctx := context.Background()
	ms := &mockService{}
	createdLabel, err := CreateLabel(
		ctx,
		ms,
		config.Label{Name: "Label 1", Description: "This is the label one", Color: "ff0000"},
		config.Repository{Org: "org", Repo: "repository"},
	)

	//if err != nil {
	//	t.Fatalf("CreateLabel: %v", err)
	//}

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
