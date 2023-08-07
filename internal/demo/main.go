package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/maotai"
	"github.com/go-jarvis/maotai/internal/demo/api"
	"github.com/tangx/envutils"
)

var s = &maotai.Server{}

func ping(c *gin.Context) {
}
func main() {
	s.WithEngine(gin.Default())

	maotai.AppendRouterGroup(s.Engine(), api.InitialGroup)

	s.Run()
}

func init() {

	config := &struct {
		Server *maotai.Server
	}{
		Server: s,
	}

	envutils.MustExport("demo", config)
	envutils.MustImport("demo", config)

}
