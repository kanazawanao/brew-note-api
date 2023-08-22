package models

import (
	"time"

	"gorm.io/gorm"
)

type Store struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	StoreType string         `json:"storeType"`
	Address   string         `json:"address"`
	Url       string         `json:"url"`
	CreatedAt *time.Time     `json:"createdAt"`
	UpdatedAt *time.Time     `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
