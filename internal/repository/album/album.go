package repository

import (
	"context"

	"github.com/irvantaufik28/Album-Api/internal/entity"
)

func (repository *albumRepository) Get(ctx context.Context, id int64) (entity.Album, error) {
	return repository.albumPSQL.Get(ctx, id)
}
