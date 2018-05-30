package users

import "github.com/ctreminiom/scientific-logs-api/app/security"

func encrypt(user People) *People {

	data := People{
		ID:            user.ID,
		Consecutive:   security.EncryptWithAES(user.Consecutive),
		Name:          security.EncryptWithAES(user.Name),
		SurName:       security.EncryptWithAES(user.SurName),
		SecondSurName: security.EncryptWithAES(user.SecondSurName),
		Phone:         security.EncryptWithAES(user.Phone),
		UserName:      security.EncryptWithAES(user.UserName),
		Password:      security.EncryptWithAES(user.Name),
	}

	return &data
}

/*

	Consecutive:   12,
	Name:          json.Name,
	SurName:       json.Surname,
	SecondSurName: json.SecondSurname,
	Phone:         json.Phone,
	UserName:      json.Username,
	Password:      json.Password}

*/
