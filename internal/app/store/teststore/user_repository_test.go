package teststore_test

import (
	"testing"

	"github.com/smakkking/http-rest-api/internal/app/model"
	"github.com/smakkking/http-rest-api/internal/app/store"
	"github.com/smakkking/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()

	u := model.TestUser(t)

	err := s.User().Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()

	// find not existed user
	_, err := s.User().FindByEmail("user@example.com")
	assert.EqualError(t, err, store.ErrUserNotFound.Error())

	// find existed user
	s.User().Create(model.TestUser(t))

	u2, err := s.User().FindByEmail("try@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
