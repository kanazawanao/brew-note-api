package models

import (
	"time"

	"gorm.io/gorm"
)

type Bean struct {
	ID             int            `json:"id" gorm:"AUTO_INCREMENT"`
	UserId         string         `json:"userId"`
	ProductionRegion string         `json:"ProductionRegion"`
	Kind           string         `json:"kind"`
	RoastLevelId   int32          `json:"roastLevelId"`
	Price          int            `json:"price"`
	Gram           int            `json:"gram"`
	CreatedAt      *time.Time     `json:"createdAt"`
	UpdatedAt      *time.Time     `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt"`
}
