package main

import "time"

func date(d time.Time) time.Time {
	year, month, day := d.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func today() time.Time {
	return date(time.Now())
}
