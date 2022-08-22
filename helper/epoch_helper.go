package helper

import "time"

// IsYesterday Check if epochTime is yesterday
func IsYesterday(epochTime int64) bool {
	oneDay := int64(6400)
	inDayMode := epochTime - time.Now().Unix()
	if inDayMode > oneDay {
		return true
	}
	return false
}
