package utils

import "time"

var (
	VNLocation, _ = time.LoadLocation("Asia/Ho_Chi_Minh")
)

func GetCurrentTimeString(format string) string {
	// Get current time
	var result string
	switch format {
	case "yyyy-mm-dd hh:mm:ss":
		result = time.Now().Format("2006-01-02 15:04:05")
	case "yyyy-mm-dd":
		result = time.Now().Format("2006-01-02")
	case "hh:mm:ss":
		result = time.Now().Format("15:04:05")
	default:
		result = time.Now().Format("2006-01-02 15:04:05")
	}
	return result
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func GetCurrentTimeUnix() int64 {
	return time.Now().Unix()
}

func ConvertTimeToString(time time.Time, format string) string {
	return time.Format(format)
}

func ConvertStringToTime(timeString string, format string) (time.Time, error) {
	return time.Parse(format, timeString)
}

func ConvertUnixTimeToString(unixTime int64, format string) string {
	return time.Unix(unixTime, 0).Format(format)
}

func ConvertStringToUnixTime(timeString string, format string) (int64, error) {
	t, err := time.Parse(format, timeString)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

func BeginningOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}
