package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"

	"google.golang.org/grpc"

	"go-micro-site/core/registry"
	"go-micro-site/specs"

	"github.com/go-zoo/bone"
)

var clients *clientMap

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c)

	// register itself
	reg, _ := registry.NewRegistry("site", "v1")
	_ = reg.Register()
	// defer unregister
	defer reg.Unregister()

	// setup clients
	clients = newClientMap()
	defer clients.Close()

	mux := bone.New()
	mux.GetFunc("/pel/:element", renderElementHandler)
	mux.GetFunc("/pel/:element-:version", renderElementHandler)

	go http.ListenAndServe(":80", mux)

	// wait for signal
	select {
	case <-c:
		log.Println("Received a signal, stopping")
	}
}

type clientId struct {
	name, version string
}

type clientValue struct {
	client specs.PageElementClient
	conn   *grpc.ClientConn
}

type clientMap struct {
	storage map[clientId]*clientValue
}

func newClientMap() *clientMap {
	return &clientMap{
		storage: map[clientId]*clientValue{},
	}
}

type renderTreeEl map[clientId]renderTreeEl

type renderedTree map[clientId]renderedTreeContents

type renderedTreeContents struct {
	render *specs.PageRender
	tree   renderedTree
}

func (c *clientMap) GetById(cID clientId) (specs.PageElementClient, error) {
	// return an already existing client if possible
	if client, ok := c.storage[cID]; ok {
		return client.client, nil
	}

	// looks like we need to create a new one
	conn, err := grpc.Dial(fmt.Sprintf("%s.%s.service.consul:80", cID.version, cID.name), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := specs.NewPageElementClient(conn)

	c.storage[cID] = &clientValue{
		client: client,
		conn:   conn,
	}

	return client, nil
}

func (c *clientMap) Get(name, version string) (specs.PageElementClient, error) {
	return c.GetById(clientId{name, version})
}

func (c *clientMap) Close() {
	for _, client := range c.storage {
		client.conn.Close()
	}
}

func renderElementHandler(w http.ResponseWriter, r *http.Request) {
	el := getUrlValue(r, "element", "")
	version := getUrlValue(r, "version", "v1")
	cID := clientId{el, version}

	tree := buildRenderTree(cID)
	log.Println("Got following tree: ", tree)

	render := renderTree(tree)
	log.Println("Rendered the tree: ", render)

	w.Header().Add("Content-Type", "text/html")
	fmt.Fprint(w, parseTree(cID, render))
}

func getUrlValue(r *http.Request, key, def string) string {
	if v := bone.GetValue(r, key); v != "" {
		return v
	}
	return def
}

func buildRenderTree(cID clientId) renderTreeEl {
	// tree := []clientId{clientId{name, version}}
	tree := renderTreeEl{}
	tree[cID] = renderTreeEl{}

	c, err := clients.GetById(cID)
	if err != nil {
		log.Printf("Failed to get client for %v (err = %v)", cID, err)
		return tree
	}

	d, err := c.Describe(context.Background(), &specs.Empty{})
	if err != nil {
		log.Printf("Failed to fetch description of %v (err = %v)", cID, err)
		return tree
	}

	for _, inc := range d.Includes {
		incCID := clientId{inc.Name, "v1"}
		incTree := buildRenderTree(incCID)
		tree[cID][incCID] = incTree[incCID]
	}

	log.Printf("Built tree for %v = %v", cID, tree)

	return tree
}

func renderTree(tree renderTreeEl) renderedTree {
	rtree := renderedTree{}

	for cID, subtree := range tree {
		r, err := renderElement(cID)
		if err != nil {
			log.Println("Failed to render element %v (err = %v)", cID, err)
			continue
		}

		rtree[cID] = renderedTreeContents{r, renderTree(subtree)}
	}

	return rtree
}

func parseTree(cID clientId, tree renderedTree) string {
	cHtml := tree[cID].render.Html

	for subCID, _ := range tree[cID].tree {
		re := regexp.MustCompile(`<element>` + subCID.name + `</element>`)
		log.Printf("we should replace <element>%s</element> with \"%s\"", subCID.name, parseTree(subCID, tree[cID].tree))
		cHtml = re.ReplaceAllString(cHtml, parseTree(subCID, tree[cID].tree))
	}

	return cHtml
}

func renderElement(cID clientId) (*specs.PageRender, error) {
	c, err := clients.Get(cID.name, cID.version)
	if err != nil {
		log.Printf("Failed to get client for %s:%s (err = %v)", cID.name, cID.version, err)
		return nil, err
	}

	return c.Render(context.Background(), &specs.RenderArgs{})
}
