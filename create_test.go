package takolabel

import (
	"reflect"
	"testing"
)

func TestCreateParse(t *testing.T) {
	create := Create{}
	err := create.Parse([]byte(`repositories:
  - some-owner/some-owner-repo-1
  - some-owner/some-owner-repo-2
  - another-owner/another-owner-repo-1
labels:
  - name: Label 1
    description: This is the label one 
    color: ff0000
  - name: Label 2
    description: This is the label two
    color: 00ff00
  - name: Label 3
    description: This is the label three
    color: 0000ff
`))
	if err != nil {
		t.Fatalf("error: %q", err)
	}
	want := CreateTarget{
		Repositories: []Repository{
			{"some-owner", "some-owner-repo-1"},
			{"some-owner", "some-owner-repo-2"},
			{"another-owner", "another-owner-repo-1"},
		},
		Labels: []Label{
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
				Color:       "0000ff",
			},
		},
	}

	if !reflect.DeepEqual(create.Target, want) {
		t.Errorf("got %v want %v", create.Target, want)
	}
}
