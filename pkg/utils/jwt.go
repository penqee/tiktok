package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"tiktok/pkg/constants"
	"time"
)

var SecretKey = []byte("penQee")

type Claims struct {
	Uid      int64  `json:"uid"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

// GenerateToken 签发給用户token
func GenerateToken(uid int64, username string) (accessToken, refreshToken string, err error) {
	now := time.Now()
	acExpireTime := now.Add(constants.AccessTokenExpireDuration)
	reExpireTime := now.Add(constants.RefreshTokenExpireDuration)
	claims := Claims{
		Uid:      uid,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: acExpireTime.Unix(),
		},
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(SecretKey)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: reExpireTime.Unix(),
	}).SignedString(SecretKey)
	if err != nil {
		return "", "", err
	}

	return
}

// ParseToken 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token,
		&Claims{}, func(t *jwt.Token) (interface{}, error) { return SecretKey, nil })
	if tokenClaims != nil {
		if Claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return Claims, nil
		}
	}
	return nil, err
}

// ParseRefreshToken 验证用户token
func ParseRefreshToken(aToken, rToken string) (newAToken, newRToken string,
	err error) {
	accessClaim, err := ParseToken(aToken)
	if err != nil {
		return "", "", nil
	}

	refreshClaim, err := ParseToken(rToken)
	if err != nil {
		return "", "", nil
	}

	if accessClaim.ExpiresAt > time.Now().Unix() {
		return GenerateToken(accessClaim.Uid, accessClaim.UserName)
	}

	if refreshClaim.ExpiresAt > time.Now().Unix() {
		return GenerateToken(accessClaim.Uid, accessClaim.UserName)
	}

	return "", "", errors.New("身份过期，请重新登陆")

}
