package config

type Label struct {
	Name        string `mapstructure:"name"`
	Description string `mapstructure:"description"`
	Color       string `mapstructure:"color"`
}
