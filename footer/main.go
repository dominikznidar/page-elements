package main

import (
	"golang.org/x/net/context"

	"page-elements/core/pageElement"
	"page-elements/specs"
)

var element *pageElement.Element

const (
	elementName    = "footer"
	elementVersion = "v1"
)

type server struct{}

func (s *server) Describe(_ context.Context, _ *specs.RenderArgs) (*specs.PageElementDescription, error) {
	return &specs.PageElementDescription{
		Name:    elementName,
		Version: elementVersion,
	}, nil
}

func (s *server) Render(_ context.Context, _ *specs.RenderArgs) (*specs.PageRender, error) {
	return element.Render("template.html", struct {
		Year string
	}{"2016"})
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
