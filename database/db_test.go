package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/sweetrpg/common.go/logging"
	"github.com/sweetrpg/db.go/constants"
	"go.mongodb.org/mongo-driver/bson"
)

type DbTestSuite struct {
	suite.Suite
}

type DBObject struct {
	Key   string `bson:"key"`
	Value string `bson:"value"`
}

func (suite *DbTestSuite) SetupTest() {
	os.Unsetenv(constants.DB_URI)
	os.Unsetenv(constants.DB_NAME)
	os.Unsetenv(constants.DB_HOST)
	os.Unsetenv(constants.DB_SCHEME)
	os.Unsetenv(constants.DB_USER)
	os.Unsetenv(constants.DB_PW)
	os.Unsetenv(constants.DB_PORT)
	os.Unsetenv(constants.DB_OPTS)
}

func (suite *DbTestSuite) TestBuildURLFromURI() {
	os.Setenv(constants.DB_URI, "mongo://user:pass@host:12345/db?opts=these")
	dbUrl, dbName := buildDbURL()
	assert.Equal(suite.T(), "mongo", dbUrl.Scheme)
	assert.Equal(suite.T(), "user", dbUrl.User.Username())
	// assert.Equal(t, dbUrl.User.Password(), "pass")
	assert.Equal(suite.T(), "host:12345", dbUrl.Host)
	assert.Equal(suite.T(), "these", dbUrl.Query().Get("opts"))
	assert.Equal(suite.T(), "db", dbName)
}

func (suite *DbTestSuite) TestBuildURLFromParts() {
	os.Setenv(constants.DB_NAME, "db")
	os.Setenv(constants.DB_HOST, "host")
	os.Setenv(constants.DB_SCHEME, "mongo")
	os.Setenv(constants.DB_USER, "user")
	os.Setenv(constants.DB_PW, "pass")
	os.Setenv(constants.DB_PORT, "12345")
	os.Setenv(constants.DB_OPTS, "opts=these")

	dbUrl, dbName := buildDbURL()
	assert.Equal(suite.T(), "mongo", dbUrl.Scheme)
	assert.Equal(suite.T(), "user", dbUrl.User.Username())
	// assert.Equal(t, dbUrl.User.Password(), "pass")
	assert.Equal(suite.T(), "host:12345", dbUrl.Host)
	assert.Equal(suite.T(), "these", dbUrl.Query().Get("opts"))
	assert.Equal(suite.T(), "db", dbName)
}

func (suite *DbTestSuite) TestInvalidURL() {
	os.Setenv(constants.DB_URI, "bogus!this is some b4d URI^#$%")

	assert.Panics(suite.T(), func() { buildDbURL() }, "Should have panicked")
}

func (suite *DbTestSuite) TestInsert() {
	os.Setenv(constants.DB_URI, os.Getenv("TEST_DB_URI"))
	logging.Init()
	SetupDatabase()

	doc := DBObject{
		Key:   "inserted-key",
		Value: "inserted-value",
	}

	id, err := Insert[DBObject](os.Getenv("TEST_COLLECTION"), doc)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), id)
}

func (suite *DbTestSuite) TestUpdate() {
	os.Setenv(constants.DB_URI, os.Getenv("TEST_DB_URI"))
	logging.Init()
	SetupDatabase()

	doc := DBObject{
		Key:   "update-key",
		Value: "update-value",
	}

	id, err := Insert[DBObject](os.Getenv("TEST_COLLECTION"), doc)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), id)

	object, err := Get[DBObject](os.Getenv("TEST_COLLECTION"), id)
	assert.NotNil(suite.T(), object)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "update-key", object.Key)
	assert.Equal(suite.T(), "update-value", object.Value)

	newObject := DBObject{
		Key:   object.Key,
		Value: "changed-value",
	}

	matched, modified, err := Update[DBObject](os.Getenv("TEST_COLLECTION"), id, newObject)
	assert.Nil(suite.T(), err)
	assert.EqualValues(suite.T(), 1, matched)
	assert.EqualValues(suite.T(), 1, modified)

	object, err = Get[DBObject](os.Getenv("TEST_COLLECTION"), id)
	assert.NotNil(suite.T(), object)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "update-key", object.Key)
	assert.Equal(suite.T(), "changed-value", object.Value)
}

func (suite *DbTestSuite) TestDelete() {
	os.Setenv(constants.DB_URI, os.Getenv("TEST_DB_URI"))
	logging.Init()
	SetupDatabase()

	doc := DBObject{
		Key:   "deleted-key",
		Value: "deleted-value",
	}

	id, err := Insert[DBObject](os.Getenv("TEST_COLLECTION"), doc)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), id)

	deleted, err := Delete[DBObject](os.Getenv("TEST_COLLECTION"), id)
	assert.True(suite.T(), deleted)
	assert.NoError(suite.T(), err)
}

func (suite *DbTestSuite) TestGet() {
	os.Setenv(constants.DB_URI, os.Getenv("TEST_DB_URI"))
	logging.Init()
	SetupDatabase()

	doc := DBObject{
		Key:   "gotten-key",
		Value: "gotten-value",
	}

	id, err := Insert[DBObject](os.Getenv("TEST_COLLECTION"), doc)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), id)

	object, err := Get[DBObject](os.Getenv("TEST_COLLECTION"), id)
	assert.NotNil(suite.T(), object)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "gotten-key", object.Key)
	assert.Equal(suite.T(), "gotten-value", object.Value)
}

func (suite *DbTestSuite) TestQuery() {
	os.Setenv(constants.DB_URI, os.Getenv("TEST_DB_URI"))
	logging.Init()
	SetupDatabase()

	// insert docs
	for i := 0; i < 10; i++ {
		doc := DBObject{
			Key:   "key-" + string(i),
			Value: "value-" + string(i),
		}
		_, err := Insert[DBObject](os.Getenv("TEST_COLLECTION"), doc)
		assert.NoError(suite.T(), err)
	}

	filter := bson.D{{}}
	sort := bson.D{{Key: "Key", Value: 1}} // Sort by key ascending
	proj := bson.D{{Key: "Key", Value: 1}}
	var start int64 = 1
	limit := 2
	models, err := Query[DBObject](os.Getenv("TEST_COLLECTION"), filter, sort, proj, start, limit)
	assert.NotNil(suite.T(), models)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(models))

	model1 := models[0]
	assert.Equal(suite.T(), "key-1", model1.Key)
	assert.Nil(suite.T(), model1.Value)

	model2 := models[1]
	assert.Equal(suite.T(), "key-2", model2.Key)
	assert.Nil(suite.T(), model2.Value)
}

func TestDbTestSuite(t *testing.T) {
	suite.Run(t, new(DbTestSuite))
}
