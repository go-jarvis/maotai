package user

import "github.com/gin-gonic/gin"

func InitialGroup(r gin.IRouter) {
	userRG := r.Group("/user")

	userRG.GET("")
}
