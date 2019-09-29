package config

import "fmt"

type AppConfig struct {
	GrpcHosts *GrpcHosts
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		GrpcHosts: &GrpcHosts{
			Account: &GrpcHost{
				Host: "127.0.0.1",
				Port: 19004,
			},
			Content: &GrpcHost{
				Host: "127.0.0.1",
				Port: 19003,
			},
		},
	}
}

type GrpcHosts struct {
	Account *GrpcHost
	Content *GrpcHost
}

type GrpcHost struct {
	Host string
	Port uint32
}

func (h *GrpcHost) Build() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}