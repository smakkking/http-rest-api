package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/smakkking/http-rest-api/internal/app/model"
	"github.com/smakkking/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := NewServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))

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

func TestServer_AuthenticateUser(t *testing.T) {
	store := teststore.New()
	u := model.TestUser(t)

	store.User().Create(u)

	testCases := []struct {
		name         string
		cookieValue  map[interface{}]interface{}
		expectedCode int
	}{
		{
			name: "auth",
			cookieValue: map[interface{}]interface{}{
				"user_id": u.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "not auth",
			cookieValue:  nil,
			expectedCode: http.StatusUnauthorized,
		},
	}

	s := NewServer(store, sessions.NewCookieStore([]byte("secret")))
	sc := securecookie.New([]byte("secret"), nil)
	myHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			cookieStr, _ := sc.Encode(sessionName, tc.cookieValue)
			req.Header.Set("Cookie", fmt.Sprintf("%s=%s", sessionName, cookieStr))
			s.authenticateUser(myHandler).ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleSessionsCreate(t *testing.T) {
	// изначально у нас бд голая, там нет никаких пользоваелей
	// поэтому их нужно создать
	u := model.TestUser(t)

	store := teststore.New()
	store.User().Create(u)

	s := NewServer(store, sessions.NewCookieStore([]byte("secret")))

	testCases := []struct {
		name          string
		payload       interface{}
		expected_code int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expected_code: http.StatusOK,
		},
		{
			name: "invalid pwd",
			payload: map[string]string{
				"email":    u.Email,
				"password": "3434254366",
			},
			expected_code: http.StatusUnauthorized,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email":    "asdflkfm@asdf.org",
				"password": u.Password,
			},
			expected_code: http.StatusUnauthorized,
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

			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expected_code, rec.Code)
		})
	}
}
