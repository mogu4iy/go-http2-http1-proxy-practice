package config

import (
	"gopkg.in/yaml.v2"
	"http2-http1.1-proxy/internal/constants"
	"os"
)

type ProxyHeaders struct {
	Upgrade          string `yaml:"Upgrade"`
	Connection       string `yaml:"Connection"`
	Host             string `yaml:"Host"`
	ProxyCacheBypass string `yaml:"proxy_cache_bypass"`
}

type Location struct {
	Path             string       `yaml:"path"`
	ProxyPass        string       `yaml:"proxy_pass"`
	ProxyHttpVersion string       `yaml:"proxy_http_version"`
	ProxyHeaders     ProxyHeaders `yaml:"proxy_set_header"`
}

type Server struct {
	Listen     uint8      `yaml:"listen"`
	ServerName string     `yaml:"server_name"`
	Locations  []Location `yaml:"locations"`
}

type UpstreamServer struct {
	Server      string `yaml:"server"`
	Weight      int    `yaml:"weight,omitempty"`
	MaxFails    int    `yaml:"max_fails,omitempty"`
	FailTimeout string `yaml:"fail_timeout,omitempty"`
	Backup      bool   `yaml:"backup,omitempty"`
}

type Upstream struct {
	Name      string           `yaml:"name"`
	Addresses []UpstreamServer `yaml:"addresses"`
}

type HTTP struct {
	Version   constants.HTTPVersion `yaml:"version"`
	Servers   []Server              `yaml:"servers"`
	Upstreams []Upstream            `yaml:"upstreams"`
}

type Config struct {
	HTTPs []HTTP `yaml:"https"`
}

func New(path string) (*Config, error) {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return &config, nil
}
