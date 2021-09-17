package http

import (
	"gaming-company-test/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Delete(c *gin.Context) {
	artistID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(errors.ErrUnprocessableEntity).SetType(gin.ErrorTypePublic)
	}

	err = h.artistUsecase.Delete(artistID)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}
