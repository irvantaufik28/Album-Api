package config

import (
	artistsRepository "github/irvantaufik/album/internal/repository/artists"

	"gorm.io/gorm"
)

type Repository struct {
	ArtistsRepository artistsRepository.ArtistRepository
}

// Function to initialize repository
func InitRepository(db *gorm.DB) Repository {
	return Repository{
		ArtistsRepository: artistsRepository.InArtistRepository(db),
	}
}
