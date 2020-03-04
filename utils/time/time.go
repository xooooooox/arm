package time

import "time"

func DateTime() string {
	return time.Now().Format("2006-01-02 13:05:06")
}
