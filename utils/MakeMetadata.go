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

type Brand struct {
	Title       string
	Description string
}

type MetadataInput struct {
	Brand       Brand
	Locale      string
	Title       string
	CurrentURL  string
	Description string
	Breadcrumbs BreadcrumbsInput
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

	brand := Brand{
		Title:       meta.Brand.Title,
		Description: meta.Brand.Description,
	}

	titleTemplate := "%s | " + brand.Title

	description := meta.Description
	if description == "" {
		description = brand.Description
	}

	return MetadataOutput{
		Locale:      meta.Locale,
		Title:       fmt.Sprintf(titleTemplate, meta.Title),
		Description: description,
		Breadcrumbs: breadcrumbs,
		BaseURL:     baseUrl,
		CurrentURL:  baseUrl + meta.CurrentURL,
		Brand:       brand,
	}
}
