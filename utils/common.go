package utils

import "strconv"

func StringToInt(value string, def int64) int64 {
	r, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		r = def
	}
	return r
}
