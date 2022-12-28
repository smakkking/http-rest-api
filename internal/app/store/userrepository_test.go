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

	u, err := s.User().Create(&model.User{
		Email: "user@mail.org",
	})

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
	u1, err := s.User().Create(&model.User{
		Email: "user@mail.org",
	})

	u2, err := s.User().FindByEmail("user@mail.org")
	assert.NoError(t, err)
	assert.Equal(t, u1, u2)
}
