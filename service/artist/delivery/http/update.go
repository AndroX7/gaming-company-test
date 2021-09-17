package http

import (
	"gaming-company-test/service/artist/delivery/http/request"
	"gaming-company-test/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Update(c *gin.Context) {
	var request request.ArtistUpdateRequest

	artistID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(errors.ErrUnprocessableEntity).SetType(gin.ErrorTypePublic)
	}

	// validate request
	if err := c.ShouldBind(&request); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	artistM, err := h.artistUsecase.Update(request, artistID)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, artistM)
}
