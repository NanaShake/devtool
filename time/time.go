package lib

import "time"

// 转换字符串格式时间 2006-01-02 15:04:05 为时间戳
func ParseLocalUnixTimeStamp(val string) (int64, error) {
	tm, err := time.ParseInLocation("2006-01-02 15:04:05", val, time.Local)
	if nil != err {
		return 0, err
	}
	return tm.UTC().Unix(), nil
}

// 转换字符串格式时间 2006-01-02 15:04:05 为 time.Time
func ParseLocalUnixTime(val string) (time.Time, error) {
	tm, err := time.ParseInLocation("2006-01-02 15:04:05", val, time.Local)
	if nil != err {
		return time.Time{}, err
	}
	return tm, nil
}

// 返回  20060102
func GetDate(tm *time.Time) int {
	return int(tm.Year())*10000 + int(tm.Month())*100 + tm.Day()
}
