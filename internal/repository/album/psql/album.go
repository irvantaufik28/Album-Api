package psql

import (
	"context"
	"workingwithdatabase/internal/entity"
)

func (repository *albumConnection) Get(ctx context.Context, id int64) (entity.Album, error) {
	var album entity.Album
	query := `SELECT id, artist_id, title, price From album WHERE id=$1`

	if err := repository.db.QueryRowContext(ctx, query, id).Scan(&album.ID, &album.Artist_id, &album.Title, &album.Price); err != nil {
		return album, err
	}

	return album, nil
}
