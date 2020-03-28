package cockroachdb

import (
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	"gitlab.com/hyperd/konga-backend"
)

type repository struct {
	db     *gorm.DB
	logger log.Logger
}

// New returns a concrete repository backed by CockroachDB
func New(db *gorm.DB, logger log.Logger) (konga.Repository, error) {
	// return  repository
	return &repository{
		db:     db,
		logger: log.With(logger, "rep", "cockroachdb"),
	}, nil
}

// Functions of type `txnFunc` are passed as arguments to our
// `runTransaction` wrapper that handles transaction retries for us
// (see implementation below).
type txnFunc func(*gorm.DB) error

// This function is used for testing the transaction retry loop.  It
// can be deleted from production code.
var forceRetryLoop txnFunc = func(db *gorm.DB) error {

	// The first statement in a transaction can be retried transparently
	// on the server, so we need to add a dummy statement so that our
	// force_retry statement isn't the first one.
	if err := db.Exec("SELECT now()").Error; err != nil {
		return err
	}
	// Used to force a transaction retry.  Can only be run as the
	// 'root' user.
	if err := db.Exec("SELECT crdb_internal.force_retry('1s'::INTERVAL)").Error; err != nil {
		return err
	}
	return nil
}
