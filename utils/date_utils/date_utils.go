package date_utils

import "time"

const (
	apiDateLayout = "02-01-2006T15:04:05Z"
)

func GetNowString() string {
	return time.Now().UTC().Format(apiDateLayout)
}
