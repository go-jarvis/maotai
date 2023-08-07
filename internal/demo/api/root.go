package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/maotai"
	"github.com/go-jarvis/maotai/internal/demo/api/user"
)

func InitialGroup(r gin.IRouter) {
	v1 := r.Group("/v1")

	maotai.AppendRouterGroup(v1, user.InitialGroup)
}
