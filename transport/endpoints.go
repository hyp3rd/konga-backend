package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"gitlab.com/hyperd/konga-backend"
)

// Endpoints collects all of the konga API endpoints.
type Endpoints struct {
	GetAPIStatusEndpoint endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided konga.Service. Useful in a konga.Service
// server.
func MakeServerEndpoints(s konga.Service) Endpoints {
	return Endpoints{
		GetAPIStatusEndpoint: MakeGetAPIStatusEndpoint(),
	}
}

// MakeGetAPIStatusEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakeGetAPIStatusEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		return GetAPIStatusResponse{Status: "Healthy"}, nil
	}
}

// GetAPIStatusRequest request object
type GetAPIStatusRequest struct{}

// GetAPIStatusResponse response object
type GetAPIStatusResponse struct {
	Status string `json:"status,omitempty"`
	Err    error  `json:"err,omitempty"`
}

func (r GetAPIStatusResponse) error() error { return r.Err }
