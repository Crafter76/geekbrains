package models

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	ID   int
	Text string `gorm:"text"`
}
