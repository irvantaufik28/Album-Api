package orm

import (
	"context"
	"github/irvantaufik/album/internal/entity"
	"log"
)

func (repository *artistsConnection) FindById(ctx context.Context, id int64) (entity.Artists, error) {
	var artist entity.Artists
	if err := repository.db.Where("id = ?", id).Take(&artist).Error; err != nil {
		return artist, err
	}
	return artist, nil
}

func (repository *artistsConnection) Create(ctx context.Context, Artists *entity.Artists) error {

	tx := repository.db.Create(Artists)
	if tx.Error != nil {
		return tx.Error
	}
	log.Print("Success create Artists")
	return nil

}
