package models

import (
	"time"

	"gorm.io/gorm"
)

type Store struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Address   string         `json:"address"`
	Url       string         `json:"url"`
	PlaceId   string         `json:"placeId"`
	CreatedAt *time.Time     `json:"createdAt"`
	UpdatedAt *time.Time     `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
