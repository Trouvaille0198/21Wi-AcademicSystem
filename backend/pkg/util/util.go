package util

import (
	"strconv"
)

func String2Int(s string) (int, error) {
	i, err := strconv.Atoi(s)
	return i, err
}

func Int2String(i int) string {
	return strconv.Itoa(i)
}
