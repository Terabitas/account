package account

import (
	"os"
	"testing"

	"github.com/nildev/account/Godeps/_workspace/src/github.com/jarcoal/httpmock"
	"github.com/nildev/account/Godeps/_workspace/src/golang.org/x/oauth2/github"
	. "github.com/nildev/account/Godeps/_workspace/src/gopkg.in/check.v1"
	"github.com/nildev/account/domain"
	"github.com/nildev/account/oauth"
	"github.com/nildev/account/provision"
	"github.com/nildev/account/storage"
)

type AccountIntegrationSuite struct{}

var _ = Suite(&AccountIntegrationSuite{})

func TestMain(m *testing.M) {
	dbname := "nildev"
	ip := os.Getenv("ND_IP_PORT")

	os.Setenv("ND_ENV", "dev")
	os.Setenv("ND_DATABASE_NAME", dbname)
	os.Setenv("ND_STORAGE", storage.STORAGE_MONGODB)
	os.Setenv("ND_MONGODB_URL", "mongodb://"+ip+"/"+dbname)

	provision.NildevInitMongoDB()
	code := m.Run()
	provision.DestroyMongoDB()
	os.Exit(code)
}

func (s *AccountIntegrationSuite) TestIfAccountWhichDoesNotExistsIsBeingRegistered(c *C) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", github.Endpoint.TokenURL,
		httpmock.NewStringResponder(200, `{"access_token":"e72e16c7e42f292c6912e7710c838347ae178b4a", "scope":"repo,gist", "token_type":"bearer"}`))

	httpmock.RegisterResponder("GET", "https://api.github.com/user",
		httpmock.NewStringResponder(200, `{
  "login": "octocat",
  "id": 1,
  "avatar_url": "https://github.com/images/error/octocat_happy.gif",
  "gravatar_id": "",
  "url": "https://api.github.com/users/octocat",
  "html_url": "https://github.com/octocat",
  "followers_url": "https://api.github.com/users/octocat/followers",
  "following_url": "https://api.github.com/users/octocat/following{/other_user}",
  "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
  "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
  "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
  "organizations_url": "https://api.github.com/users/octocat/orgs",
  "repos_url": "https://api.github.com/users/octocat/repos",
  "events_url": "https://api.github.com/users/octocat/events{/privacy}",
  "received_events_url": "https://api.github.com/users/octocat/received_events",
  "type": "User",
  "site_admin": false,
  "name": "monalisa octocat",
  "company": "GitHub",
  "blog": "https://github.com/blog",
  "location": "San Francisco",
  "email": "octocat@github.com",
  "hireable": false,
  "bio": "There once was...",
  "public_repos": 2,
  "public_gists": 1,
  "followers": 20,
  "following": 0,
  "created_at": "2008-01-14T04:33:35Z",
  "updated_at": "2008-01-14T04:33:35Z",
  "total_private_repos": 100,
  "owned_private_repos": 100,
  "private_gists": 81,
  "disk_usage": 10000,
  "collaborators": 8,
  "plan": {
    "name": "Medium",
    "space": 400,
    "private_repos": 20,
    "collaborators": 0
  }
}`))

	rez, err := Register(oauth.PROVIDER_GITHUB, "xxxx", domain.Data{})

	c.Assert(rez, Equals, true)
	c.Assert(err, IsNil)
}

//
//func (s *AccountIntegrationSuite) TestIfCorrectErrorIsReturnedWhenBadProviderIsPassed(c *C) {
//	c.Assert(true, Equals, false)
//}
//
//func (s *AccountIntegrationSuite) TestIfCorrectErrorIsReturnedWhenBadEmailIsPassed(c *C) {
//	c.Assert(true, Equals, false)
//}
//
//func (s *AccountIntegrationSuite) TestIfEmailWhichAlreadyExistsIsNotCreatedAgain(c *C) {
//	c.Assert(true, Equals, false)
//}
//
//func (s *AccountIntegrationSuite) TestIfUsernameWhichAlreadyExistsIsNotCreatedAgain(c *C) {
//	c.Assert(true, Equals, false)
//}
//
//func (s *AccountIntegrationSuite) TestIfAuthTokenWhichAlreadyExistsIsNotCreatedAgain(c *C) {
//	c.Assert(true, Equals, false)
//}
//
//func (s *AccountIntegrationSuite) TestIfTryingCreateAccountWhichAlreadyExistsReturnsCorrectError(c *C) {
//	c.Assert(true, Equals, false)
//}
//
//func (s *AccountIntegrationSuite) TestIfUserIsAbleToLoginByEmail(c *C) {
//	c.Assert(true, Equals, false)
//}
//
//func (s *AccountIntegrationSuite) TestIfUserIsAbleToLoginByUsername(c *C) {
//	c.Assert(true, Equals, false)
//}
//
//func (s *AccountIntegrationSuite) TestIfCredentialsAreWrongWhenTryingToAuthByEmailCorrectErrorIsReturned(c *C) {
//	c.Assert(true, Equals, false)
//}
//
//func (s *AccountIntegrationSuite) TestIfCredentialsAreWrongWhenTryingToAuthByUsernameCorrectErrorIsReturned(c *C) {
//	c.Assert(true, Equals, false)
//}
//
//func (s *AccountIntegrationSuite) TestIfTokenIsWrongWhenTryingToAuthByUsernameCorrectErrorIsReturned(c *C) {
//	c.Assert(true, Equals, false)
//}
//
//func (s *AccountIntegrationSuite) TestIfTokenIsWrongWhenTryingToAuthByEmailCorrectErrorIsReturned(c *C) {
//	c.Assert(true, Equals, false)
//}
//
//func (s *AccountIntegrationSuite) TestIfCreatedAccountIsActive(c *C) {
//	c.Assert(true, Equals, false)
//}
