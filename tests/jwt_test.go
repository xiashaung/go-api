package tests

import (
	"github.com/stretchr/testify/assert"
	"go-api/lib"
	"testing"
)

var tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkX2F0IjoxNjg1MDA0MjEyLCJpZCI6MTB9.c1SlOV9WBJfMc3ivMzi3stOidTYh3AUD9vk-iWm1xak"

func TestJwtEncode(t *testing.T) {
	var id = 10
	token := lib.JwtEncode(id, 1684997091)
	assert.NotEqual(t, "", token)
}

// certbot certonly --preferred-challenges dns --manual -d *.demitasse.cn --server https://acme-v02.api.letsencrypt.org/directory
func TestJwtDecode(t *testing.T) {
	claim := lib.JwtDecode(tokenString)
	assert.Equal(t, float64(10), claim["id"])
}
