package config

import (
	artistsController "github/irvantaufik/album/internal/controller/artists"
	artistService "github/irvantaufik/album/internal/service/artists"
)

type Controller struct {
	ArtistController artistsController.ArtistController
}

// Function to initialize handler
func InitController(artistService artistService.ArtistsService) Controller {
	return Controller{
		ArtistController: artistsController.InitArtistsController(artistService),
	}
}
