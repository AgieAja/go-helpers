package validation

import (
	"encoding/base64"
	"regexp"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"os"
)

//CheckPasswordHash - validate encrypt password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//HashPassword - Encrypt password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//IsEmail - validate format email,return true if correct email
func IsEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	res := re.MatchString(email)
	return res
}

//IsNumeric - validation for input field is numeric
//return true for integer,return false for not integer
func IsNumeric(strNoHp string) bool {
	_, err := strconv.Atoi(strNoHp)
	if err != nil {
		return false
	}
	return true
}

//SignatureSocialLogin - verify social login
func SignatureSocialLogin(email, name, sig string) bool {
	clientSecret := os.Getenv("CLIENT_SECRET")
	myToken := email + name + clientSecret
	mySignature := base64.StdEncoding.EncodeToString([]byte(myToken))
	if mySignature != sig {
		return false
	}

	return true
}
