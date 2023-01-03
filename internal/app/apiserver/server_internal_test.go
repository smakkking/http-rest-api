package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smakkking/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := NewServer(teststore.New())

	testCases := []struct {
		name          string
		payload       interface{}
		expected_code int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "asdf@example.org",
				"password": "34gfg",
			},
			expected_code: http.StatusCreated,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"password": "34",
			},
			expected_code: http.StatusUnprocessableEntity,
		},
		{
			name:          "invalid payload",
			payload:       "jsla;dflkjsd",
			expected_code: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}

			json.NewEncoder(b).Encode(tc.payload)

			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expected_code, rec.Code)
		})
	}

}
