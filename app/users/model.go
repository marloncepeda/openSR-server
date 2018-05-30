package users

// People ...
type People struct {
	tablename struct{} `sql:"public.peoples"`

	ID            uint   `gorm:"primary_key, AUTO_INCREMENT"`
	Consecutive   string `sql:"type:VARCHAR(100), not null"`
	Name          string `sql:"type:VARCHAR(100), not null"`
	SurName       string `sql:"type:VARCHAR(100), not null"`
	SecondSurName string `sql:"type:VARCHAR(100), not null"`
	Phone         string `sql:"type:VARCHAR(100), not null"`
	UserName      string `sql:"type:VARCHAR(100), not null"`
	Password      string `sql:"type:VARCHAR(100), not null"`
}
