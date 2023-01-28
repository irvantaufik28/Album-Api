package main

import (
	"fmt"
	"github/irvantaufik/album/internal/config"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Get the environment
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Load db config
	db := config.OpenDB(os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	defer config.CloseDB(db)

	// Initialize gin
	r := gin.Default()

	port := os.Getenv("PORT")

	// Init clean arch
	repository := config.InitRepository(db)
	service := config.InitService(repository.ArtistsRepository)
	controller := config.InitController(service.ArtistsService)

	// Create the API
	artistsRoutes := r.Group("/api/v1/artists")
	{
		artistsRoutes.GET("/:id", controller.ArtistController.FindById)
		artistsRoutes.POST("/", controller.ArtistController.Create)

	}

	// Run the gin gonic in port 5000
	runWithPort := fmt.Sprintf("0.0.0.0:%s", port)
	r.Run(runWithPort)
}
