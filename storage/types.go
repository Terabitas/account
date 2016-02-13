package storage

import (
	"os"

	"github.com/nildev/account/Godeps/_workspace/src/github.com/juju/errors"
	"github.com/nildev/account/domain"
)

const (
	STORAGE_MONGODB = "mongodb"
)

type (
	// AccountRepository type
	AccountRepository interface {
		Save(*domain.Account) error
		GetByEmail(email string) (*domain.Account, error)
		GetByUsername(username string) (*domain.Account, error)
		GetById(id string) (*domain.Account, error)
	}
)

// MakeRepository factory
func MakeRepositoryFromEnv() (AccountRepository, error) {
	storage := os.Getenv("ND_STORAGE")
	switch storage {
	case STORAGE_MONGODB:
		return makeMongoDB(), nil
	}

	return nil, errors.Trace(errors.Errorf("Storage [%s] is not supported", storage))
}
