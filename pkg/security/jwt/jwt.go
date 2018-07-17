package jwt

import (
	"github.com/yushuailiu/easygin/pkg/config"
	"time"
	"crypto/md5"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	jwt.StandardClaims
}

func getJwtToken() string {
	jwtToken := config.GetConfig().Section("security").Key("jwt_token").String()

	return jwtToken
}

func GenerateToken(username string, password string) (string, error) {
	now := time.Now()

	expirTime := now.Add(time.Hour)

	usernameMd5 := md5.Sum([]byte(username))
	passwordMd5 := md5.Sum([]byte(password))
	claims := Claims{
		string(usernameMd5[:]),
		string(passwordMd5[:]),
		jwt.StandardClaims{
			ExpiresAt:	expirTime.Unix(),
			Issuer:		config.GetConfig().Section("").Key("name").String(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(getJwtToken())
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return getJwtToken(), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
