package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"sort"
	"time"

	consul "github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
)

type Registry struct {
	client         *consul.Client
	serviceID      string
	serviceName    string
	serviceVersion string
	systemTool     bool
}

type Clients map[string][]string
type State map[string]string

var (
	consulAddr, traefikDomain string
)

func init() {
	consulAddr = envString("CONSUL", "localhost:8500")
	traefikDomain = envString("DOMAIN", "")
}

func NewRegistry(name, version string, systemTool bool) (*Registry, error) {
	config := consul.DefaultConfig()
	config.Address = consulAddr
	c, err := consul.NewClient(config)
	return &Registry{
		client:         c,
		serviceName:    name,
		serviceVersion: version,
		systemTool:     systemTool,
	}, err
}

func (r *Registry) Register() error {
	r.serviceID = uuid.NewV4().String()

	// compose tags
	tags := []string{r.serviceVersion}
	if traefikDomain == "" {
		tags = append(tags, "traefik.enable=false")
	} else {
		tags = append(tags, "traefik.frontend.rule=Host:"+traefikDomain)
	}
	if !r.systemTool {
		tags = append(tags, "micro-element")
	}

	// register the service
	service := &consul.AgentServiceRegistration{
		ID:      r.serviceID,
		Name:    r.serviceName,
		Tags:    tags,
		Port:    80,
		Address: GetLocalIP(),
		Check: &consul.AgentServiceCheck{
			TTL: "20s",
			DeregisterCriticalServiceAfter: "30s",
			Status: "passing",
		},
	}

	err := r.client.Agent().ServiceRegister(service)
	log.Printf("Registered service '%s' in consul (err = %v);", r.serviceName, err)

	t := time.NewTicker(10 * time.Second)
	go func() {
		for ts := range t.C {
			_ = r.client.Agent().PassTTL("service:"+r.serviceID, fmt.Sprintf("I'm alive at %s", ts))
		}
	}()

	return nil
}

func (r *Registry) Unregister() error {
	err := r.client.Agent().ServiceDeregister(r.serviceID)
	log.Printf("Unregistered the service '%s'", r.serviceName)
	return err
}

func (r *Registry) FetchAvailableClients() (Clients, error) {
	// fetch all services from catalog
	services, _, _ := r.client.Catalog().Services(nil)
	log.Printf("Fetched following services: %v", services)

	clients := Clients{}
	for serviceName, tags := range services {
		if !containsMicroElementTag(tags) {
			continue
		}

		versions := getVersionTags(tags)

		log.Printf("Found valid element %s:%v", serviceName, versions)

		clients[serviceName] = versions
	}

	return clients, nil
}

func (r *Registry) FetchCurrentState() (*State, error) {
	// read state from consul
	res, qm, err := r.client.KV().Get("micro/state", nil)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return &State{}, nil
	}

	log.Printf("Read state from consul: %v; err = %v; qm = %v", res, err, qm)
	buf := bytes.NewReader(res.Value)

	// decode JSON to State
	state := &State{}
	decoder := json.NewDecoder(buf)
	err = decoder.Decode(state)

	return state, err
}

func (r *Registry) UpdateState(newState State) error {
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	enc.Encode(newState)

	_, err := r.client.KV().Put(&consul.KVPair{Key: "micro/state", Value: buf.Bytes()}, nil)
	return err
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func envString(envVar, defaultValue string) string {
	if v := os.Getenv(envVar); v != "" {
		return v
	}
	return defaultValue
}

func containsMicroElementTag(tags []string) bool {
	for _, tag := range tags {
		if tag == "micro-element" {
			return true
		}
	}

	return false
}

var versionRe = regexp.MustCompile(`^v\d+$`)

func getVersionTags(tags []string) []string {
	versions := []string{}
	for _, tag := range tags {
		if versionRe.MatchString(tag) {
			versions = append(versions, tag)
		}
	}

	sort.Sort(sort.Reverse(sort.StringSlice(versions)))
	return versions
}
