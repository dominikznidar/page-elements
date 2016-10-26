package main

import (
	"fmt"
	"go-micro-site/specs"

	"google.golang.org/grpc"
)

type clientId struct {
	name, version string
}

type clientConnection struct {
	client specs.PageElementClient
	conn   *grpc.ClientConn
}

type clientConnections struct {
	storage map[clientId]*clientConnection
}

func newClientConnections() *clientConnections {
	return &clientConnections{
		storage: map[clientId]*clientConnection{},
	}
}

func (c *clientConnections) GetById(cID clientId) (specs.PageElementClient, error) {
	if cID.version == "off" {
		return nil, fmt.Errorf("Can't connect to client (%s) that is marked as off", cID.name)
	}
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

	c.storage[cID] = &clientConnection{
		client: client,
		conn:   conn,
	}

	return client, nil
}

func (c *clientConnections) Get(name, version string) (specs.PageElementClient, error) {
	return c.GetById(clientId{name, version})
}

func (c *clientConnections) Close() {
	for _, client := range c.storage {
		client.conn.Close()
	}
}
