package store_test

import (
	"testing"

	"github.com/smakkking/http-rest-api/internal/app/model"
	"github.com/smakkking/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	// find not existed user
	_, err := s.User().FindByEmail("user@example.com")
	assert.Error(t, err)

	// find existed user
	s.User().Create(model.TestUser(t))

	u2, err := s.User().FindByEmail("try@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
