package models

import (
	"time"

	"gorm.io/gorm"
)

type Bean struct {
	ID               int            `json:"id" gorm:"AUTO_INCREMENT"`
	UserId           string         `json:"userId"`
	ProductionRegion string         `json:"productionRegion"`
	Kind             string         `json:"kind"`
	RoastLevelId     int32          `json:"roastLevelId"`
	ProcessingId     int            `json:"processingId"`
	Altitude         int            `json:"altitude"`
	Farm             string         `json:"farm"`
	Flavor           string         `json:"flavor"`
	Memo             string         `json:"memo"`
	RoastedAt        *time.Time     `json:"roastedAt"`
	CreatedAt        *time.Time     `json:"createdAt"`
	UpdatedAt        *time.Time     `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `json:"deletedAt"`
}
