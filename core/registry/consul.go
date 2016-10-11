package registry

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	consul "github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
)

type Registry struct {
	client         *consul.Client
	serviceID      string
	serviceName    string
	serviceVersion string
}

var (
	consulAddr, traefikDomain string
)

func init() {
	consulAddr = envString("CONSUL", "localhost:8500")
	traefikDomain = envString("DOMAIN", "")
}

func NewRegistry(name, version string) (*Registry, error) {
	config := consul.DefaultConfig()
	config.Address = consulAddr
	c, err := consul.NewClient(config)
	return &Registry{
		client:         c,
		serviceName:    name,
		serviceVersion: version,
	}, err
}

func (r *Registry) Register() error {
	r.serviceID = uuid.NewV4().String()

	// hostname, _ := os.Hostname()
	ip := GetLocalIP()
	log.Printf("IP for service %s set to %s", r.serviceName, ip)

	// compose tags
	tags := []string{r.serviceVersion}
	if traefikDomain == "" {
		tags = append(tags, "traefik.enable=false")
	} else {
		tags = append(tags, "traefik.frontend.rule=Host:"+traefikDomain)
	}

	// register the service
	service := &consul.AgentServiceRegistration{
		ID:      r.serviceID,
		Name:    r.serviceName,
		Tags:    tags,
		Port:    80,
		Address: ip,
		Check: &consul.AgentServiceCheck{
			TTL: "20s",
			DeregisterCriticalServiceAfter: "30s",
			Status: "passing",
		},
	}

	err := r.client.Agent().ServiceRegister(service)
	log.Printf("Registered service '%s' in consul (err = %s);", r.serviceName, err)

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
