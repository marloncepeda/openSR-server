package users

type loginModel struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type registerModel struct {
	Name          string `form:"name" json:"name" binding:"required"`
	Surname       string `form:"surname" json:"surname" binding:"required"`
	SecondSurname string `form:"secondSurname" json:"secondSurname" binding:"required"`
	Phone         uint   `form:"phone" json:"phone" binding:"required"`
	Username      string `form:"username" json:"username" binding:"required"`
	Password      string `form:"password" json:"password" binding:"required"`
}
