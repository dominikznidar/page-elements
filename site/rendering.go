package main

import (
	"context"
	"go-micro-site/specs"
	"log"
	"regexp"
)

const (
	statusReady    = iota
	statusPending  = iota
	statusStarted  = iota
	statusDone     = iota
	statusIncluded = iota
)

type renderTreeEl map[clientId]renderTreeEl

type renderedTree map[clientId]renderedTreeContents

type renderedTreeContents struct {
	render *specs.PageRender
	tree   renderedTree
}

type renderingQueue struct {
	queue      map[clientId]*queueElement
	args       *specs.RenderArgs
	topElement clientId
}

type queueElement struct {
	status              int
	render              *specs.PageRender
	queue               map[clientId]bool
	requiresSubElements bool
}

type queuedRenderResult struct {
	cID    clientId
	render *specs.PageRender
}

// func render(cID clientId, args *specs.RenderArgs) string {
// 	go render2(cID, args)
// 	tree := buildRenderTree(cID, args)
// 	render := renderTree(tree, args)
// 	return parseTree(cID, render)
// }

// func buildRenderTree(cID clientId, args *specs.RenderArgs) renderTreeEl {
// 	tree := renderTreeEl{}
// 	tree[cID] = renderTreeEl{}

// 	c, err := clients.GetById(cID)
// 	if err != nil {
// 		log.Printf("Failed to get client for %v (err = %v)", cID, err)
// 		return tree
// 	}

// 	d, err := c.Describe(context.Background(), args)
// 	if err != nil {
// 		log.Printf("Failed to fetch description of %v (err = %v)", cID, err)
// 		return tree
// 	}

// 	for _, inc := range d.Includes {
// 		incCID := clientId{inc.Name, "v1"}
// 		incTree := buildRenderTree(incCID, args)
// 		tree[cID][incCID] = incTree[incCID]
// 	}

// 	return tree
// }

// func renderTree(tree renderTreeEl, args *specs.RenderArgs) renderedTree {
// 	rtree := renderedTree{}

// 	for cID, subtree := range tree {
// 		r, err := renderElement(cID, args)
// 		if err != nil {
// 			log.Println("Failed to render element %v (err = %v)", cID, err)
// 			continue
// 		}

// 		rtree[cID] = renderedTreeContents{r, renderTree(subtree, args)}
// 	}

// 	return rtree
// }

// func parseTree(cID clientId, tree renderedTree) string {
// 	cHtml := tree[cID].render.Html

// 	for subCID, _ := range tree[cID].tree {
// 		re := regexp.MustCompile(`<element>` + subCID.name + `</element>`)
// 		cHtml = re.ReplaceAllString(cHtml, parseTree(subCID, tree[cID].tree))
// 	}

// 	return cHtml
// }

func renderElement(cID clientId, args *specs.RenderArgs) (*specs.PageRender, error) {
	c, err := clients.Get(cID.name, cID.version)
	if err != nil {
		log.Printf("Failed to get client for %s:%s (err = %v)", cID.name, cID.version, err)
		return nil, err
	}

	return c.Render(context.Background(), args)
}

func describeElement(cID clientId, args *specs.RenderArgs) (*specs.PageElementDescription, error) {
	c, err := clients.GetById(cID)
	if err != nil {
		log.Printf("Failed to get client for %v (err = %v)", cID, err)
		return nil, err
	}

	return c.Describe(context.Background(), args)
}

func render(cID clientId, args *specs.RenderArgs) string {
	q := newRenderingQueue()
	if err := q.buildQueue(cID, args); err != nil {
		log.Printf("Failed to build the render queue; err = %v", err)
		return ""
	}

	// log.Printf("Done with building the queue; %v", q)

	output, _ := q.render()
	// log.Printf("Rendered the page; err = %v; page = %v", err, output)

	return output.Html
}

func newRenderingQueue() *renderingQueue {
	return &renderingQueue{queue: map[clientId]*queueElement{}}
}

func (q *renderingQueue) buildQueue(cID clientId, args *specs.RenderArgs) error {
	q.args = args
	q.topElement = cID

	deltach := make(chan int)
	errch := make(chan error)
	defer close(errch)
	defer close(deltach)

	go q.addToQueue(cID, args, true, deltach, errch)

	todo := 1

	for {
		select {
		case i := <-deltach:
			todo += i
			// log.Printf("Updating with %d item(s) (now have %d)", i, todo)

			if todo == 0 {
				return nil
			}
		case e := <-errch:
			return e
		}
	}

	return nil
}

func (q *renderingQueue) addToQueue(cID clientId, args *specs.RenderArgs, inititalElement bool, deltach chan int, errch chan error) {
	// get the description
	d, err := describeElement(cID, args)
	if err != nil {
		log.Printf("Failed to get description for %v while adding element to queue (err = %v)", cID, err)
		deltach <- -1

		// send error back for initial element
		if inititalElement {
			errch <- err
		}

		return
	}
	// log.Printf("Received description for %v (%v)", cID, d)

	// start queueing up includes
	incQueue := map[clientId]bool{}
	for _, inc := range d.GetIncludes() {
		subCID := getActiveClientIdFor(inc.Name)
		go q.addToQueue(subCID, args, false, deltach, errch)
		incQueue[subCID] = false
	}

	q.queue[cID] = &queueElement{
		queue:               incQueue,
		requiresSubElements: inititalElement,
	}

	deltach <- (-1 + len(d.GetIncludes()))
}

func (q *renderingQueue) render() (*specs.PageRender, error) {
	ch := make(chan *queuedRenderResult)
	defer close(ch)

	q.startRendering(ch)

renderLoop:
	for {
		select {
		case r := <-ch:
			q.queue[r.cID].status = statusDone
			q.queue[r.cID].render = r.render

			// log.Printf("Received a new render for %v (%v)", r.cID, r.render)

			if q.allRendered() {
				// return nil, nil
				// log.Println("all done with rendering")
				break renderLoop
			}
			// log.Println("not everything ready yet ...")

			q.startRendering(ch)
		}
	}

	return q.glueTogether()
}

func (q *renderingQueue) startRendering(ch chan *queuedRenderResult) {
	for cID, el := range q.queue {
		if el.status == statusReady && (!el.requiresSubElements || el.subElementsDone(q)) {
			elCID := cID
			q.queue[elCID].status = statusPending
			go (func(elCID clientId) {
				r, err := renderElement(elCID, q.args)
				if err != nil {
					log.Printf("Failed to render %v; err = %v", elCID, err)
				}
				ch <- &queuedRenderResult{
					cID:    elCID,
					render: r,
				}
			})(elCID)
		}
	}
}

func (q *renderingQueue) allRendered() bool {
	for _, el := range q.queue {
		if el.status != statusDone {
			// log.Printf("%v not done yet (%d)", cID, el.status)
			return false
		}
	}

	return true
}

func (q *renderingQueue) glueTogether() (*specs.PageRender, error) {
	return q.getGluedElement(q.topElement)
}

func (q *renderingQueue) getGluedElement(cID clientId) (*specs.PageRender, error) {
	r := q.queue[cID].render
	for subCID, _ := range q.queue[cID].queue {
		replacement := ""
		if q.queue[subCID].status >= statusDone {
			replacementEl, _ := q.getGluedElement(subCID)
			replacement = replacementEl.Html
		}

		re := regexp.MustCompile(`<element>` + subCID.name + `</element>`)
		r.Html = re.ReplaceAllString(r.Html, replacement)
	}

	return r, nil
}

func (e *queueElement) subElementsDone(q *renderingQueue) bool {
	for cID, status := range e.queue {
		if status != true {
			if q.queue[cID].status == statusDone {
				e.queue[cID] = true
				continue
			}
			return false
		}
	}

	return true
}
