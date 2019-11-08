package main

import (
	"net/http"

	"github.com/lxc/lxd/lxd/response"
)

// APIEndpoint represents a URL in our API.
type APIEndpoint struct {
	Name   string // Name for this endpoint.
	Path   string // Path pattern for this endpoint.
	Get    APIEndpointAction
	Put    APIEndpointAction
	Post   APIEndpointAction
	Delete APIEndpointAction
	Patch  APIEndpointAction
}

// APIEndpointAlias represents an alias URL of and APIEndpoint in our API.
type APIEndpointAlias struct {
	Name string // Name for this alias.
	Path string // Path pattern for this alias.
}

// APIEndpointAction represents an action on an API endpoint.
type APIEndpointAction struct {
	Handler func(r *http.Request) response.Response
}