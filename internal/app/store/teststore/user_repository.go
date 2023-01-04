package teststore

import (
	"github.com/smakkking/http-rest-api/internal/app/model"
	"github.com/smakkking/http-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}
	// здесь мы создали в БД запись о новом пользователе, кроме этого в объект ORM положили его id

	u.ID = len(r.users) + 1
	r.users[u.ID] = u

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, store.ErrUserNotFound
}

func (r *UserRepository) FindByID(id int) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, store.ErrUserNotFound
	}
	return u, nil
}
