package takolabel

import "github.com/tommy6073/takolabel/config"

type CreateTarget struct {
	repositories []string
	labels       []config.Label
}

type DeleteTarget struct {
	repositories []string
	labels       []string
}
