package gelf

import (
	"fmt"
	"io"

	"github.com/luraproject/lura/config"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

// Namespace is the key to look for extra configuration details
const Namespace = "github_com/devopsfaith/krakend-gelf"

var (
	newTCPWriter = gelf.NewTCPWriter
	newUDPWriter = gelf.NewUDPWriter
	// ErrWrongConfig is the error returned when there is no config under the namespace
	ErrWrongConfig = fmt.Errorf("getting the extra config for the krakend-gelf module")
	// ErrMissingAddr is the error returned when there address param is empty
	ErrMissingAddr = fmt.Errorf("missing addr to send gelf logs")
)

// NewWriter returns an io.Writer to write gelf logs to a server
func NewWriter(cfg config.ExtraConfig) (io.Writer, error) {
	logconfig, ok := ConfigGetter(cfg).(Config)
	if !ok {
		return nil, ErrWrongConfig
	}
	if logconfig.Addr == "" {
		return nil, ErrMissingAddr
	}

	if logconfig.EnableTCP {
		return newTCPWriter(logconfig.Addr)
	}
	return newUDPWriter(logconfig.Addr)
}

// ConfigGetter implements the config.ConfigGetter interface
func ConfigGetter(e config.ExtraConfig) interface{} {
	v, ok := e[Namespace]
	if !ok {
		return nil
	}
	tmp, ok := v.(map[string]interface{})
	if !ok {
		return nil
	}
	cfg := Config{}
	if v, ok := tmp["address"]; ok {
		cfg.Addr = v.(string)
	}
	if v, ok := tmp["enable_tcp"]; ok {
		cfg.EnableTCP = v.(bool)
	}
	return cfg
}

// Config is the custom config struct containing the params for the Writer
type Config struct {
	Addr      string
	EnableTCP bool
}
