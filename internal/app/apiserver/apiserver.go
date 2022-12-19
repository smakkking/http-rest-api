package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// Upcase to avaiable import from another package
func New(c *Config) *APIserver {
	return &APIserver{
		config: c,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (a *APIserver) Start() error {
	if err := a.configLogger(); err != nil {
		return err
	}

	a.logger.Info("startng api server")
	a.configRouter()

	return http.ListenAndServe(a.config.BindAddr, a.router)
}

func (s *APIserver) configLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIserver) configRouter() {
	// здесь задаем ресурсы и их обработчики
	s.router.HandleFunc("/hello", s.handlehello())
}

func (s *APIserver) handlehello() http.HandlerFunc {
	// здесь можно определять локальные переменные и делать вычисления

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
