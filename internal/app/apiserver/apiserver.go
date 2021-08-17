package apiserver

import (
	"experiments/internal/app/handler"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type APIServer struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	handler *handler.Handler
}

func New(config *Config) *APIServer {
	return &APIServer{
		config:  config,
		logger:  logrus.New(),
		router:  mux.NewRouter(),
		handler: handler.New(),
	}
}

func (server *APIServer) Start() error {
	if err := server.configureLogger(); err != nil {
		return err
	}
	server.configureRouter()

	server.logger.Info("Starting API Server on port ", server.config.BindPort)
	return http.ListenAndServe(server.config.BindPort, server.router)
}

func (server *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(server.config.LogLevel)
	if err != nil {
		return err
	}
	server.logger.SetLevel(level)
	return nil
}

func (server *APIServer) configureRouter() {
	server.router.HandleFunc("/health", server.handler.HandleHealth).Methods("GET")
	server.router.HandleFunc("/{address}/health", server.handler.HandleRequestAddressHealth).Methods("GET")

}
