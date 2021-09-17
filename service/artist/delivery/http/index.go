package http

import (
	"gaming-company-test/lib/response_util"
	"gaming-company-test/service/artist/delivery/http/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Index(c *gin.Context) {
	artists, artistPagination, err := h.artistUsecase.Index(request.NewArtistPaginationConfig(c.Request.URL.Query()), c)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, response_util.IndexResponse{
		Data: artists,
		Meta: artistPagination,
	})
}
