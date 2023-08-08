package maotai

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Listen  string `env:""`
	AppName string `env:""`

	e *gin.Engine
}

func (s *Server) SetDefaults() {
	// compatible for heroku
	port := os.Getenv("PORT")
	if port != "" {
		s.Listen = fmt.Sprintf(":%s", port)
	}

	if s.Listen == "" {
		s.Listen = ":8080"
	}

	if s.AppName == "" {
		s.AppName = "appname"
	}

	if s.e == nil {
		s.e = s.defaultEngine()
	}
}

func (s *Server) defaultEngine() *gin.Engine {

	engine := gin.New()
	logger := gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"", "/", "/ping", "/liveness"},
	})
	engine.Use(logger, gin.Recovery())

	return engine
}

func (s *Server) WithEngine(e *gin.Engine) {
	s.e = e
}

func (s *Server) Engine() *gin.Engine {
	return s.e
}

func (s *Server) Run() {

	s.e.Run(s.Listen)
}

func (s *Server) RegisteRouter(f RouterGroupFunc) {
	s.e.Handle(http.MethodGet, "", pingHandler)
	// s.e.Handle(http.MethodGet, "/ping", pingHandler)
	// s.e.Handle(http.MethodGet, "/liveness", pingHandler)

	r := s.e.Group(s.AppName)
	f(r)
}
