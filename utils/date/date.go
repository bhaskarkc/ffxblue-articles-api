package date

import "time"

// https://golang.org/pkg/time/#pkg-constants
const (
	apiDateLayout = "2006-01-02T15:04:05.000Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}

// https://golang.org/pkg/time/#Parse
func Validate(date string) bool {
	_, err := time.Parse(apiDateLayout, date)
	if err != nil {
		return false
	}
	return true
}
