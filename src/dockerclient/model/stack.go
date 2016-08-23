package model

import (
	"github.com/docker/engine-api/types/swarm"
)

// bundle stores the contents of services and stack name
type Bundle struct {
	Stack     BundleService `json:"Stack"`
	Namespace string        `json:"Namespace"`
}

// BundleService content services spec map and stack version
// Correspondence docker daemon type BundleFile
type BundleService struct {
	Version  string                      `json:"Version"`
	Services map[string]RolexServiceSpec `json:"Services"`
}

type RolexServiceSpec struct {
	Name         string              `json:"Name"`
	Labels       map[string]string   `json:"Labels"`
	TaskTemplate swarm.TaskSpec      `json:"TaskTemplate"`
	Mode         swarm.ServiceMode   `json:"Mode"`
	UpdateConfig *swarm.UpdateConfig `json:"UpdateConfig"`
	Networks     []string            `json:"Networks"`
	EndpointSpec *swarm.EndpointSpec `json:"EndpointSpec"`
}

type RolexService struct {
	ID string `json:"ID"`
	swarm.Meta
	Spec         RolexServiceSpec   `josn:"Spec"`
	Endpoint     swarm.Endpoint     `json:"Endpoint"`
	UpdateStatus swarm.UpdateStatus `json:"UpdateStatus"`
}
