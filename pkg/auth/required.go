package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ctreminiom/openSR-server/pkg/auth/jwt"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthorizationHandler ..
func AuthorizationHandler() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")

		fmt.Println(token)

		//Check if the token has the correct format
		Bearer := "Bearer "

		if !strings.Contains(token, Bearer) {

			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Your request is not authorized",
			})

			c.Abort()
			return
		}
		//Check if the token has the correct format

		// Check if the token has the token format
		tokenWithOutBearer := strings.Split(token, Bearer)

		if len(tokenWithOutBearer) < 2 {

			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "An authorization token was not supplied",
			})
		}
		// Check if the token has the token format

		//Validate the token
		valid, err := jwt.Decode(tokenWithOutBearer[1])

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid authorization token",
			})

			c.Abort()
			return
		}
		//Validate the token

		c.Set("public_id", valid.Claims.(jwtGo.MapClaims)["foo"])
		c.Next()
	}

}
