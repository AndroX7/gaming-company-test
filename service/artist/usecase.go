package artist

import (
	"gaming-company-test/lib/request_util"
	"gaming-company-test/models"
	"gaming-company-test/service/artist/delivery/http/request"

	"gaming-company-test/lib/response_util"

	"github.com/gin-gonic/gin"
)

type Usecase interface {
	Index(config request_util.PaginationConfig, c *gin.Context) ([]models.Artist, response_util.PaginationMeta, error)
	Show(artistID uint64) (*models.Artist, error)
	Create(request request.ArtistRequest) (*models.Artist, error)
	Update(request request.ArtistRequest, artistID uint64) (*models.Artist, error)
	Delete(artistID uint64) error
}
