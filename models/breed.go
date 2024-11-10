package models

import "time"

type Breed struct {
	ID          string    `json:"id"`
	NameEN      string    `json:"name_en"`
	NameTH      string    `json:"name_th"`
	ShortName   string    `json:"short_name"`
	Remark      *string   `json:"remark"`
	CreatedAt   time.Time `json:"createdAt"`
	CreatedByID string    `json:"createdByID"`
	CreatedBy   string    `json:"createdBy"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UpdatedByID string    `json:"updatedByID"`
	UpdatedBy   string    `json:"updatedBy"`
}

func (m *Breed) TableName() string {
	return "breed"
}
