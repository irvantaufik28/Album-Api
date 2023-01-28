package repository

import (
	"context"
	"github/irvantaufik/album/internal/entity"
	"github/irvantaufik/album/internal/repository/artists/orm"

	"gorm.io/gorm"
)

type ArtistRepository interface {
	FindById(ctx context.Context, id int64) (entity.Artists, error)
	FindAll(ctx context.Context) ([]entity.Artists, error)
	Create(ctx context.Context, artist *entity.Artists) error
}

type artistsRepository struct {
	artistsOrm orm.ArtistsOrm
}

func InArtistRepository(db *gorm.DB) ArtistRepository {
	return &artistsRepository{
		artistsOrm: orm.InitArstistOrm(db),
	}
}
