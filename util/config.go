package util

type Repository struct {
	Org  string `mapstructure:"org"`
	Repo string `mapstructure:"repo"`
}

type Label struct {
	Name        string `mapstructure:"name"`
	Description string `mapstructure:"description"`
	Color       string `mapstructure:"color"`
}
