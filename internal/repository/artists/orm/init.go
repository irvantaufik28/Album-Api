package orm

import (
	"context"
	"github/irvantaufik/album/internal/entity"

	"gorm.io/gorm"
)

type ArtistsOrm interface {
	FindById(ctx context.Context, id int64) (entity.Artists, error)
	Create(ctx context.Context, artist *entity.Artists) error
}

type artistsConnection struct {
	db *gorm.DB
}

func InitArstistOrm(db *gorm.DB) ArtistsOrm {
	return &artistsConnection{
		db: db,
	}
}
