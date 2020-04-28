package utils

import "strconv"

func ParseInt(s string) (n int, err error) {
	return strconv.Atoi(s)
}
