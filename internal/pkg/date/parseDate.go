package date

import "time"

func ParseDate(dateString string) (time.Time, error) {
	layout := "02/01/2006"
	return time.Parse(layout, dateString)
}
