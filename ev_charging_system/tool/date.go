package tool

import (
	"time"
)

// UnixToTimeString 时间戳转换"yyyy-MM-dd HH:mm:ss" 格式
func UnixToTimeString(unixTime int64) string {
	// 将时间戳转换为时间对象
	utcTime := time.Unix(unixTime, 0)

	// 指定时区为 UTC
	utcLocation := time.UTC
	utcTime = utcTime.In(utcLocation)

	// 格式化时间为字符串
	utcFormattedTime := utcTime.Format(time.DateTime)
	return utcFormattedTime
}

// GetNowTime 获取当前时间，并转换为"yyyy-MM-dd HH:mm:ss" 格式
func GetNowTime() string {
	// 获取当前时间
	currentTime := time.Now()

	// 指定时区为 UTC
	utcLocation := time.UTC
	currentTime = currentTime.In(utcLocation)

	// 格式化时间为字符串
	utcFormattedTime := currentTime.Format(time.DateTime)
	return utcFormattedTime
}
