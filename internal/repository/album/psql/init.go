package psql

import (
	"context"
	"database/sql"
	"workingwithdatabase/internal/entity"
)

type AlbumPSQL interface {
	Get(ctx context.Context, id int64) (album entity.Album, err error)
	// Create(ctx context.Context, album entity.Album) (entity.Album, error)
}

type albumConnection struct {
	db *sql.DB
}

func InitAlbumPSQL(db *sql.DB) AlbumPSQL {
	return &albumConnection{
		db: db,
	}
}
