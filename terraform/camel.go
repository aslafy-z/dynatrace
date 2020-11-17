package terraform

import (
	"strings"
	"unicode"
)

func unCamel(s string) string {
	s = strings.ReplaceAll(s, "ID", "Id")
	result := ""
	lastCharWasLower := false
	for _, c := range s {
		currentIsUpper := unicode.IsUpper(c)
		if lastCharWasLower && currentIsUpper {
			result = result + "_"
		}
		result = result + string(unicode.ToLower(c))
		lastCharWasLower = !currentIsUpper
	}
	return result
}
