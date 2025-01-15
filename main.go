package certdash

import (
	"context"
	"fmt"
	"net/http"
)

type Config struct {
}

// CreateConfig creates the DEFAULT plugin configuration - no access to config yet!
func CreateConfig() *Config {
	return &Config{}
}

type CertDashPlugin struct {
	Config *Config
	name   string
	next   http.Handler
}

// New created a new plugin, with a config that's been set (possibly) by the admin
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config == nil {
		return nil, fmt.Errorf("config can not be nil")
	}

	plugin := &CertDashPlugin{
		Config: config,
		next:   next,
		name:   name,
	}

	return plugin, nil
}

func (t *CertDashPlugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// FIXME

	t.next.ServeHTTP(rw, req)
}
