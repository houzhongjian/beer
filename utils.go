package beer

import "strconv"

func parstInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}

//dir .
func dir(dir string) string {
	return dir
}