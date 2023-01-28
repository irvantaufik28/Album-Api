package artists

import (
	artistsService "github/irvantaufik/album/internal/service/artists"

	"github.com/gin-gonic/gin"
)

type ArtistController interface {
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type artirsController struct {
	artists artistsService.ArtistsService
}

func InitArtistsController(artists artistsService.ArtistsService) ArtistController {
	return &artirsController{
		artists: artists,
	}
}
