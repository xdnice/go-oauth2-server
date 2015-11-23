package session

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RichardKnop/go-oauth2-server/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// SessionTestSuite needs to be exported so the tests run
type SessionTestSuite struct {
	suite.Suite
	service *Service
}

// The SetupSuite method will be run by testify once, at the very
// start of the testing suite, before any tests are run.
func (suite *SessionTestSuite) SetupSuite() {
	// Overwrite internal vars so we don't affect existing session
	storageSessionName = "test_session"
	userSessionKey = "test_user"

	// Initialise the service
	r, err := http.NewRequest("GET", "http://1.2.3.4/foo/bar", nil)
	if err != nil {
		log.Fatal(err)
	}
	w := httptest.NewRecorder()
	suite.service = NewService(config.NewConfig(), r, w)
}

// The TearDownSuite method will be run by testify once, at the very
// end of the testing suite, after all tests have been run.
func (suite *SessionTestSuite) TearDownSuite() {
	suite.service.ClearUserSession()
}

// The SetupTest method will be run before every test in the suite.
func (suite *SessionTestSuite) SetupTest() {
	//
}

// The TearDownTest method will be run after every test in the suite.
func (suite *SessionTestSuite) TearDownTest() {
	//
}

func (suite *SessionTestSuite) TestService() {
	// No public methods should work before StartSession has been called
	userSession, err := suite.service.GetUserSession()
	if assert.NotNil(suite.T(), err) {
		assert.Equal(suite.T(), errSessonNotStarted.Error(), err.Error())
	}

	// Call the StartSession method so internal session object gets set
	if err := suite.service.StartSession(); err != nil {
		log.Fatal(err)
	}

	// Since StartSession has not been called yet, this should return an error
	userSession, err = suite.service.GetUserSession()
	if assert.NotNil(suite.T(), err) {
		assert.Equal(suite.T(), "User session type assertion error", err.Error())
	}

	log.Print(userSession)
}

// TesSessionTestSuite ...
// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TesSessionTestSuite(t *testing.T) {
	suite.Run(t, new(SessionTestSuite))
}
