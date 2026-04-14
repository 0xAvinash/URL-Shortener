package auth

import (
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/argon2"
)

type TokenPair struct {
	AccessToken string
	RefreshToken string
}


var jwtSecret = []byte("secret")

func HashPassword(password string) string {
	hash := argon2.IDKey([]byte(password), jwtSecret, 1, 64*1024, 4, 32)
	return base64.StdEncoding.EncodeToString(hash)
}

func sign(token *jwt.Token) (string, error) {
    return token.SignedString(jwtSecret)
}


func generateTokens(userID uint64) (TokenPair, error) {
    access := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(15 * time.Minute).Unix(),
    })

    refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
    })

    accessToken, err := sign(access)
    if err != nil {
        return TokenPair{}, err
    }

    refreshToken, err := sign(refresh)
    if err != nil {
        return TokenPair{}, err
    }

    return TokenPair{
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
    }, nil
}