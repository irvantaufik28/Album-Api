package repository

import (
	"context"

	"github.com/irvantaufik28/Album-Api/internal/entity"
	"github.com/irvantaufik28/Album-Api/internal/repository/album/psql"
)

type AlbumRepository interface {
	Get(ctx context.Context, id int64) (entity.Album, error)
	// Create(ctx context.Context, album entity.Album) (entity.Album, error)

	// GetCache(ctx context.Context, id int64) error
	// SetCache(ctx context.Context, album entity.Album) error
}

type albumRepository struct {
	albumPSQL psql.AlbumPSQL
}

func InitAlbumRepository(albumPSQL psql.AlbumPSQL) AlbumRepository {
	return &albumRepository{
		albumPSQL: albumPSQL,
	}
}
