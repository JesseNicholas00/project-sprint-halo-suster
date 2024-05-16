package helper

import "time"

func MustParseDateOnly(timeString string) time.Time {
	res, err := time.Parse(time.DateOnly, timeString)
	if err != nil {
		panic(err)
	}
	return res
}
