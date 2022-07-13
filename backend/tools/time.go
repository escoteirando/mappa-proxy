package tools

import "time"

func DaysAgo(days int) time.Time {
	return time.Now().AddDate(0, 0, -days)
}