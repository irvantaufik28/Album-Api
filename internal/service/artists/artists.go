package artists

import (
	"context"
	"github/irvantaufik/album/internal/entity"
)

func (service *artistsService) FindById(ctx context.Context, id int64) (artists entity.Artists, err error) {
	artists, err = service.artists.FindById(ctx, id)
	if err != nil {
		return artists, err
	}

	return artists, nil
}

func (sercvice *artistsService) Create(ctx context.Context, artists *entity.Artists) (*entity.Artists, error) {
	var newArtists *entity.Artists

	err := sercvice.artists.Create(ctx, artists)

	if err != nil {
		return newArtists, err
	}
	return artists, err
}
