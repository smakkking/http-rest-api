package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"github.com/smakkking/http-rest-api/internal/app/model"
	"github.com/smakkking/http-rest-api/internal/app/store"
)

const (
	sessionName = "abcde"
)

var (
	errIncorectEmailOrPassword = errors.New("incorrect email or password")
)

// более легковесная реализация сервера, может только обрабатывать запросы
type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

func NewServer(store store.Store, sStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        store,
		sessionStore: sStore,
	}

	s.configRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// здесь расположена основная логика работы с ресурсами API
func (s *server) configRouter() {
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/sessions", s.handleSessionCreate()).Methods("POST")
	// ...
}

func (s *server) handleSessionCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// проверка на правильность ввода всего, чтобы не крашнулся например json парсер
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err) // т.е мы здесь ошибку не просто в толкаем наверх, а закидываем в специальный обработчик
			return
		}

		// проверка ввода, аутентификация
		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			// либо пароль неправильный, либо не нашли по имени
			s.error(w, r, http.StatusUnauthorized, errIncorectEmailOrPassword)
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = u.ID
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// декодим входящий json
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err) // т.е мы здесь ошибку не просто в толкаем наверх, а закидываем в специальный обработчик
			return
		}
		// создаем пользователя
		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err) // аналогично
			return
		}
		u.Sanitize()

		// отсылаем ответ об успешном завершении
		s.respond(w, r, http.StatusCreated, u)
	}
}

// рендерим ошибку в процессе обработки ресурсов
func (s *server) error(w http.ResponseWriter, r *http.Request, status_code int, err error) {
	s.respond(w, r, status_code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, status_code int, data interface{}) {
	w.WriteHeader(status_code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
