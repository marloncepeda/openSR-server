package users

// People ...
type People struct {
	tablename struct{} `sql:"public.peoples"`

	ID            string `sql:"type:VARCHAR(100), pk"`
	Consecutive   string `sql:"type:VARCHAR(100), not null"`
	Name          string `sql:"type:VARCHAR(100), not null"`
	SurName       string `sql:"type:VARCHAR(100), not null"`
	SecondSurName string `sql:"type:VARCHAR(100), not null"`
	Phone         string `sql:"type:VARCHAR(100), not null"`
	UserName      string `sql:"type:VARCHAR(100), unique, not null"`
	Password      string `sql:"type:VARCHAR(100), not null"`
}

type loginModel struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type registerModel struct {
	Name          string `form:"name" json:"name" binding:"required"`
	Surname       string `form:"surname" json:"surname" binding:"required"`
	SecondSurname string `form:"secondSurname" json:"secondSurname" binding:"required"`
	Phone         string `form:"phone" json:"phone" binding:"required"`
	Username      string `form:"username" json:"username" binding:"required"`
	Password      string `form:"password" json:"password" binding:"required"`
}
