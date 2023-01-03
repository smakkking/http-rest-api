package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/smakkking/http-rest-api/internal/app/model"
	"github.com/smakkking/http-rest-api/internal/app/store"
)

// более легковесная реализация сервера, может только обрабатывать запросы
type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func NewServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
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
	// ...
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err) // т.е мы здесь ошибку не просто в толкаем наверх, а закидываем в специальный обработчик
			return
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err) // аналогично
			return
		}
		u.Sanitize()
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
