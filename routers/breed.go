package routers

import (
	"github.com/gin-gonic/gin"

	"breed/database"
	breedHandler "breed/feature/breed/delivery/http"
	breedRepository "breed/feature/breed/repository"
	breedUsecase "breed/feature/breed/usecase"
)

func BreedRoute(route *gin.RouterGroup) {
	// repo
	newBreedRepo := breedRepository.NewBreedRepository(database.DB)

	// usecase
	newBreedUsecase := breedUsecase.NewBreedUsecase(newBreedRepo)

	// handler
	newBreedHandler := breedHandler.NewBreedHandler(newBreedUsecase)

	{
		route.POST("/breed-inquiry", newBreedHandler.ListBreed)
	}
}
