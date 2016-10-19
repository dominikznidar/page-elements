package pageElement

import (
	"bytes"
	"go-micro-site/core/registry"
	"go-micro-site/core/render"
	"go-micro-site/specs"
	"html/template"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

type Element struct {
	t *template.Template
	c *Config
}

type Config struct {
	Name         string
	Version      string
	AssetFunc    func(name string) ([]byte, error)
	AssetDirFunc func(name string) ([]string, error)
	Server       specs.PageElementServer
	SystemTool   bool
}

func (e *Element) Run() {
	c := make(chan os.Signal, 1)
	signal.Notify(c)

	// register itself
	reg, _ := registry.NewRegistry(e.c.Name, e.c.Version, e.c.SystemTool)
	_ = reg.Register()
	// defer unregister
	defer reg.Unregister()

	// setup grpc server
	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Panic("Failed to listen;", err)
	}
	grpcServer := grpc.NewServer()
	specs.RegisterPageElementServer(grpcServer, e.c.Server)
	go grpcServer.Serve(lis)

	// wait for signal
	select {
	case <-c:
		log.Println("Received a signal, stopping")
	}
}

func (e *Element) Render(template string, data interface{}) (*specs.PageRender, error) {
	buf := &bytes.Buffer{}
	err := e.t.ExecuteTemplate(buf, template, data)
	if err != nil {
		return nil, err
	}

	return &specs.PageRender{
		Html: buf.String(),
	}, nil
}

func MustNew(c *Config) *Element {
	// Init element
	e := &Element{c: c}

	// Parse templates
	e.t = render.GetTemplates(e.c.AssetDirFunc, e.c.AssetFunc)

	return e
}
