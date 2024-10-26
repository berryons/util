package util

import (
	"runtime"
)

/**
 * @author Ted Yoo
 * @breif OS util
 * @date 2024.10.26
 */

type ERuntimeOS int

const (
	UNKNOWN ERuntimeOS = iota
	WINDOWS
	MACOS
	LINUX
)

func GetRuntimeOS() ERuntimeOS {
	os := runtime.GOOS
	switch os {
	case "windows":
		return WINDOWS
	case "darwin":
		return MACOS
	case "linux":
		return LINUX
	default:
		return UNKNOWN
	}
}
