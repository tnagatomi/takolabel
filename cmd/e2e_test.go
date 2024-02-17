package cmd_test

import (
	"bytes"
	"context"
	"github.com/google/go-github/v41/github"
	"github.com/tnagatomi/takolabel/cmd"
	"golang.org/x/oauth2"
	"os"
	"path/filepath"
	"testing"
)

var (
	owner = "tnagatomi"
	repo1 = "test-repository-for-takolabel-1"
	repo2 = "test-repository-for-takolabel-2"
)

func TestE2E(t *testing.T) {
	defer changeWorkingDir(t, filepath.Join("..", "testdata"))()

	confirmInput := bytes.NewBufferString("y\n")

	emptyCmd := cmd.NewDeleteCmd(confirmInput)

	// Ensure there are no labels in the repositories
	err := emptyCmd.Execute()
	if err != nil {
		t.Fatalf("failed to execute empty command %v", err)
	}

	createCmd := cmd.NewCreateCmd()

	err = createCmd.Execute()
	if err != nil {
		t.Fatalf("failed to execute create command %v", err)
	}

	client := newGitHubClient(t)

	labels, _, err := client.Issues.ListLabels(context.Background(), owner, repo1, nil)
	if err != nil {
		t.Fatalf("failed to get labels from %s: %v", repo1, err)
	}
	if len(labels) != 3 {
		t.Errorf("expected 3 labels for %q, but got %d", repo1, len(labels))
	}
	if *labels[0].Name != "Label 1" && *labels[0].Color != "ff0000" && *labels[0].Description != "" {
		t.Errorf("expected label is name: Label 1 color: ff0000 for %q but got name: %s color: %s", repo1, *labels[0].Name, *labels[0].Color)
	}
	if *labels[1].Name != "Label 2" && *labels[1].Color != "00ff00" && *labels[1].Description != "This is the label two by create" {
		t.Errorf("expected label is name: Label 2 color: 00ff00 description: This is the label two by create for %q but got name: %s color: %s description: %s", repo1, *labels[1].Name, *labels[1].Color, *labels[1].Description)
	}
	if *labels[2].Name != "Label 3" && *labels[2].Color != "0000ff" && *labels[2].Description != "This is the label three by create" {
		t.Errorf("expected label is name: Label 3 color: 0000ff description: This is the label three by create for %q but got name: %s color: %s description: %s", repo1, *labels[2].Name, *labels[2].Color, *labels[2].Description)
	}

	labels, _, err = client.Issues.ListLabels(context.Background(), owner, repo2, nil)
	if err != nil {
		t.Fatalf("failed to get labels from %s: %v", repo2, err)
	}
	if len(labels) != 3 {
		t.Errorf("expected 3 labels for %q, but got %d", repo2, len(labels))
	}
	if *labels[0].Name != "Label 1" && *labels[0].Color != "ff0000" && *labels[0].Description != "" {
		t.Errorf("expected label is name: Label 1 color: ff0000 for %q but got name: %s color: %s", repo2, *labels[0].Name, *labels[0].Color)
	}
	if *labels[1].Name != "Label 2" && *labels[1].Color != "00ff00" && *labels[1].Description != "This is the label two by create" {
		t.Errorf("expected label is name: Label 2 color: 00ff00 description: This is the label two by create for %q but got name: %s color: %s description: %s", repo2, *labels[1].Name, *labels[1].Color, *labels[1].Description)
	}
	if *labels[2].Name != "Label 3" && *labels[2].Color != "0000ff" && *labels[2].Description != "This is the label three by create" {
		t.Errorf("expected label is name: Label 3 color: 0000ff description: This is the label three by create for %q but got name: %s color: %s description: %s", repo2, *labels[2].Name, *labels[2].Color, *labels[2].Description)
	}

	confirmInput = bytes.NewBufferString("y\n")
	deleteCmd := cmd.NewDeleteCmd(confirmInput)

	err = deleteCmd.Execute()
	if err != nil {
		t.Fatalf("failed to execute delete command %v", err)
	}

	labels, _, err = client.Issues.ListLabels(context.Background(), owner, repo1, nil)
	if err != nil {
		t.Fatalf("failed to get labels from %q: %v", repo1, err)
	}
	if len(labels) != 2 {
		t.Errorf("expected 2 labels for %q, but got %d", repo1, len(labels))
	}

	labels, _, err = client.Issues.ListLabels(context.Background(), owner, repo2, nil)
	if err != nil {
		t.Fatalf("failed to get labels from %q: %v", repo2, err)
	}
	if len(labels) != 2 {
		t.Errorf("expected 2 labels for %q, but got %d", repo2, len(labels))
	}

	syncCmd := cmd.NewSyncCmd()

	err = syncCmd.Execute()
	if err != nil {
		t.Fatalf("failed to execute create command %v", err)
	}

	labels, _, err = client.Issues.ListLabels(context.Background(), owner, repo1, nil)
	if err != nil {
		t.Fatalf("failed to get labels from %s: %v", repo1, err)
	}
	if len(labels) != 3 {
		t.Errorf("expected 3 labels for %q, but got %d", repo1, len(labels))
	}
	if *labels[0].Name != "Label 1" && *labels[0].Color != "ff0000" && *labels[0].Description != "" {
		t.Errorf("expected label is name: Label 1 color: ff0000 for %q but got name: %s color: %s", repo1, *labels[0].Name, *labels[0].Color)
	}
	if *labels[1].Name != "Label 2" && *labels[1].Color != "00ff00" && *labels[1].Description != "This is the label two by sync" {
		t.Errorf("expected label is name: Label 2 color: 00ff00 description: This is the label two by create for %q but got name: %s color: %s description: %s", repo1, *labels[1].Name, *labels[1].Color, *labels[1].Description)
	}
	if *labels[2].Name != "Label 3" && *labels[2].Color != "0000ff" && *labels[2].Description != "This is the label three by sync" {
		t.Errorf("expected label is name: Label 3 color: 0000ff description: This is the label three by create for %q but got name: %s color: %s description: %s", repo1, *labels[2].Name, *labels[2].Color, *labels[2].Description)
	}

	labels, _, err = client.Issues.ListLabels(context.Background(), owner, repo2, nil)
	if err != nil {
		t.Fatalf("failed to get labels from %s: %v", repo2, err)
	}
	if len(labels) != 3 {
		t.Errorf("expected 3 labels for %q, but got %d", repo2, len(labels))
	}
	if *labels[0].Name != "Label 1" && *labels[0].Color != "ff0000" && *labels[0].Description != "" {
		t.Errorf("expected label is name: Label 1 color: ff0000 for %q but got name: %s color: %s", repo2, *labels[0].Name, *labels[0].Color)
	}
	if *labels[1].Name != "Label 2" && *labels[1].Color != "00ff00" && *labels[1].Description != "This is the label two by sync" {
		t.Errorf("expected label is name: Label 2 color: 00ff00 description: This is the label two by create for %q but got name: %s color: %s description: %s", repo2, *labels[1].Name, *labels[1].Color, *labels[1].Description)
	}
	if *labels[2].Name != "Label 3" && *labels[2].Color != "0000ff" && *labels[2].Description != "This is the label three by sync" {
		t.Errorf("expected label is name: Label 3 color: 0000ff description: This is the label three by create for %q but got name: %s color: %s description: %s", repo2, *labels[2].Name, *labels[2].Color, *labels[2].Description)
	}

	confirmInput = bytes.NewBufferString("y\n")
	emptyCmd2 := cmd.NewEmptyCmd(confirmInput)
	err = emptyCmd2.Execute()
	if err != nil {
		t.Fatalf("failed to execute empty command %v", err)
	}

	labels, _, err = client.Issues.ListLabels(context.Background(), owner, repo1, nil)
	if err != nil {
		t.Fatalf("failed to get labels from %q: %v", repo1, err)
	}
	if len(labels) != 0 {
		t.Errorf("expected 0 labels for %q, but got %d", repo1, len(labels))
	}

	labels, _, err = client.Issues.ListLabels(context.Background(), owner, repo2, nil)
	if err != nil {
		t.Fatalf("failed to get labels from %q: %v", repo2, err)
	}
	if len(labels) != 0 {
		t.Errorf("expected 0 labels for %q, but got %d", repo2, len(labels))
	}
}

// newGitHubClient is helper function to create new GitHub client
func newGitHubClient(t *testing.T) *github.Client {
	t.Helper()

	githubToken := os.Getenv("TAKOLABEL_TOKEN")
	if githubToken == "" {
		t.Fatalf("TAKOLABEL_TOKEN environment variable is not set")
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	return github.NewClient(tc)
}

// changeWorkingDir is helper function to change working directory for test
func changeWorkingDir(t *testing.T, newDir string) func() {
	t.Helper()

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current directory: %s", err)
	}

	err = os.Chdir(newDir)
	if err != nil {
		t.Fatalf("failed to change directory to %s: %s", newDir, err)
	}

	return func() {
		err := os.Chdir(originalDir)
		if err != nil {
			t.Fatalf("failed to change back to original directory: %s", err)
		}
	}
}
