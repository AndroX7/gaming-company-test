package http

import (
	"gaming-company-test/app/api/middleware"
	"gaming-company-test/service/artist"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	artistUsecase artist.Usecase
}

func New(artistUC artist.Usecase) *Handler {
	return &Handler{
		artistUsecase: artistUC,
	}
}

func (h *Handler) Register(r *gin.Engine, m *middleware.Middleware) {
	artistRoute := r.Group("/artist", m.BasicHandle())
	{
		artistRoute.GET("", h.Index)
		artistRoute.GET("/:id", h.Show)
		artistRoute.POST("", h.Create)
		artistRoute.DELETE("/:id", h.Delete)
		artistRoute.PUT("/:id", h.Update)
	}
}
