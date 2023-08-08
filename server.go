package maotai

import (
	"fmt"
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
		s.AppName = "demo"
	}

	if s.e == nil {
		s.e = gin.Default()
	}
}

func (s *Server) Initialize() {
}

// func (s *Server) WithAppName(name string) {
// 	s.AppName = name
// }

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
	// f(s.e)
	r := s.e.Group(s.AppName)
	f(r)
}
