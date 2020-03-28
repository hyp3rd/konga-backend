package inmemory

import (
	"errors"
	"sync"

	"github.com/go-kit/kit/log"
	"gitlab.com/hyperd/konga-backend"
)

// Response errors
var (
	ErrInconsistentID = errors.New("inconsistent ID")
	ErrAlreadyExists  = errors.New("already exists")
	ErrNotFound       = errors.New("not found")
)

type repository struct {
	mtx    sync.RWMutex
	logger log.Logger
}

// NewInmemService returns an in-memory storage
func NewInmemService(logger log.Logger) (konga.Repository, error) {
	return &repository{
		logger: log.With(logger, "repository", "inmemory"),
	}, nil
}
