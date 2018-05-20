package users

import (
	"github.com/jinzhu/gorm"
)

// User .....
type User struct {
	gorm.Model
	//Code           int    `gorm:"AUTO_INCREMENT, not null"`
	Name string `gorm:"type:varchar(20), not null"`
	//FirstLastName  string `gorm:"type:varchar(20), not null"`
	//SecondLastName string `gorm:"type:varchar(20), not null"`
	//Phone          int    `gorm:"type:smallint(15), not null"`
	//NickName       string `gorm:"type:varchar(30), not null"`
	Role    string `gorm:"type:varchar(30), not null"`
	Picture string `gorm:"type:varchar(1000)"`
}
