package transport

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nilspolek/leaderBoard/repo"
)

const (
	DefaultAddr = ":8080"
)

type Server struct {
	router *mux.Router
	repo   repo.Repo
	Config ServerConfig
}

type ServerConfig struct {
	Addr string
}

func New(repo repo.Repo, conf ServerConfig) *Server {
	if conf.Addr == "" {
		conf.Addr = DefaultAddr
	}
	return &Server{
		router: mux.NewRouter(),
		repo:   repo,
		Config: conf,
	}
}

func (s *Server) SetupRoutes() {
	var endpoints Endpoints = Endpoints{
		repo: s.repo,
	}
	s.AddRoute("GET", "/user/{id}", endpoints.GetUser)
	s.AddRoute("POST", "/user", endpoints.CreateUser)
	s.AddRoute("GET", "/game/{id}", endpoints.GetGameScore)
	s.AddRoute("POST", "/game", endpoints.CreateGame)
	s.AddRoute("GET", "/team/{id}", endpoints.GetTeam)
	s.AddRoute("POST", "/team", endpoints.CreateTeam)
}

func (s *Server) AddRoute(method, path string, handler http.HandlerFunc) {
	s.router.HandleFunc(path, handler).Methods(method)
}

func (s *Server) ServeHTTP() {
	http.Handle("/", s.router)
	http.ListenAndServe(s.Config.Addr, nil)
}
