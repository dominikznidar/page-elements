package main

import (
	"go-micro-site/core/registry"
	"go-micro-site/core/render"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-zoo/bone"
)

var t *template.Template

type State map[string]*struct {
	Enabled string
	Tags    []string
}

func main() {
	// load templates
	t = render.GetTemplates(AssetDir, Asset)

	// register itself
	reg, _ := registry.NewRegistry("dashboard", "v1", true)
	_ = reg.Register()
	defer reg.Unregister()

	mux := bone.New()
	mux.GetFunc("/", renderPageHandler(reg))
	mux.PostFunc("/save", storeStateHandler(reg))

	go http.ListenAndServe(":80", mux)

	log.Println("Ready to serve http requests ...")

	// intercept shutdown signals
	c := make(chan os.Signal, 1)
	signal.Notify(c)

	// wait for signal
	select {
	case <-c:
		log.Println("Received a signal, stopping")
	}
}

func renderPageHandler(reg *registry.Registry) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get state
		state, _ := reg.FetchCurrentState()
		// get available clients
		clients, _ := reg.FetchAvailableClients()

		log.Printf("Current state = %v", state)
		log.Printf("Available clients = %v", clients)

		// join state and clients toggether

		w.Header().Set("Content-type", "text/html")

		err := t.Execute(w, composeState(state, clients))
		if err != nil {
			log.Printf("Failed to render the template; err = %v", err)
		}
	}
}

func storeStateHandler(reg *registry.Registry) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// update state
		newState := registry.State{}
		r.ParseForm()
		for serviceName, values := range r.Form {
			log.Printf("received service %s with values %v", serviceName, values)
			newState[serviceName] = values[0]
		}

		log.Printf("Composed a new state: %v", newState)
		reg.UpdateState(newState)

		// redirect back to the list
		http.Redirect(w, r, "/?done=1", 302)
	}
}

func composeState(state *registry.State, clients registry.Clients) *State {
	s := &State{}

	for name, tags := range clients {
		activeTag, _ := (*state)[name]
		(*s)[name] = &struct {
			Enabled string
			Tags    []string
		}{
			Enabled: activeTag,
			Tags:    tags,
		}
	}

	return s
}
