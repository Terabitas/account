package storage

import (
	"github.com/nildev/account/Godeps/_workspace/src/github.com/juju/errors"
	"github.com/nildev/account/Godeps/_workspace/src/github.com/nildev/lib/registry"
	"github.com/nildev/account/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"
	"github.com/nildev/account/domain"
	"github.com/nildev/account/provision"
)

type (
	mongodb struct{}
)

func makeMongoDB() *mongodb {
	return &mongodb{}
}

// Save method
func (ds *mongodb) Save(acc *domain.Account) error {
	if acc == nil {
		return errors.Trace(errors.Errorf("acc is nil!"))
	}

	session, err := registry.CreateMongoDBClient()
	if err != nil {
		return err
	}

	collection := session.DB(registry.GetDatabaseName()).C(provision.TABLE_ACCOUNTS)
	_, err = collection.Upsert(bson.M{"email": acc.Email}, acc)

	if err != nil {
		return err
	}

	return nil
}

// GetByEmail method
func (ds *mongodb) GetByEmail(email string) (*domain.Account, error) {
	session, err := registry.CreateMongoDBClient()
	if err != nil {
		return nil, err
	}
	collection := session.DB(registry.GetDatabaseName()).C(provision.TABLE_ACCOUNTS)
	acc := &domain.Account{}
	collection.Find(bson.M{"email": email}).One(&acc)

	return acc, nil
}

// GetByUsername method
func (ds *mongodb) GetByUsername(username string) (*domain.Account, error) {
	session, err := registry.CreateMongoDBClient()
	if err != nil {
		return nil, err
	}
	collection := session.DB(registry.GetDatabaseName()).C(provision.TABLE_ACCOUNTS)
	acc := &domain.Account{}
	collection.Find(bson.M{"userName": username}).One(&acc)

	return acc, nil
}

// GetById method
func (ds *mongodb) GetById(id string) (*domain.Account, error) {
	session, err := registry.CreateMongoDBClient()
	if err != nil {
		return nil, err
	}
	collection := session.DB(registry.GetDatabaseName()).C(provision.TABLE_ACCOUNTS)
	acc := &domain.Account{}
	collection.FindId(bson.M{"_id": id}).One(&acc)

	return acc, nil
}
