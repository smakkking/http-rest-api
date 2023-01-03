package sqlstore

import (
	"database/sql"

	"github.com/smakkking/http-rest-api/internal/app/model"
	"github.com/smakkking/http-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}
	// здесь мы создали в БД запись о новом пользователе, кроме этого в объект ORM положили его id
	if err := r.store.db.QueryRow(
		"INSERT INTO users(email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email=$1",
		email,
	).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows { // специально проверяем, чтобы заменить на нашу ошибку
			return nil, store.ErrUserNotFound
		}
		return nil, err // если ошибка любая другая, то просто проводим ее наверх
	}

	return u, nil
}
