package models

import (
	"time"

	"gorm.io/gorm"
)

type Bean struct {
	ID             string         `json:"id"`
	UserId        string          `json:"userId"`
	ProductionArea string         `json:"productionArea"`
	PlantationName string         `json:"plantationName"`
	Kind           string         `json:"kind"`
	RoastLevel     string         `json:"roastLevel"`
	Price          string         `json:"price"`
	CreatedAt      *time.Time     `json:"createdAt"`
	UpdatedAt      *time.Time     `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt"`
}
