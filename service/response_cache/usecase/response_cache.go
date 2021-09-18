package usecase

import (
	"gaming-company-test/lib/redis"
	"gaming-company-test/service/artist"
	"gaming-company-test/service/response_cache"
	"os"
)

type Usecase struct {
	redis            redis.Client
	routeGroups      map[string]string
	artistRepository artist.Repository
}

func New(
	artistRepository artist.Repository,
) response_cache.Usecase {
	redis := redis.NewClient(redis.Credentials{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}, os.Getenv("APP_ENV"))

	routeGroups := map[string]string{
		"admin":  "/admin",
		"user":   "",
		"public": "/public",
		"server": "/server",
	}

	return &Usecase{
		redis:            redis,
		routeGroups:      routeGroups,
		artistRepository: artistRepository,
	}
}
