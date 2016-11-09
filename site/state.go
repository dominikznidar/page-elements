package main

import (
	"page-elements/core/registry"
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
		state, cIndex, _ := reg.WaitForNewState(cstate.cIndex, 30*time.Second)
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
