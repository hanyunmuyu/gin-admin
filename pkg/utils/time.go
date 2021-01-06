package utils

import (
	"database/sql/driver"
	"time"
)

type Time time.Time

const (
	timeFormat = "2006-01-02 15:04:05"
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	if t.String() != "0001-01-01 00:00:00" {
		b = time.Time(t).AppendFormat(b, timeFormat)
	}
	b = append(b, '"')
	return b, nil
}
func (t *Time) Scan(v interface{}) error {
	*t = Time(v.(time.Time))
	return nil
}
func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}
func (t Time) Value() (driver.Value, error) {
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return time.Time(t), nil
}
