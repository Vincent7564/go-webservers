package model

import "gorm.io/gorm"

type Task struct {
	Title       string
	Description string
	Is_Active   bool
	gorm.Model
}
