package album

import (
	"context"

	"github.com/irvantaufik28/Album-Api/internal/entity"
)

func (service *albumService) GetAlbum(ctx context.Context, id int64) (album entity.Album, err error) {
	album, err = service.album.Get(ctx, id)
	if err != nil {
		return album, err
	}
	return album, nil
}
