package main

import (
	"golang.org/x/net/context"

	"go-micro-site/core/pageElement"
	"go-micro-site/specs"
)

var element *pageElement.Element

const (
	elementName    = "page-home"
	elementVersion = "v1"
)

type server struct{}

func (s *server) Describe(_ context.Context, _ *specs.Empty) (*specs.PageElementDescription, error) {
	return &specs.PageElementDescription{
		Name:    elementName,
		Version: elementVersion,
		Includes: []*specs.PageElementIncludes{
			&specs.PageElementIncludes{Name: "header"},
			&specs.PageElementIncludes{Name: "footer"},
			&specs.PageElementIncludes{Name: "recommendations"},
		},
	}, nil
}

func (s *server) Render(_ context.Context, _ *specs.RenderArgs) (*specs.PageRender, error) {
	r, err := element.Render("template.html", nil)
	if err != nil {
		return r, err
	}

	r.PageTitle = "Micro Home Page"
	return r, nil
}

func main() {
	element = pageElement.MustNew(&pageElement.Config{
		Name:         elementName,
		Version:      elementVersion,
		AssetFunc:    Asset,
		AssetDirFunc: AssetDir,
		Server:       new(server),
	})
	element.Run()
}
