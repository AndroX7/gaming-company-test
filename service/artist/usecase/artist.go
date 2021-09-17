package usecase

import (
	"gaming-company-test/lib/s3"
	"gaming-company-test/service/artist"
)

type Usecase struct {
	artistRepo artist.Repository
	s3         s3.S3Client
}

func New(
	artistRepo artist.Repository,
	s3 s3.S3Client,
) artist.Usecase {
	return &Usecase{
		artistRepo: artistRepo,
		s3:         s3,
	}
}
