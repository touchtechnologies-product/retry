package retry

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
	inMemType   string
	key         string
	backOffTime int
	maxRetry    int
}

func (suite *PackageTestSuite) SetupTest() {
	suite.inMemType = "*retry.InMemManager"
	suite.key = "test-key"
	suite.backOffTime = 2
	suite.maxRetry = 3
}

func (suite *PackageTestSuite) TestInitNewInMemManager() {
	manager, err := NewManager(InMemType, suite.backOffTime, suite.maxRetry)
	suite.NoError(err)
	suite.Equal(suite.inMemType, fmt.Sprint(reflect.TypeOf(manager)))
}

func (suite *PackageTestSuite) TestInitInvalidTypeManager() {
	_, err := NewManager("invalid-type", suite.backOffTime, suite.maxRetry)
	suite.Error(err)
}

func (suite *PackageTestSuite) TestGetInitialRetryCountInMem() {
	manager, err := NewManager(InMemType, suite.backOffTime, suite.maxRetry)
	suite.NoError(err)
	suite.Equal(suite.inMemType, fmt.Sprint(reflect.TypeOf(manager)))

	count := manager.GetRetryCount(suite.key)
	suite.Equal(0, count)
}

func (suite *PackageTestSuite) TestAddRetryCountInMem() {
	manager, err := NewManager(InMemType, suite.backOffTime, suite.maxRetry)
	suite.NoError(err)
	suite.Equal(suite.inMemType, fmt.Sprint(reflect.TypeOf(manager)))

	manager.AddRetryCount(suite.key)
	count := manager.GetRetryCount(suite.key)
	suite.Equal(1, count)
}

func (suite *PackageTestSuite) TestGetAddedRetryCountInMem() {
	manager, err := NewManager(InMemType, suite.backOffTime, suite.maxRetry)
	suite.NoError(err)
	suite.Equal(suite.inMemType, fmt.Sprint(reflect.TypeOf(manager)))

	count := manager.GetRetryCount(suite.key)
	suite.Equal(0, count)
}

func (suite *PackageTestSuite) TestClearRetryCountInMem() {
	manager, err := NewManager(InMemType, suite.backOffTime, suite.maxRetry)
	suite.NoError(err)
	suite.Equal(suite.inMemType, fmt.Sprint(reflect.TypeOf(manager)))

	manager.AddRetryCount(suite.key)
	manager.ClearRetryCount(suite.key)
	count := manager.GetRetryCount(suite.key)
	suite.Equal(0, count)
}

func (suite *PackageTestSuite) TestDelayProcessFollowBackOffTimeInMem() {
	manager, err := NewManager(InMemType, suite.backOffTime, suite.maxRetry)
	suite.NoError(err)
	suite.Equal(suite.inMemType, fmt.Sprint(reflect.TypeOf(manager)))

	manager.AddRetryCount(suite.key)

	secBefore := time.Now()
	manager.DelayProcessFollowBackOffTime(suite.key)
	secAfter := time.Now()
	suite.Equal(time.Duration(suite.backOffTime)*time.Second, secAfter.Sub(secBefore).Truncate(time.Second))
}

func (suite *PackageTestSuite) TestIsMaximumRetryInMem() {
	manager, err := NewManager(InMemType, suite.backOffTime, suite.maxRetry)
	suite.NoError(err)
	suite.Equal(suite.inMemType, fmt.Sprint(reflect.TypeOf(manager)))

	manager.AddRetryCount(suite.key)
	manager.AddRetryCount(suite.key)
	manager.AddRetryCount(suite.key)

	suite.Equal(true, manager.IsMaximumRetry(suite.key))
}

func (suite *PackageTestSuite) TestIsNotMaximumRetryInMem() {
	manager, err := NewManager(InMemType, suite.backOffTime, suite.maxRetry)
	suite.NoError(err)
	suite.Equal(suite.inMemType, fmt.Sprint(reflect.TypeOf(manager)))

	manager.AddRetryCount(suite.key)
	manager.AddRetryCount(suite.key)

	suite.Equal(false, manager.IsMaximumRetry(suite.key))
}

func TestAppSuite(t *testing.T) {
	suite.Run(t, new(PackageTestSuite))
}
