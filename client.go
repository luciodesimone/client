package client

import (
	"net/http"

	"github.com/luciods1/logger"
)

type ServiceInfo struct {
	Name     string
	Endpoint string
}

type HTTPAction struct {
	Method string
	Path   string
}

type Config struct {
	BaseURL string
	// TODO: review if this should change to be an object
	Credentials string
	Logger      logger.Logger
	HTTPClient  *http.Client
	UserAgent   string
}
