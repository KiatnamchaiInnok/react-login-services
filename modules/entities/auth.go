package entities

import (
	"github.com/golang-jwt/jwt/v4"
)

type UsersCredentials struct {
	Username string `json:"username" db:"username" form:"username"`
	Password string `json:"password" db:"password" form:"password"`
}

type UsersPassport struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UsersClaims struct {
	Id       int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type UsersLoginRes struct {
	AccessToken string `json:"access_token"`
}
