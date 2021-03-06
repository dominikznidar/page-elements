package main

import (
	"context"
	"page-elements/specs"
	"log"
	"regexp"
	"sync"
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
	queuemu    sync.Mutex
	args       *specs.RenderArgs
	argsmu     sync.Mutex
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

func render(cID clientId, args *specs.RenderArgs) (string, error) {
	q := newRenderingQueue()
	if err := q.buildQueue(cID, args); err != nil {
		log.Printf("Failed to build the render queue; err = %v", err)
		return "", err
	}

	output, _ := q.render()
	return output.Html, nil
}

func newRenderingQueue() *renderingQueue {
	return &renderingQueue{queue: map[clientId]*queueElement{}}
}

func (q *renderingQueue) buildQueue(cID clientId, args *specs.RenderArgs) error {
	q.args = args
	q.topElement = cID

	deltach := make(chan int)
	defer close(deltach)
	errch := make(chan error)
	defer close(errch)

	go q.addToQueue(cID, args, true, deltach, errch)

	todo := 1

	for {
		select {
		case i := <-deltach:
			todo += i
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

		// send error back for initial element
		if inititalElement {
			errch <- err
			return
		}

		// decrease todo list
		deltach <- -1
		return
	}

	// start queueing up includes
	incQueue := map[clientId]bool{}
	for _, inc := range d.GetIncludes() {
		subCID := getActiveClientIdFor(inc.Name)
		go q.addToQueue(subCID, args, false, deltach, errch)
		incQueue[subCID] = false
	}

	q.queuemu.Lock()
	q.queue[cID] = &queueElement{
		queue:               incQueue,
		requiresSubElements: inititalElement,
	}
	q.queuemu.Unlock()

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

			q.queuemu.Lock()
			q.queue[r.cID].status = statusDone
			q.queue[r.cID].render = r.render
			q.queuemu.Unlock()
			q.updateArgs(r.render)

			if q.allRendered() {
				break renderLoop
			}

			q.startRendering(ch)
		}
	}

	return q.glueTogether()
}

func (q *renderingQueue) startRendering(ch chan *queuedRenderResult) {
	for cID, el := range q.queue {
		if el.status == statusReady && (!el.requiresSubElements || el.subElementsDone(q)) {
			elCID := cID
			q.queuemu.Lock()
			q.queue[elCID].status = statusPending
			q.queuemu.Unlock()
			go (func(elCID clientId) {
				r, err := renderElement(elCID, q.getArgs())
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
		if qEl, ok := q.queue[subCID]; ok && qEl.status >= statusDone {
			replacementEl, _ := q.getGluedElement(subCID)
			replacement = replacementEl.Html
		}

		re := regexp.MustCompile(`<element>` + subCID.name + `</element>`)
		r.Html = re.ReplaceAllString(r.Html, replacement)
	}

	return r, nil
}

func (q *renderingQueue) getArgs() *specs.RenderArgs {
	q.argsmu.Lock()
	defer q.argsmu.Unlock()
	return q.args
}

func (q *renderingQueue) updateArgs(r *specs.PageRender) {
	q.argsmu.Lock()
	defer q.argsmu.Unlock()

	if t := r.PageTitle; t != "" {
		q.args.Add("pageTitle", t)
	}
}

func (e *queueElement) subElementsDone(q *renderingQueue) bool {
	for cID, status := range e.queue {
		if status != true {
			if qEl, ok := q.queue[cID]; ok && qEl.status >= statusDone {
				e.queue[cID] = true
				continue
			}
			return false
		}
	}

	return true
}
