package lib

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
)

type jwtObj struct {
}

var (
	Jwt = jwtObj{}
)

var (
	jwtKey = []byte("asdasdasd")
)

func (j jwtObj) Encode(id int, expireAt int64) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         id,
		"expired_at": expireAt,
	})
	str, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println(err.Error())
	}
	return str
}

func (j jwtObj) Decode(jwtString string) jwt.MapClaims {
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		log.Println(err.Error())
	}
	return token.Claims.(jwt.MapClaims)
}
