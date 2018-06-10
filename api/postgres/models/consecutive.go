package models

import (
	"github.com/jinzhu/gorm"
)

// Consecutive ...
type Consecutive struct {
	ID string `gorm:"primary_key"`

	// FK
	ConsecutiveType ConsecutiveType
	Type            string `gorm:"not null"`
	// FK

	Consecutive  string `gorm:"not null"` //Summary of all entities are using this consecutive
	HasPrefix    bool   `gorm:"not null; default:true"`
	Prefix       string `gorm:"unique; not null"`
	InitialRange string `gorm:"not null"`
	FinalRange   string `gorm:"not null"`
}

// TableName return a table name :v
func (Consecutive) TableName() string { return "Consecutive" }

// Save ...
func (c *Consecutive) Save(db *gorm.DB) error {

	err := db.Create(c).Error

	if err != nil {
		return err
	}

	return nil
}

// Update ...
func (c *Consecutive) Update(db *gorm.DB) error {

	row := Consecutive{}

	row.InitialRange = c.InitialRange
	row.FinalRange = c.FinalRange

	err := db.Model(c).Updates(row).Error

	if err != nil {
		return err
	}

	return nil
}

// AddConsecutive ...
func (c *Consecutive) AddConsecutive(db *gorm.DB) error {

	row := Consecutive{}

	// ADD +1
	row.Consecutive = c.Consecutive

	err := db.Model(c).Updates(row).Error

	if err != nil {
		return err
	}

	return nil
}
