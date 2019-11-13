package date

import (
	"errors"
	"time"
)

//ConvertStrToDate - convert string to date
func ConvertStrToDate(myStr string) (time.Time, error) {
	myDate, err := time.Parse("2006-01-02", myStr)
	if err != nil {
		return myDate, errors.New("ConvertStrToDate err = " + err.Error())
	}
	return myDate, nil
}

//GetDate - get only date
func GetDate() (time.Time, error) {
	t := time.Now()
	dateStr := t.Format("2006-01-02")
	date, err := ConvertStrToDate(dateStr)
	if err != nil {
		return date, err
	}

	return date, nil
}

//GetDateUnix - generate date format unixtime
func GetDateUnix() int64 {
	t := time.Now()
	dateUnix := t.Unix() * 1000
	return dateUnix
}
