package main

import (
	"golang.org/x/net/context"

	"go-micro-site/core/pageElement"
	"go-micro-site/specs"
)

var element *pageElement.Element

const (
	elementName    = "header"
	elementVersion = "v1"
)

type server struct{}

func (s *server) Describe(_ context.Context, _ *specs.RenderArgs) (*specs.PageElementDescription, error) {
	return &specs.PageElementDescription{
		Name:    elementName,
		Version: elementVersion,
		Includes: []*specs.PageElementIncludes{
			&specs.PageElementIncludes{Name: "navigation"},
		},
	}, nil
}

func (s *server) Render(_ context.Context, _ *specs.RenderArgs) (*specs.PageRender, error) {
	return element.Render("template.html", nil)
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
