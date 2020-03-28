package implementation

import (
	"github.com/go-kit/kit/log"
	"gitlab.com/hyperd/konga-backend"
)

// service implements the konga Service
type service struct {
	repository titanic.Repository
	logger     log.Logger
}

// NewService creates and returns a new konga service instance
func NewService(rep konga.Repository, logger log.Logger) konga.Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}
