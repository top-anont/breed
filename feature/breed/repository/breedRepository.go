package repository

import (
	"breed/domain"
	requestForm "breed/domain/requestForm"
	responseForm "breed/domain/responseForm"
	"breed/models"
	"breed/utils"

	"gorm.io/gorm"
)

type breedRepository struct {
	db *gorm.DB
}

func NewBreedRepository(db *gorm.DB) domain.BreedRepository {
	return &breedRepository{db}
}

func (r *breedRepository) ListBreed(req requestForm.ListBreedRequest) (res []responseForm.ListBreedResponse, err error) {
	query := r.db.Model(&models.Breed{}).
		Select("id, name_th, name_en, short_name, remark")

	if req.Keyword != "" {
		keyWord := "%" + req.Keyword + "%"
		query = query.Where("name_th ILIKE ? OR name_en ILIKE ? ", keyWord, keyWord)
	}

	if len(req.IDs) > 0 {
		query = query.Where("id IN (?)", req.IDs)
	}

	if len(req.ShortNames) > 0 {
		query = query.Where("short_name IN (?)", req.ShortNames)
	}

	if err := query.Find(&res).Error; err != nil {
		utils.Logger.Error(err.Error())

		return nil, err
	}

	return res, nil
}
