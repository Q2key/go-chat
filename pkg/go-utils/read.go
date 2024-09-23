package go_utils

import "os"

func ReadContent(path string) (string, error) {
	dat, err := os.ReadFile(path)
	Check(err)
	return string(dat), nil
}
