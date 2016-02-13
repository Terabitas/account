package account // import "github.com/nildev/account"

import (
	"github.com/nildev/account/domain"
	"github.com/nildev/account/oauth"
	"github.com/nildev/account/storage"
)

// Register based on requested provider exchanges given authCode for token
// and retrieves data.
// Account is persisted along with given arbitrary `data`
func Register(provider string, authCode string, data domain.Data) (result bool, err error) {
	reader, err := oauth.MakeReader(provider)
	if err != nil {
		return false, err
	}

	token, err := reader.Token(authCode)
	if err != nil {
		return false, err
	}

	writer, err := storage.MakeRepositoryFromEnv()
	if err != nil {
		return false, err
	}

	email, err := reader.Email()
	if err != nil {
		return false, err
	}

	uname, err := reader.Username()
	if err != nil {
		return false, err
	}

	avatar, err := reader.Avatar()
	if err != nil {
		return false, err
	}

	acc := domain.MakeAccount(
		*email,
		*uname,
		*avatar,
		token,
		data,
	)

	err = writer.Save(acc)

	if err != nil {
		return false, err
	}

	return true, nil
}
