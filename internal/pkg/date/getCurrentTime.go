package date

import "time"

func GetCurrentTime() (time.Time, error) {
	str := time.Now().Format("2006-01-02 15:04:05")
	t, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
