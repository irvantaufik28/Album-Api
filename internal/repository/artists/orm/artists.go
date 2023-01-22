package orm

import (
	"context"
	"github/irvantaufik/album/internal/entity"
)

func (repository *artistsConnection) FindById(ctx context.Context, id int64) (entity.Artists, error) {
	var artist entity.Artists
	if err := repository.db.Find(&artist).Error; err != nil {
		return artist, err
	}
	return artist, nil

}
