package repository

import (
	"context"
	"workingwithdatabase/internal/entity"
	"workingwithdatabase/internal/repository/album/psql"
)

type AlbumRepository interface {
	Get(ctx context.Context, id int64) (entity.Album, error)
	Create(ctx context.Context, album entity.Album) (entity.Album, error)

	GetCache(ctx context.Context, id int64) error
	SetCache(ctx context.Context, album entity.Album) error
}

type albumRepository struct {
	albumPSQL psql.AlbumPSQL
}

func InitAlbumRepository(album albumRepository) albumRepository {
	return &albumRepository{
		albumPSQL: albumPSQL,
	}
}
