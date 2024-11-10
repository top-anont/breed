package usecase

import (
	"breed/domain"
	requestForm "breed/domain/requestForm"
	responseForm "breed/domain/responseForm"
)

type breedUsecase struct {
	breedRepo domain.BreedRepository
}

func NewBreedUsecase(breedRepo domain.BreedRepository) domain.BreedUsecase {
	return &breedUsecase{breedRepo}
}

func (u *breedUsecase) ListBreed(req requestForm.ListBreedRequest) (res []responseForm.ListBreedResponse, err error) {

	return u.breedRepo.ListBreed(req)
}
