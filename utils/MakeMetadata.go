package utils

import (
	"fmt"
)

type Metadata struct {
	Title       string
	Description string
}

var titleTemplate = "%s | Акигами"

func MakeMetadata(meta Metadata) Metadata {
	return Metadata{
		Title:       fmt.Sprintf(titleTemplate, meta.Title),
		Description: meta.Description,
	}
}
