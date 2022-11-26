package tools

import (
	"strconv"
	"time"
)

func GetCurrentTime() string {
	return strconv.Itoa(int(
		time.Now().UnixNano() / int64(time.Millisecond),
	))
}
