package http

import (
	"gaming-company-test/service/artist/delivery/http/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var request request.ArtistCreateRequest

	if err := c.ShouldBind(&request); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	artistM, err := h.artistUsecase.Create(request)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, artistM)
}
