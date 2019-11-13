package generatenumber

import (
	"crypto/rand"
	"errors"
	"io"
	"time"
)

//GenerateRandomCode - function for generate random number
func GenerateRandomCode(max int) (string, error) {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		return "", errors.New("GenerateRandomCode err = " + err.Error())
	}

	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b), nil
}

//GenerateNoTrans - generate transaction no
func GenerateNoTrans(myParam string) string {
	var myNo string
	myLength := 5
	myDate := time.Now()
	strDate := myDate.Format("20060102")
	myID, err := GenerateRandomCode(myLength)
	if err != nil {
		return ""
	}

	myNo = myParam + strDate + "-" + myID
	return myNo
}
