package cluster

import (
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/network"
)

// Provider provides clustering config details
type Provider interface {
	IsManager() bool
	IsAgent() bool
	GetListenAddress() string
	GetRemoteAddress() string
	ListenClusterEvents() <-chan struct{}
	AllocateEndpoint(string, string, *network.EndpointSettings) (string, types.NetworkCreateRequest, error)
	DeallocateEndpoint(id string) error
}
