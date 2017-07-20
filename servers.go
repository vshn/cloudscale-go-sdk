package cloudscale

import (
	"context"
	"fmt"
	"net/http"
)

const serverBasePath = "v1/servers"

type Server struct {
	HREF            string      `json:"href"`
	UUID            string      `json:"uuid"`
	Name            string      `json:"name"`
	Status          string      `json:"status"`
	Flavor          Flavor      `json:"flavor"`
	Image           Image       `json:"image"`
	Volumes         []Volume    `json:"volumes"`
	Interfaces      []Interface `json:"interfaces"`
	SSHFingerprints []string    `json:"ssh_fingerprints"`
	SSHHostKeys     []string    `json:"ssh_host_keys"`
	AntiAfinityWith []string    `json:"anti-affinity-with"`
}

type Flavor struct {
	Slug      string `json:"slug"`
	Name      string `json:"name"`
	VCPUCount int    `json:"vcpu_count"`
	MemoryGB  int    `json:"memory_gb"`
}

type Image struct {
	Slug            string `json:"slug"`
	Name            string `json:"name"`
	OperatingSystem string `json:"operating_system"`
}

type Volume struct {
	Type       string `json:"ssd"`
	DevicePath string `json:"device_path"`
	SizeGB     int    `json:"SizeGB"`
}

type Interface struct {
	Type     string    `json:"type"`
	Adresses []Address `json:"addresses"`
}

type Address struct {
	Version      int    `json:"version"`
	Address      string `json:"address"`
	PrefixLenght string `json:"prefix_lenght"`
	Gateway      string `json:"gateway"`
	ReversePtr   string `json:"reverse_prt"`
}

type ServerRequest struct {
	Name         string   `json:"name"`
	Flavor       string   `json:"flavor"`
	Image        string   `json:"image"`
	VolumeSizeGB int      `json:"volume_size_gb"`
	SSHKeys      []string `json:"ssh_keys"`
}

type ServerService interface {
	Create(ctx context.Context, createRequest *ServerRequest) (*Server, error)
	Get(ctx context.Context, serverID string) (*Server, error)
	Update(ctx context.Context, serverID string) error
	Delete(ctx context.Context, serverID string) error
	List(ctx context.Context) ([]Server, error)
	Reboot(ctx context.Context, serverID string) error
	Start(ctx context.Context, serverID string) error
	Stop(ctx context.Context, serverID string) error
}

type ServerServiceOperations struct {
	client *Client
}

func (s ServerServiceOperations) Create(ctx context.Context, createRequest *ServerRequest) (*Server, error) {
	path := serverBasePath

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, createRequest)
	if err != nil {
		return nil, err
	}

	server := new(Server)

	err = s.client.Do(ctx, req, server)
	if err != nil {
		return nil, err
	}

	return server, nil
}

func (s ServerServiceOperations) Update(ctx context.Context, serverID string) error {
	return nil
}

func (s ServerServiceOperations) Get(ctx context.Context, serverID string) (*Server, error) {
	path := fmt.Sprintf("%s/%s", serverBasePath, serverID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	server := new(Server)
	err = s.client.Do(ctx, req, server)
	if err != nil {
		return nil, err
	}

	return server, nil
}
func (s ServerServiceOperations) Delete(ctx context.Context, serverID string) error {
	path := fmt.Sprintf("%s/%s", serverBasePath, serverID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	return s.client.Do(ctx, req, nil)

}
func (s ServerServiceOperations) Reboot(ctx context.Context, serverID string) error {
	path := fmt.Sprintf("%s/%s/reboot", serverBasePath, serverID)
	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return err
	}
	return s.client.Do(ctx, req, nil)
}

func (s ServerServiceOperations) Start(ctx context.Context, serverID string) error {
	path := fmt.Sprintf("%s/%s/start", serverBasePath, serverID)
	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return err
	}
	return s.client.Do(ctx, req, nil)
}

func (s ServerServiceOperations) Stop(ctx context.Context, serverID string) error {
	path := fmt.Sprintf("%s/%s/stop", serverBasePath, serverID)
	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return err
	}
	return s.client.Do(ctx, req, nil)
}

func (s ServerServiceOperations) List(ctx context.Context) ([]Server, error) {
	path := serverBasePath

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	servers := []Server{}
	err = s.client.Do(ctx, req, &servers)
	if err != nil {
		return nil, err
	}

	return servers, nil
}
