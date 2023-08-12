package utils

import (
	"fmt"
	"os"
)

type BreadcrumbsInput [][2]string
type BreadcrumbsOutput struct {
	Position int
	Url      string
	Title    string
}

type MetadataInput struct {
	Locale      string
	Title       string
	CurrentURL  string
	Description string
	Breadcrumbs BreadcrumbsInput
}

type Brand struct {
	Title       string
	Description string
}

type MetadataOutput struct {
	Locale      string
	Brand       Brand
	Title       string
	BaseURL     string
	CurrentURL  string
	Description string
	Breadcrumbs []BreadcrumbsOutput
}

var titleTemplate = "%s | Акигами"

func MakeMetadata(meta MetadataInput) MetadataOutput {

	breadcrumbs := []BreadcrumbsOutput{}

	baseUrl := os.Getenv("BASE_URL")

	for i, breadcrumb := range meta.Breadcrumbs {
		breadcrumbs = append(breadcrumbs, BreadcrumbsOutput{
			Position: i + 1,
			Url:      baseUrl + breadcrumb[0],
			Title:    breadcrumb[1],
		})
	}

	return MetadataOutput{
		Locale:      meta.Locale,
		Title:       fmt.Sprintf(titleTemplate, meta.Title),
		Description: meta.Description,
		Breadcrumbs: breadcrumbs,
		BaseURL:     baseUrl,
		CurrentURL:  baseUrl + meta.CurrentURL,
		Brand: Brand{
			Title:       "Акигами",
			Description: "Какое-то логичное описание",
		},
	}
}
