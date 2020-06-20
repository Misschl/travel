package service

import (
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(uid int, secret string, maxAge int64) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		//"exp":  time.Now().Add(time.Minute * 15).Unix(),
		"exp": maxAge,
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string, secret string) (float64, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}
	return claim.Claims.(jwt.MapClaims)["uid"].(float64), err
}
