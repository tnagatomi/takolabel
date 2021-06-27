package takolabel

import (
	"reflect"
	"testing"
)

func TestDeleteParse(t *testing.T) {
	delete_ := Delete{}
	err := delete_.Parse([]byte(`repositories:
  - some-owner/some-owner-repo-1
  - some-owner/some-owner-repo-2
  - another-owner/another-owner-repo-1
labels:
  - Label 1
  - Label 2
  - Label 3
`))
	if err != nil {
		t.Fatalf("error: %q", err)
	}

	want := DeleteTarget{
		Repositories: []string{
			"some-owner/some-owner-repo-1",
			"some-owner/some-owner-repo-2",
			"another-owner/another-owner-repo-1",
		},
		Labels: []string{
			"Label 1",
			"Label 2",
			"Label 3",
		},
	}

	if !reflect.DeepEqual(delete_.Target, want) {
		t.Errorf("got %v want %v", delete_.Target, want)
	}
}
