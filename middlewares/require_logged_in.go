package middlewares

import (
	"net/http"
	"strings"

	"github.com/necmettindev/currency-conversion/controllers"

	"github.com/gin-gonic/gin"
	"gopkg.in/dgrijalva/jwt-go.v3"
)

// Claims object
type Claims struct {
	Username string `json:"username"`
	ID       uint   `json:"id"`
	jwt.StandardClaims
}

// Remove "Bearer " from "Authorization" token string
func stripBearer(tok string) (string, error) {
	if len(tok) > 6 && strings.ToLower(tok[0:7]) == "bearer " {
		return tok[7:], nil
	}
	return tok, nil
}

func RequireLoggedIn(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := stripBearer(c.Request.Header.Get("Authorization"))
		if err != nil {
			controllers.HTTPRes(c, http.StatusUnauthorized, err.Error(), nil)
			c.Abort()
			return
		}

		tokenClaims, err := jwt.ParseWithClaims(
			token,
			&Claims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtSecret), nil
			},
		)
		if err != nil {
			controllers.HTTPRes(c, http.StatusUnauthorized, err.Error(), nil)
			c.Abort()
			return
		}

		if tokenClaims != nil {
			claims, ok := tokenClaims.Claims.(*Claims)

			if ok && tokenClaims.Valid {
				c.Set("user_id", claims.ID)
				c.Set("user_username", claims.Username)
				c.Next()
				return
			}
		}

		controllers.HTTPRes(c, http.StatusUnauthorized, "Unauthorized", nil)
		c.Abort()
	}
}
