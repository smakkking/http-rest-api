package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //  установка драйвера postgres через пустое имя
)

type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

func New(c *Config) *Store {
	return &Store{
		config: c,
	}
}

func (s *Store) Open() error {
	fmt.Println(s.config.DataBaseDriver)
	fmt.Println(s.config.DataBaseURL)

	db, err := sql.Open(s.config.DataBaseDriver, s.config.DataBaseURL) // соединение создается лениво

	if err != nil {

		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	// когда сервак заканчивает работу, мы можем сюда поместить доп операции
	s.db.Close()
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}
