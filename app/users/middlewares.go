package users

import "github.com/ctreminiom/scientific-logs-api/app/security"

func encrypt(user People) *People {

	data := People{
		ID:            security.EncryptWithAES(user.ID),
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
