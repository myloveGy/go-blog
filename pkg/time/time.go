package time

import (
	"time"
)

var loc, _ = time.LoadLocation("Asia/Shanghai")

func Date() string {
	return time.Now().In(loc).Format("2006-01-02")
}

func DateTime() string {
	return time.Now().In(loc).Format("2006-01-02 15:04:05")
}

func FormatDate(t time.Time) string {
	return t.In(loc).Format("2006-01-02")
}

func FormatDateTime(t time.Time) string {
	return t.In(loc).Format("2006-01-02 15:04:05")
}
