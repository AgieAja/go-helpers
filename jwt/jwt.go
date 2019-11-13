package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//MyClaims - struct jwt
type MyClaims struct {
	jwt.StandardClaims
	UserID uint `json:"user_id"`
}

//GenerateTokenJwt - generate token jwt
func GenerateTokenJwt(userID uint, appName, appSecretKey string) (string, error) {
	loginExpDuration := time.Duration(24*60) * time.Minute
	jwtSigningMethod := jwt.SigningMethodHS256
	jwtSignatureKey := []byte(appSecretKey)

	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    appName,
			ExpiresAt: time.Now().Add(loginExpDuration).Unix(),
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err := token.SignedString(jwtSignatureKey)
	if err != nil {
		return "", errors.New("GenerateTokenJwt err = " + err.Error())
	}

	return signedToken, nil
}
