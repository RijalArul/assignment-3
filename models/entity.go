package models

import (
	"gorm.io/gorm"
)

type Element struct {
	gorm.Model
	Water       int
	Wind        int
	StatusWater string
	StatusWind  string
}
