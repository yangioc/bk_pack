package dto

import (
	"github.com/golang-jwt/jwt/v5"
)

// JWT 資料結構
type JwtStruct struct {
	Claims interface{} `json:"claims"`
	jwt.RegisteredClaims
}
