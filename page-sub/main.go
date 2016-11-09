package main

import (
	"golang.org/x/net/context"

	"page-elements/core/pageElement"
	"page-elements/specs"
)

var element *pageElement.Element

const (
	elementName    = "page-sub"
	elementVersion = "v1"
)

type server struct{}

func (s *server) Describe(_ context.Context, _ *specs.RenderArgs) (*specs.PageElementDescription, error) {
	return &specs.PageElementDescription{
		Name:    elementName,
		Version: elementVersion,
		Includes: []*specs.PageElementIncludes{
			&specs.PageElementIncludes{Name: "header"},
			&specs.PageElementIncludes{Name: "footer"},
		},
	}, nil
}

func (s *server) Render(_ context.Context, _ *specs.RenderArgs) (*specs.PageRender, error) {
	r, err := element.Render("template.html", nil)
	if err != nil {
		return r, err
	}

	r.PageTitle = "Subpage"
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
