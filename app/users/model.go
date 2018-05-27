package users

import (
	"github.com/jinzhu/gorm"
)

// Users ...
type Users struct {
	gorm.Model

	Consecutive   uint   `gorm:"AUTO_INCREMENT, not null"`
	Name          string `gorm:"type:varchar(20), not null"`
	SurName       string `gorm:"type:varchar(20), not null"`
	SecondSurName string `gorm:"type:varchar(20), not null"`
	Phone         uint   `gorm:"type:smallint(15), not null"`
	UserName      string `gorm:"type:varchar(10), not null"`
	Password      string `gorm:"type:varchar(15), not null"`
}

/*
	gorm.Model struct have those variables.

  		ID        uint `gorm:"primary_key"`
  		CreatedAt time.Time
  		UpdatedAt time.Time
  		DeletedAt *time.Time
*/

// Export ...
func Export(db *gorm.DB) {
	db.AutoMigrate(&Users{})

}
