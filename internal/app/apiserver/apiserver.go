package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/smakkking/http-rest-api/internal/app/store/sqlstore"
)

// логика управления сервером
func Start(c *Config) error {
	db, err := newDB(c.DataBaseUrl)

	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(c.SessionKey))

	s := NewServer(store, sessionStore)
	return http.ListenAndServe(c.BindAddr, s)
}

func newDB(database_url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", database_url)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
