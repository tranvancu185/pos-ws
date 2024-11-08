package utime

import "time"

var (
	VNLocation, _ = time.LoadLocation("Asia/Ho_Chi_Minh")
)

const (
	YYYY_MM_DD_HH_MM_SS = "YYYY_MM_DD_HH_MM_SS"
	YYYY_MM_DD          = "YYYY_MM_DD"
	HH_MM_SS            = "HH_MM_SS"
	YYYYMMDD            = "YYYYMMDD"
	YYYYMMDDHHMMSS      = "YYYYMMDDHHMMSS"
	YYMMDD              = "YYMMDD"
)

var DATE_TIME_FORMAT_VALUES = map[string]string{
	YYYY_MM_DD_HH_MM_SS: "2006-01-02 15:04:05",
	YYYY_MM_DD:          "2006-01-02",
	HH_MM_SS:            "15:04:05",
	YYYYMMDD:            "20060102",
	YYYYMMDDHHMMSS:      "20060102150405",
	YYMMDD:              "060102",
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func GetCurrentTimeString(format string) string {
	// Get current time
	var result string
	currenTime := GetCurrentTime()

	switch format {
	case YYYY_MM_DD_HH_MM_SS:
		result = currenTime.Format(DATE_TIME_FORMAT_VALUES[YYYY_MM_DD_HH_MM_SS])
	case YYYY_MM_DD:
		result = currenTime.Format(DATE_TIME_FORMAT_VALUES[YYYY_MM_DD])
	case HH_MM_SS:
		result = currenTime.Format(DATE_TIME_FORMAT_VALUES[HH_MM_SS])
	case YYYYMMDD:
		result = currenTime.Format(DATE_TIME_FORMAT_VALUES[YYYYMMDD])
	case YYYYMMDDHHMMSS:
		result = currenTime.Format(DATE_TIME_FORMAT_VALUES[YYYYMMDDHHMMSS])
	case YYMMDD:
		result = currenTime.Format(DATE_TIME_FORMAT_VALUES[YYMMDD])
	default:
		result = currenTime.Format(DATE_TIME_FORMAT_VALUES[YYYY_MM_DD_HH_MM_SS])
	}
	return result
}

func GetCurrentTimeUnix() int64 {
	return GetCurrentTime().Unix()
}

func StartOfCurrentDay() time.Time {
	return StartOfDay(GetCurrentTime())
}

func EndOfCurrentDay() time.Time {
	return EndOfDay(GetCurrentTime())
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

func StartOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}
