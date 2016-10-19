package main

import (
	"go-micro-site/core/registry"
	"log"
	"time"
)

type state struct {
	cIndex uint64
	state  *registry.State
}

var cstate = &state{}

func initState(reg *registry.Registry) error {
	// start polling current state from consul
	go startPollingState(reg)

	return nil
}

func startPollingState(reg *registry.Registry) {
	for {
		log.Printf("Fetching new state from consul; cIndex = %d", cstate.cIndex)
		state, cIndex, err := reg.WaitForNewState(cstate.cIndex, 30*time.Second)
		log.Printf("Received new state from consul; state = %+v; cIndex = %d; err = %v", state, cIndex, err)

		cstate.cIndex = cIndex
		cstate.state = state
	}
}

func getVersionFor(element string) string {
	version, ok := (*cstate.state)[element]
	if !ok || version == "" {
		return "off"
	}

	return version
}

func getActiveClientIdFor(element string) clientId {
	return clientId{element, getVersionFor(element)}
}
