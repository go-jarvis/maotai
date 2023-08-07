package maotai

import "github.com/gin-gonic/gin"

type Handler struct {
	Method   string
	Path     string
	Hanlders []gin.HandlerFunc
}

func NewHandler(method string, path string, hanlders ...gin.HandlerFunc) *Handler {

	return &Handler{
		Method:   method,
		Path:     path,
		Hanlders: hanlders,
	}
}

func Handle(r gin.IRoutes, h *Handler) {
	r.Handle(h.Method, h.Path, h.Hanlders...)
}
func HanldeAny(r gin.IRoutes, h *Handler) {
	r.Any(h.Path, h.Hanlders...)
}

/* RouterGroup */
// type IRouter = gin.IRouter
type RouterGroupFunc = func(r gin.IRouter)

func AppendRouterGroup(parent gin.IRouter, f RouterGroupFunc) {
	f(parent)
}

func Use(r gin.IRouter, middlewares ...gin.HandlerFunc) gin.IRoutes {
	return r.Use(middlewares...)
}
