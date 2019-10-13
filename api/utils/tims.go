package utils

import "time"

// 毫秒
func ToMilliSecond(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// 当前时间的毫秒值
func NowMilliSecond() int64 {
	return ToMilliSecond(time.Now())
}

//ToTime 毫秒转时间
func ToTime(msec int64) time.Time {
	if msec == 0 {
		return time.Unix(0, 0)
	}
	t := msec / 1000
	return time.Unix(t, msec*1e6%t)
}

func ToTimeString(msec int64) string {
	return ToDatetimeString(ToTime(msec))
}
func ToDatetimeString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
