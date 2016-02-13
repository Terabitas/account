package provision

import (
	"fmt"

	"github.com/nildev/account/Godeps/_workspace/src/github.com/nildev/lib/registry"
	"github.com/nildev/account/Godeps/_workspace/src/gopkg.in/mgo.v2"
)

var (
	TABLE_ACCOUNTS = "accounts"
)

// NildevInitMongoDB init
func NildevInitMongoDB() {
	// In which environment we are
	//env := registry.GetEnv()
	//fmt.Printf("%s", env)

	session, err := registry.CreateMongoDBClient()
	if err != nil {
		fmt.Printf("%s", err)
	}

	err = session.DB(registry.GetDatabaseName()).C(TABLE_ACCOUNTS).Create(&mgo.CollectionInfo{})

	if err != nil {
		fmt.Printf("%s", err)
	}

	usernameIndex := mgo.Index{
		Key:        []string{"userName"},
		Unique:     true,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	}

	err = session.DB(registry.GetDatabaseName()).C(TABLE_ACCOUNTS).EnsureIndex(usernameIndex)

	if err != nil {
		fmt.Printf("%s", err)
	}

	emailIndex := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	}

	err = session.DB(registry.GetDatabaseName()).C(TABLE_ACCOUNTS).EnsureIndex(emailIndex)

	if err != nil {
		fmt.Printf("%s", err)
	}
}

func DestroyMongoDB() {
	session, err := registry.CreateMongoDBClient()
	if err != nil {
		fmt.Printf("%s", err)
	}

	err = session.DB(registry.GetDatabaseName()).C(TABLE_ACCOUNTS).DropCollection()

	if err != nil {
		fmt.Printf("%s", err)
	}
}
