package consecutive

import "github.com/jinzhu/gorm"

// Consecutive ...
type Consecutive struct {
	ID string `gorm:"primary_key"`

	// FK
	//ConsecutiveType ConsecutiveType
	//Type            string `gorm:"not null"`
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

// Update function ONLY update the InitialRange and FinalRange
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

// Increase function increase the consecutive count ++
func (c *Consecutive) Increase(db *gorm.DB) error {

	row := Consecutive{}

	row.Consecutive = c.Consecutive

	err := db.Model(c).Updates(row).Error

	if err != nil {
		return err
	}

	return nil

}

// Registred ... NAMEEE
func (c *Consecutive) Registred(db *gorm.DB) (string, error) {

	var row Consecutive

	err := db.Where("ID = ? AND consecutive = ?", c.ID, c.Consecutive).First(&row).Error

	if err != nil {
		return "", err
	}

	return row.Consecutive, nil
}

// Consecutives function return all consecutives
func (c *Consecutive) Consecutives(db *gorm.DB) ([]Consecutive, error) {

	var consecutives []Consecutive

	err := db.Find(&consecutives).Error

	if err != nil {
		return nil, err
	}

	return consecutives, nil
}

// ConsecutiveGet ... NAMEEE
func (c *Consecutive) ConsecutiveGet(db *gorm.DB) (Consecutive, error) {

	var consecutive Consecutive

	err := db.Where("ID = ?", c.ID).First(&consecutive).Error

	if err != nil {
		return consecutive, err
	}

	return consecutive, nil
}
