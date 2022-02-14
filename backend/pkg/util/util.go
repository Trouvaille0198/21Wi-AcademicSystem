package util

import (
	"strconv"
)

func String2Int(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Int2String(i int) string {
	return strconv.Itoa(i)
}
