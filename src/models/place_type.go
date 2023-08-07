package models

import (
	"time"

	"gorm.io/gorm"
)

type PlaceType struct {
	ID        string
	Key       string
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt
}
