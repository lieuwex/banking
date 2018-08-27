package utils

import "time"

func Date(d time.Time) time.Time {
	year, month, day := d.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func Today() time.Time {
	return Date(time.Now())
}
