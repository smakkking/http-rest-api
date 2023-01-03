package model_test

import (
	"testing"

	"github.com/smakkking/http-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
)

type test_data struct {
	n    int
	want int
}

func TestUser_BeforeCreate(t *testing.T) {
	testCases := []struct {
		name    string
		m       func() *model.User
		isValid bool
	}{
		{
			name: "wrong email",
			m: func() *model.User {
				u := model.TestUser(t)
				u.Email = "sdlkgdglk"
				return u
			},
			isValid: false,
		},
		{
			name: "wrong password",
			m: func() *model.User {
				u := model.TestUser(t)
				u.Password = "123"
				return u
			},
			isValid: false,
		},
		{
			name: "all ok",
			m: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.m().Validate())
			} else {
				assert.Error(t, tc.m().Validate())
			}
		})

	}
}

func TestUser_Validate(t *testing.T) {
	u := model.TestUser(t)

	err := u.Validate()
	assert.NoError(t, err)
}
