package domain

import (
	requestForm "breed/domain/requestForm"
	responseForm "breed/domain/responseForm"
)

type BreedUsecase interface {
	ListBreed(req requestForm.ListBreedRequest) (res []responseForm.ListBreedResponse, err error)
}

type BreedRepository interface {
	ListBreed(req requestForm.ListBreedRequest) (res []responseForm.ListBreedResponse, err error)
}
