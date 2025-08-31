package util

import (
	"time"
)

func IsValidDate(strDate string) (time.Time, bool) {
	t, err := time.Parse("02/01/2006", strDate)
	if err == nil {
		return t, true
	}
	return t, false
}
