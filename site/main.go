package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"go-micro-site/core/registry"
	"go-micro-site/specs"

	"github.com/go-zoo/bone"
)

var clients *clientConnections

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c)

	// register itself
	reg, _ := registry.NewRegistry("site", "v1", true)
	_ = reg.Register()
	defer reg.Unregister()

	// setup clients
	clients = newClientConnections()
	defer clients.Close()

	// initialize state
	initState(reg)

	mux := bone.New()
	mux.GetFunc("/pel/:element", renderElementHandler)
	mux.GetFunc("/pel/:element/:version", renderElementHandler)

	go http.ListenAndServe(":80", mux)

	// wait for signal
	select {
	case <-c:
		log.Println("Received a signal, stopping")
	}
}

func renderElementHandler(w http.ResponseWriter, r *http.Request) {
	el := getUrlValue(r, "element", "")
	version := getUrlValue(r, "version", getVersionFor(el))
	format := getQueryValue(r, "format", "pretty")

	var (
		cID clientId
	)
	args := getFormValuesAsRenderArgs(r)

	if format == "pretty" {
		cID = getActiveClientIdFor("skeleton")
		args.Add("element", el)
	} else if format == "snippet" {
		cID = clientId{el, version}
	} else {
		http.Error(w, "Invalid format", http.StatusBadRequest)
		return
	}

	renderedHtml := render(cID, args)

	w.Header().Add("Content-Type", "text/html")
	fmt.Fprint(w, renderedHtml)
}

func getUrlValue(r *http.Request, key, def string) string {
	if v := bone.GetValue(r, key); v != "" {
		return v
	}
	return def
}

func getQueryValue(r *http.Request, key, def string) string {
	if v := r.URL.Query().Get(key); v != "" {
		return v
	}
	return def
}

func getFormValuesAsRenderArgs(r *http.Request) *specs.RenderArgs {
	r.ParseMultipartForm(2 << 20) // 2 MB
	return specs.RenderArgsFromUrlValues(r.Form)
}
