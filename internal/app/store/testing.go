package store

import (
	"fmt"
	"strings"
	"testing"
)

// пишем тестовый хелпер, он будет подключаться к реальной бд

func TestStore(t *testing.T, databaseUrl string) (*Store, func(...string)) {
	t.Helper() // указываем, что это хелпер, т.е. его не нужно тестировать и где либо учитывать

	// подключаемся к бд, заводим основую переменную работы с настройками, как и в главном скрипте
	config := NewConfig()
	config.DataBaseURL = databaseUrl

	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}
		s.Close()
	}
}
