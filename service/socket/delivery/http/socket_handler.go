package http

import (
	"gaming-company-test/app/api/middleware"
	"gaming-company-test/service/socket"
	"log"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

type Handler struct {
	socketUsecase socket.Usecase
}

func New(socketUC socket.Usecase) *Handler {
	return &Handler{
		socketUsecase: socketUC,
	}
}

func (h *Handler) Register(r *gin.Engine, m *middleware.Middleware) {
	r.GET("/socket.io/", gin.WrapH(h.Socket()))
	r.POST("/socket.io/", gin.WrapH(h.Socket()))
	r.Handle("WS", "/socket.io/", gin.WrapH(h.Socket()))
	r.Handle("WSS", "/socket.io/", gin.WrapH(h.Socket()))
}

func (h *Handler) Socket() *socketio.Server {
	s, _ := h.socketUsecase.Socket()
	return s
}

func (h *Handler) Test(c *gin.Context) {
	log.Println(c.Request.Header)
}
