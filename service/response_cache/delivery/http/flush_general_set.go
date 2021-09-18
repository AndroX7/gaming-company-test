package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gaming-company-test/service/response_cache/delivery/http/request"
)

func (h *Handler) FlushAllSet(c *gin.Context) {
	var req request.ResponseCacheFlushRequest
	if err := c.ShouldBind(&req); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	go h.responseCacheUsecase.FlushAllFromSet(req.SetName)

	c.JSON(http.StatusOK, nil)
}
