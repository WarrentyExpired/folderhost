package utils

import (
	"strings"
)

func ReplacePathPrefix(fullPath string, realPrefix string) string {
	fullPath = strings.ReplaceAll(fullPath, "\\", "/")
	realPrefix = strings.ReplaceAll(realPrefix, "\\", "/")

	if strings.HasPrefix(fullPath, realPrefix) {
		realPrefixLen := len(realPrefix)

		if realPrefixLen >= len(fullPath) {
			return "./"
		}

		if string(fullPath[realPrefixLen]) == "/" {
			return "." + fullPath[realPrefixLen:]
		}

		return "./" + fullPath[realPrefixLen:]
	}

	return fullPath
}
