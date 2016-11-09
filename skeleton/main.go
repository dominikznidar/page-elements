package main

import (
	"golang.org/x/net/context"

	"page-elements/core/pageElement"
	"page-elements/specs"
)

var element *pageElement.Element

const (
	elementName    = "skeleton"
	elementVersion = "v1"
)

type server struct{}

func (s *server) Describe(_ context.Context, args *specs.RenderArgs) (*specs.PageElementDescription, error) {
	includes := []*specs.PageElementIncludes{}

	if el := args.Get("element"); el != "" {
		includes = append(includes, &specs.PageElementIncludes{Name: el})
	}

	return &specs.PageElementDescription{
		Name:     elementName,
		Version:  elementVersion,
		Includes: includes,
	}, nil
}

func (s *server) Render(_ context.Context, args *specs.RenderArgs) (*specs.PageRender, error) {
	return element.Render("template.html", args)
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
