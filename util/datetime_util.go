package util

import (
	"strconv"
	"time"
)

func ConvertDateTime(dateString string) (time.Time, error) {
	format := "2006-01-02 15:04:05"
	location, err := time.LoadLocation("UTC")
	if err != nil {
		return time.Now(), err
	}

	dateTime, err := time.ParseInLocation(format, dateString, location)
	if err != nil {
		return dateTime, err
	}

	return dateTime, nil
}

func StringToEpoch(epochString string) (time.Time, error) {
	epochInt, err := strconv.ParseInt(epochString, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	epochTime := time.Unix(epochInt, 0)
	utcEpochTime := epochTime.UTC()

	return utcEpochTime, nil
}

func DateToEpoch(inputTime time.Time) int64 {
	return inputTime.Unix()
}

func NumberToEpoch(epochInt int64) (time.Time, error) {
	epochTime := time.Unix(epochInt, 0)
	utcEpochTime := epochTime.UTC()

	return utcEpochTime, nil
}

func TimeToStringDateOnly(t time.Time) string {
	return t.Format("2006-01-02")
}
