package common

import (
	"time"
)

func NowUTC() time.Time {
	return time.Now()
}

func NowJST() {

}

func Time2String(t time.Time) string {
	return ""
}

func String2Time(ts string) time.Time {
	jst := "JST"
	t, _ := time.Parse("2006-01-02 15:04:05 MST", ts+" "+jst)
	return t
}
