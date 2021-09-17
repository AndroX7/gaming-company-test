package http

import (
	"gaming-company-test/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Show(c *gin.Context) {
	artistID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(errors.ErrUnprocessableEntity).SetType(gin.ErrorTypePublic)
	}

	artistM, err := h.artistUsecase.Show(artistID)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, artistM)
}
