package commons

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthClaims struct {
	jwt.StandardClaims
	UserName string
}

func CreateToken(userName string) (string, error) {
	now := time.Now()
	expiredAt := now.Add(time.Minute * 15)
	mapClaims := AuthClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt.Unix(),
			IssuedAt:  now.Unix(),
			NotBefore: now.Unix(),
		},
		UserName: userName,
	}

	secretKey, _, err := GenerateKey()
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, mapClaims)

	return token.SignedString(secretKey)
}

func ClaimToken(token string) (*AuthClaims, error) {
	claim := AuthClaims{}
	_, err := jwt.ParseWithClaims(token, &claim, func(token *jwt.Token) (interface{}, error) {
		_, pub, _ := GenerateKey()
		return pub, nil
	})
	if err != nil {
		return nil, err
	}

	return &claim, nil
}
