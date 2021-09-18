package request

import (
	"gaming-company-test/lib/request_util"
	"time"
)

type ArtistCreateRequest struct {
	ArtistName  string     `json:"artist_name" form:"artist_name" binding:"required"`
	AlbumName   string     `json:"album_name" form:"album_name" binding:"required"`
	ImageUrl    string     `json:"image_url" form:"image_url" binding:"required"`
	ReleaseDate *time.Time `json:"release_date" form:"release_date" binding:"required" time_format:"2006-01-02"`
	Price       float64    `json:"price" form:"price" binding:"required"`
	SampleUrl   string     `json:"sample_url" form:"sample_url" binding:"required"`
}

type ArtistUpdateRequest struct {
	ArtistName  *string    `json:"artist_name" form:"artist_name"`
	AlbumName   *string    `json:"album_name" form:"album_name"`
	ImageUrl    *string    `json:"image_url" form:"image_url"`
	ReleaseDate *time.Time `json:"release_date" form:"release_date" time_format:"2006-01-02"`
	Price       *float64   `json:"price" form:"price"`
	SampleUrl   *string    `json:"sample_url" form:"sample_url"`
}

func NewArtistPaginationConfig(conditions map[string][]string) request_util.PaginationConfig {
	request_util.OverrideKey(conditions, "id", "artists.id")
	request_util.OverrideKey(conditions, "artist_name", "artists.artist_name")
	request_util.OverrideKey(conditions, "album_name", "artists.album_name")

	filterable := map[string]string{
		"artist.id":           request_util.IdType,
		"artists.album_name":  request_util.StringType,
		"artists.artist_name": request_util.StringType,
		"release_date":        request_util.DateType,
	}
	return request_util.NewRequestPaginationConfig(conditions, filterable)
}
