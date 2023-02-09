package utils

import (
	"github.com/ywh147906/load-test/common/values"

	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")

type Claims struct {
	TokenInfo
	jwt.StandardClaims
}

type TokenInfo struct {
	MapId  values.Integer
	RoleId values.RoleId
}
