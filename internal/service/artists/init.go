package artists

import (
	"context"
	"github/irvantaufik/album/internal/entity"
	artistsRepository "github/irvantaufik/album/internal/repository/artists"
)

type ArtistsService interface {
	FindById(ctx context.Context, id int64) (entity.Artists, error)
	FindAll(ctx context.Context) ([]entity.Artists, error)
	Create(ctx context.Context, artists *entity.Artists) (*entity.Artists, error)
}

type artistsService struct {
	artists artistsRepository.ArtistRepository
}

func InitArtistsService(artists artistsRepository.ArtistRepository) ArtistsService {
	return &artistsService{
		artists: artists,
	}
}
