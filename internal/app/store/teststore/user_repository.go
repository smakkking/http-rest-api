package teststore

import (
	"github.com/smakkking/http-rest-api/internal/app/model"
	"github.com/smakkking/http-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}
	// здесь мы создали в БД запись о новом пользователе, кроме этого в объект ORM положили его id

	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, store.ErrUserNotFound
	}
	return u, nil
}
