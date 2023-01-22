package config

import (
	artistsRepository "github/irvantaufik/album/internal/repository/artists"
	artistsService "github/irvantaufik/album/internal/service/artists"
)

type Service struct {
	ArtistsService artistsService.ArtistsService
}

// Function to initialize usecase
func InitService(artistsRepository artistsRepository.ArtistRepository) Service {
	return Service{
		ArtistsService: artistsService.InitArtistsService(artistsRepository),
	}
}
