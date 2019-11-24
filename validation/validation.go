package validation

import (
	"regexp"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"google.golang.org/api/oauth2/v2"
	"net/http"
)

//CheckPasswordHash - validate encrypt password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//HashPassword - Encryp password
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

//VerifyIDTokenGoogle - verify id token google
func VerifyIDTokenGoogle(idToken string) (*oauth2.Tokeninfo, error) {
	httpClient := &http.Client{}
	oauth2Service, err := oauth2.New(httpClient)
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, err
	}

	return tokenInfo, nil
}
