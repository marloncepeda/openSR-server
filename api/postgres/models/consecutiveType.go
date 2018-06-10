package models

import (
	"github.com/jinzhu/gorm"
)

// ConsecutiveType ..
type ConsecutiveType struct {
	ID          string `gorm:"primary_key"`
	Name        string `gorm:"unique; not null"`
	Description string `gorm:"unique; not null"`
}

// TableName return a table name :v
func (ConsecutiveType) TableName() string { return "Consecutive_Type" }

// Save ..
func (t *ConsecutiveType) Save(db *gorm.DB) error {

	err := db.Create(t).Error

	if err != nil {
		return err
	}

	return nil

}

// Update ...
func (t *ConsecutiveType) Update(db *gorm.DB) error {

	row := ConsecutiveType{}

	row.Name = t.Name
	row.Description = t.Description

	err := db.Model(t).Updates(row).Error

	if err != nil {
		return err
	}

	return nil
}

// Delete ...
func (t *ConsecutiveType) Delete(db *gorm.DB) error {

	err := db.Delete(t).Error

	if err != nil {
		return err
	}

	return nil
}

// GetTypes ...
func (t *ConsecutiveType) GetTypes(db *gorm.DB) ([]ConsecutiveType, error) {

	var types []ConsecutiveType

	err := db.Find(&types).Error

	if err != nil {
		return nil, err
	}

	return types, nil

}

// GetType ...
func (t *ConsecutiveType) GetType(db *gorm.DB) (ConsecutiveType, error) {

	var result ConsecutiveType

	err := db.Where("ID = ?", t.ID).First(&result).Error

	if err != nil {
		return result, err
	}

	return result, nil
}
