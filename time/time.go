package time

import (
	"strconv"
	t "time"
)

func TickNow() string {
	return strconv.FormatInt(t.Now().UnixNano(), 10)
}

func TickInt64() int64 {
	return t.Now().UnixNano()
}

func String(tick int64) string {
	return strconv.FormatInt(tick, 10)
}

func Int64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return i, nil
}
