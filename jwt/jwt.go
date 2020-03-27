package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var (
	applicationName  = os.Getenv("APP_NAME")
	jwtSigningMethod = jwt.SigningMethodHS256
	jwtSignatureKey  = []byte(os.Getenv("APP_SECRET_KEY_JWT"))
)

//MyClaims - struct jwt
type MyClaims struct {
	jwt.StandardClaims
	UserID uint `json:"user_id"`
}

//GenerateTokenJwt - generate token jwt
func GenerateTokenJwt(userID uint, loginExpDuration time.Duration) (string, error) {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    applicationName,
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
