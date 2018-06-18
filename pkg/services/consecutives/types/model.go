package types

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Model ....
type Model struct {
	ID          string `gorm:"primary_key"`
	Name        string `gorm:"primary_key"`
	Description string `gorm:"primary_key"`
}

// TableName ....
func (Model) TableName() string { return "consecutives_type" }

func (t *Model) save(db *gorm.DB) error {

	err := db.Create(t).Error

	if err != nil {
		return err
	}

	return nil
}

func (t *Model) last(db *gorm.DB) (string, error) {

	consecutiveType := Model{}

	err := db.Last(&consecutiveType).Error

	if err != nil {
		return "0", err
	}

	return consecutiveType.ID, nil
}

func (t *Model) types(db *gorm.DB) ([]Model, error) {

	result := []Model{}

	err := db.Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *Model) query(where, value string, db *gorm.DB) (Model, error) {

	result := Model{}

	var err error

	switch where {

	case "id":
		err = db.Where("id = ?", value).First(&result).Error

	case "name":
		err = db.Where("name = ?", value).First(&result).Error

	case "description":
		err = db.Where("description = ?", value).First(&result).Error

	default:
		err = errors.New("Invalid parameter")
	}

	if err != nil {
		return result, err
	}

	return result, nil
}
