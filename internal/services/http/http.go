package http

import (
	"net"
	"net/http"
	"time"
)

//Config for http package
type Config struct {
	TimeoutInSeconds     int `required:"true" split_words:"true"`
	DialTimeoutInSeconds int `required:"true" split_words:"true"`
	MaxConnsPerHost      int `required:"true" split_words:"true"`
}

//Service ...
type Service struct {
	HTTPClient http.Client
	Config     Config
}

//Initialize the http client to be shared for all services
func Initialize(config Config) Service {
	return Service{
		Config: config,
		HTTPClient: http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   time.Duration(config.DialTimeoutInSeconds) * time.Second,
					KeepAlive: 10 * time.Second,
				}).DialContext,
				TLSHandshakeTimeout: 10 * time.Second,

				ExpectContinueTimeout: 10 * time.Second,
				ResponseHeaderTimeout: 60 * time.Second,
				MaxConnsPerHost:       config.MaxConnsPerHost,
			},
			Timeout: time.Duration(config.TimeoutInSeconds) * time.Second,
		},
	}
}
