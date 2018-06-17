package jwt

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

//Encode ...
func Encode(id string) (string, error) {

	var key = fmt.Sprintf("%v", viper.Get("jwt.key"))

	token := jwt.New(jwt.SigningMethodHS256)

	claims := make(jwt.MapClaims)

	claims["foo"] = id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token.Claims = claims

	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
