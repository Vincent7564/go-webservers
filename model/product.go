package model

import "gorm.io/gorm"

type Product struct {
	Name  string
	Price int8
	gorm.Model
}
