// Copyright (c) 2022 Takayuki NAGATOMI
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
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

// ConfigCreate sets target for takolabel create
type ConfigCreate struct {
	Labels []Label
	Repos  []Repo
}

// ConfigDelete sets target for takolabel delete
type ConfigDelete struct {
	Labels []string
	Repos  []Repo
}

// ConfigSync sets target for takolabel sync
type ConfigSync struct {
	Labels []Label
	Repos  []Repo
}

// ConfigEmpty sets target for takolabel empty
type ConfigEmpty struct {
	Repos []Repo
}

// Label defines properties for GitHub label
type Label struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Color       string `yaml:"color"`
}

// Repo defines owner and repo for GitHub
type Repo struct {
	Owner string
	Repo  string
}

// configCreateYAML is YAML config struct used for takolabel
type configCreateYAML struct {
	Repositories []string `yaml:"repositories"`
	Labels       []Label  `yaml:"labels"`
}

// configDeleteYAML is YAML config struct used for takolabel
type configDeleteYAML struct {
	Repositories []string `yaml:"repositories"`
	Labels       []string `yaml:"labels"`
}

// configDeleteYAML is YAML config struct used for takolabel
type configSyncYAML struct {
	Repositories []string `yaml:"repositories"`
	Labels       []Label  `yaml:"labels"`
}

// Parse config file for takolabel create
func (c *ConfigCreate) Parse(filename string) error {
	f, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("read file failed: %v", err)
	}

	y := configCreateYAML{}
	if err := yaml.Unmarshal(f, &y); err != nil {
		return fmt.Errorf("yaml unmarshal failed: %v", err)
	}

	c.Labels = y.Labels
	for _, r := range y.Repositories {
		s := strings.Split(r, "/")
		if len(s) != 2 {
			return fmt.Errorf("repository %s is not properly formatted in setting yaml file", r)
		}
		c.Repos = append(c.Repos, Repo{
			Owner: s[0],
			Repo:  s[1],
		})
	}
	return nil
}

// Parse config file for takolabel delete
func (c *ConfigDelete) Parse(filename string) error {
	f, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("read file failed: %v", err)
	}

	y := configDeleteYAML{}
	if err := yaml.Unmarshal(f, &y); err != nil {
		return fmt.Errorf("yaml unmarshal failed: %v", err)
	}

	c.Labels = y.Labels
	for _, r := range y.Repositories {
		s := strings.Split(r, "/")
		if len(s) != 2 {
			return fmt.Errorf("repository %s is not properly formatted in setting yaml file", r)
		}
		c.Repos = append(c.Repos, Repo{
			Owner: s[0],
			Repo:  s[1],
		})
	}
	return nil
}

// Parse config file for takolabel sync
func (c *ConfigSync) Parse(filename string) error {
	f, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("read file failed: %v", err)
	}

	y := configSyncYAML{}
	if err := yaml.Unmarshal(f, &y); err != nil {
		return fmt.Errorf("yaml unmarshal failed: %v", err)
	}

	c.Labels = y.Labels
	for _, r := range y.Repositories {
		s := strings.Split(r, "/")
		if len(s) != 2 {
			return fmt.Errorf("repository %s is not properly formatted in setting yaml file", r)
		}
		c.Repos = append(c.Repos, Repo{
			Owner: s[0],
			Repo:  s[1],
		})
	}
	return nil
}

// Parse config file for takolabel empty
func (c *ConfigEmpty) Parse(filename string) error {
	f, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("read file failed: %v", err)
	}

	y := configDeleteYAML{}
	if err := yaml.Unmarshal(f, &y); err != nil {
		return fmt.Errorf("yaml unmarshal failed: %v", err)
	}

	for _, r := range y.Repositories {
		s := strings.Split(r, "/")
		if len(s) != 2 {
			return fmt.Errorf("repository %s is not properly formatted in setting yaml file", r)
		}
		c.Repos = append(c.Repos, Repo{
			Owner: s[0],
			Repo:  s[1],
		})
	}
	return nil
}
