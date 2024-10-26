package util

import "os"

/**
 * @author Ted Yoo
 * @breif Directory util
 * @date 2024.10.26
 */

func ExistsDir(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateDir(path string, mode os.FileMode) error {
	return os.MkdirAll(path, mode)
}

func CreateAllDir(path string, mode os.FileMode) error {
	return os.MkdirAll(path, mode)
}

func DeleteDir(path string) error {
	return os.RemoveAll(path)
}
