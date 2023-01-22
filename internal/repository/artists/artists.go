package repository

import (
	"context"
	"github/irvantaufik/album/internal/entity"
)

func (repository *artistsRepository) FindById(ctx context.Context, id int64) (entity.Artists, error) {
	return repository.artistsOrm.FindById(ctx, id)
}
