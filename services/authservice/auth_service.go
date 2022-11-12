package authservice

import (
	"time"

	"github.com/necmettindev/currency-conversion/models/user"

	jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

type AuthService interface {
	IssueToken(u user.User) (string, error)
	ParseToken(token string) (*Claims, error)
}

type authService struct {
	jwtSecret string
}

func NewAuthService(jwtSecret string) AuthService {
	return &authService{
		jwtSecret: jwtSecret,
	}
}

type Claims struct {
	Username string `json:"username"`
	ID       uint   `json:"id"`
	jwt.StandardClaims
}

func (auth *authService) IssueToken(u user.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := Claims{
		u.Username,
		u.ID,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Currency Conversion",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString([]byte(auth.jwtSecret))
}

func (auth *authService) ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(auth.jwtSecret), nil
		},
	)

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
