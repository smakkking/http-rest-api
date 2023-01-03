package sqlstore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

// пишем тестовый хелпер, он будет подключаться к реальной бд

func TestDB(t *testing.T, databaseUrl string) (*sql.DB, func(...string)) {
	t.Helper() // указываем, что это хелпер, т.е. его не нужно тестировать и где либо учитывать

	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}

		}
		db.Close()
	}
}
