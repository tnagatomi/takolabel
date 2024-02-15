// Copyright (c) 2021 Takayuki NAGATOMI
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package takolabel

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestConfigCreateParse(t *testing.T) {
	c := ConfigCreate{}
	err := c.Parse("testdata/takolabel_create.yml")
	if err != nil {
		t.Fatalf("error parsing config: %v", err)
	}
	want := ConfigCreate{
		Repos: []Repo{
			{"some-owner", "some-owner-repo-1"},
			{"some-owner", "some-owner-repo-2"},
			{"another-owner", "another-owner-repo-1"},
		},
		Labels: []Label{
			{
				Name:  "Label 1",
				Color: "ff0000",
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

	if diff := cmp.Diff(want, c); diff != "" {
		t.Errorf("Parse() mismatch (-want +got):\n%s", diff)
	}
}

func TestConfigDeleteParse(t *testing.T) {
	c := ConfigDelete{}
	err := c.Parse("testdata/takolabel_delete.yml")
	if err != nil {
		t.Fatalf("error parsing config: %v", err)
	}
	want := ConfigDelete{
		Repos: []Repo{
			{"some-owner", "some-owner-repo-1"},
			{"another-owner", "another-owner-repo-1"},
		},
		Labels: []string{
			"Label 1",
			"Label 2",
			"Label 3",
		},
	}

	if diff := cmp.Diff(want, c); diff != "" {
		t.Errorf("Parse() mismatch (-want +got):\n%s", diff)
	}
}

func TestConfigSyncParse(t *testing.T) {
	c := ConfigSync{}
	err := c.Parse("testdata/takolabel_sync.yml")
	if err != nil {
		t.Fatalf("error parsing config: %v", err)
	}
	want := ConfigSync{
		Repos: []Repo{
			{"some-owner", "some-owner-repo-1"},
			{"some-owner", "some-owner-repo-2"},
			{"another-owner", "another-owner-repo-1"},
		},
		Labels: []Label{
			{
				Name:  "Label 1",
				Color: "ff0000",
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

	if diff := cmp.Diff(want, c); diff != "" {
		t.Errorf("Parse() mismatch (-want +got):\n%s", diff)
	}
}

func TestConfigEmptyParse(t *testing.T) {
	c := ConfigEmpty{}
	err := c.Parse("testdata/takolabel_delete.yml")
	if err != nil {
		t.Fatalf("error parsing config: %v", err)
	}
	want := ConfigEmpty{
		Repos: []Repo{
			{"some-owner", "some-owner-repo-1"},
			{"another-owner", "another-owner-repo-1"},
		},
	}

	if diff := cmp.Diff(want, c); diff != "" {
		t.Errorf("Parse() mismatch (-want +got):\n%s", diff)
	}
}
