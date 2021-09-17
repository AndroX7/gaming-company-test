package usecase

import (
	"fmt"
	"gaming-company-test/models"

	"gaming-company-test/lib/request_util"
	"gaming-company-test/lib/response_util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (u *Usecase) Index(paginationConfig request_util.PaginationConfig, c *gin.Context) ([]models.Artist, response_util.PaginationMeta, error) {
	meta := response_util.PaginationMeta{
		Offset: paginationConfig.Offset(),
		Limit:  paginationConfig.Limit(),
		Total:  0,
	}

	if search, ok := c.Request.URL.Query()["search"]; ok {
		paginationConfig.AddScope(func(db *gorm.DB) *gorm.DB {
			return db.Where("AND name like ? OR code like ? ", fmt.Sprint("%", search[0], "%"), fmt.Sprint("%", search[0], "%"))
		})
	}

	artists, err := u.artistRepo.FindAll(paginationConfig)
	if err != nil {
		return nil, meta, err
	}

	total, err := u.artistRepo.Count(paginationConfig)
	if err != nil {
		return nil, meta, err
	}
	meta.Total = total

	return artists, meta, nil
}
