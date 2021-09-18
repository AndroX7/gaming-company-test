package database_transaction_callbacks

import (
	"gaming-company-test/events"
	"gaming-company-test/service/artist"
	"gaming-company-test/service/response_cache"

	"gorm.io/gorm"
)

type Client interface {
	FlushResponseCache(db *gorm.DB)
	FlushRedisKey(db *gorm.DB)
}

type Callback struct {
	db                   *gorm.DB
	events               events.Client
	responseCacheUsecase response_cache.Usecase
	artistUsecase        artist.Usecase
}

func New(
	db *gorm.DB,
	events events.Client,
	responseCacheUsecase response_cache.Usecase,
	artistUsecase artist.Usecase,
) Client {
	return &Callback{
		db:                   db,
		events:               events,
		responseCacheUsecase: responseCacheUsecase,
		artistUsecase:        artistUsecase,
	}
}
