package utils

import "time"

// AddMonth 指定した日付に月を追加する
func AddMonth(t time.Time, dMonth int) time.Time {
	year := t.Year()
	month := t.Month()
	day := t.Day()
	newMonth := int(month) - dMonth
	newLastDay := getLastDay(year, newMonth)
	var newDay int
	if day > newLastDay {
		newDay = newLastDay
	} else {
		newDay = day
	}

	return time.Date(year, time.Month(newMonth), newDay, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// その月の最終日を求める
func getLastDay(year, month int) int {
	t := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.Local)
	t = t.AddDate(0, 0, -1)
	return t.Day()
}
