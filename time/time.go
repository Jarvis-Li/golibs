package time

import (
	"time"
)

func IsOut(ts int64, limit time.Duration) bool {
	then := time.Unix(ts, 0)
	now := time.Now()
	return now.Sub(then) > limit
}

func TruncateToDay(t time.Time) time.Time {
	loc := t.Location()
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, loc)
}

func Remain(start int64, limit time.Duration) (remain int64) {
	end := time.Unix(start, 0).Add(limit).Unix()
	now := time.Now().Unix()
	if end < now {
		return
	}
	return end - now
}

func Today() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// DruationToday() 返回现在相对于今天0点经过的时间
func DurationToday() time.Duration {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return now.Sub(today)
}

func DateStrCn(t time.Time) string {
	return t.Format("2006年1月2日")
}

func Str(t time.Time) string {
	return t.Format("01-02 15:04")
}
