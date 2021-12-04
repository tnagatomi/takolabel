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
	"context"
	"github.com/google/go-github/v33/github"
)

func CreateLabel(ctx context.Context, issuesService *github.IssuesService, label Label, owner string, repo string) (*github.Label, error) {
	githubLabel := &github.Label{
		Name:        github.String(label.Name),
		Description: github.String(label.Description),
		Color:       github.String(label.Color),
	}
	createdLabel, _, err := issuesService.CreateLabel(ctx, owner, repo, githubLabel)
	return createdLabel, err
}

func DeleteLabel(ctx context.Context, issuesService *github.IssuesService, label string, owner string, repo string) error {
	_, err := issuesService.DeleteLabel(ctx, owner, repo, label)
	return err
}
