package pageElement

import (
	"bytes"
	"go-micro-site/core/registry"
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
	e.t = getTemplates(e.c)

	return e
}

func getTemplates(c *Config) *template.Template {
	var t, tmpl *template.Template

	// find all tempates
	views, err := c.AssetDirFunc("templates")
	if err != nil {
		log.Fatal("Failed to get views;", err)
	}

	// parse all the templates
	for _, v := range views {
		// initialize t if not done yet
		if t == nil {
			t = template.New(v)
		}
		// get pointer to current template
		if v == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(v)
		}

		// get the contents of the current template
		b, err := c.AssetFunc("templates/" + v)
		if err != nil {
			log.Fatal("Failed to read a template file;", err)
		}

		// parse the template
		_, err = tmpl.Parse(string(b))
		if err != nil {
			log.Fatal("Failed to parse a template;", err)
		}
	}

	return t
}
