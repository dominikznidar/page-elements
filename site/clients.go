package main

import (
	"fmt"
	"page-elements/specs"
	"log"

	"github.com/bogdanovich/dns_resolver"
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

	// looks like we need to open a new connection
	// get the IP of the element
	ip, err := resolveClientAddr(cID)
	if err != nil {
		return nil, err
	}
	// connect to it
	conn, err := grpc.Dial(fmt.Sprintf("%s:80", ip), grpc.WithInsecure())
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

func resolveClientAddr(cID clientId) (string, error) {
	resolver := dns_resolver.New([]string{"consul"})
	ips, err := resolver.LookupHost(fmt.Sprintf("%s.%s.service.consul", cID.version, cID.name))
	if err != nil {
		log.Printf("Failed to lookup an address; err = %+v", err)
		return "", err
	}

	return ips[0].String(), nil
}
