package getDate

import (
	"strings"
	"time"
)

func GetCurrentDate(format string) string {
	// Go's reference time: 2006-01-02 15:04:05
	const defaultFormat = "2006-01-02"

	// If user passes empty or invalid format, use default
	if format == "" || !strings.Contains(format, "2006") {
		format = defaultFormat
	}

	return time.Now().Format(format)
}

func FormatDate(t1, t2 time.Time) string {
	var d time.Time
	if !t1.IsZero() {
		d = t1
	} else {
		d = t2
	}
	if d.IsZero() {
		return ""
	}
	return d.Format("02 Jan,2006")
}
