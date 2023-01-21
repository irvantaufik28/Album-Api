package album

import (
	"context"

	"github.com/irvantaufik28/Album-Api/entity"
	albumRepository "github.com/irvantaufik28/Album-Api/entity/repository/album"
)

type AlbumService interface {
	GetAlbum(ctx context.Context, id int64) (entity.Album, error)
}

type albumService struct {
	album albumRepository.AlbumRepository
}

func InitAlbumService(album albumRepository.AlbumRepository) albumService {
	return &albumService{
		album: :album,
	}
}