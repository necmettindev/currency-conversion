package middlewares

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/dgrijalva/jwt-go.v3"
)

func SetUserContext(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := stripBearer(c.Request.Header.Get("Authorization"))

		tokenClaims, _ := jwt.ParseWithClaims(
			token,
			&Claims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtSecret), nil
			},
		)

		if tokenClaims != nil {
			claims, ok := tokenClaims.Claims.(*Claims)
			if ok && tokenClaims.Valid {
				c.Set("user_id", claims.ID)
				c.Set("user_username", claims.Username)
				c.Request = setToContext(c, "user_id", claims.ID)
				c.Request = setToContext(c, "user_username", claims.Username)
			}
		}

		c.Next()
	}
}

func setToContext(c *gin.Context, key interface{}, value interface{}) *http.Request {
	return c.Request.WithContext(context.WithValue(c.Request.Context(), key, value))
}
