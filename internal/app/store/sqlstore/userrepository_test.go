package sqlstore_test

import (
	"testing"

	"github.com/smakkking/http-rest-api/internal/app/model"
	"github.com/smakkking/http-rest-api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("users")

	s := sqlstore.New(db)

	u := model.TestUser(t)

	err := s.User().Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("users")

	s := sqlstore.New(db)

	// find not existed user
	_, err := s.User().FindByEmail("user@example.com")
	assert.Error(t, err)

	// find existed user
	s.User().Create(model.TestUser(t))

	u2, err := s.User().FindByEmail("try@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
