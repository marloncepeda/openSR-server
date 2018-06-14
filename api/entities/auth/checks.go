package auth

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/ctreminiom/openSR-server/api/entities/user"
	"github.com/ctreminiom/openSR-server/api/security"
	"github.com/ctreminiom/openSR-server/api/security/jwt"
)

func validate(header string, db *gorm.DB) (HTTPCode int, response string) {

	parameters := format(header)

	var sign user.User
	sign.UserName = encrypt(parameters[0])
	sign.Password = encrypt(parameters[1])

	confirmed, identification := sign.Login(db)

	if !confirmed {
		return http.StatusUnauthorized, "incorrect username or password"
	}

	return 200, jwt.Encode(identification)
}

func format(header string) []string {

	pattern := strings.NewReplacer("Basic", "", " ", "")

	cleanHeader := pattern.Replace(header)

	base64Header, _ := base64.StdEncoding.DecodeString(cleanHeader)

	stringHeader := string(base64Header[:])

	return strings.Split(stringHeader, ":")
}

func encrypt(text string) string { return security.Encrypt(text) }

func decrypt(text string) string { return security.Decrypt(text) }
