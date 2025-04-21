package tmpl

import (
	"unicode"
)

func CapitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func Variable(mapping map[string]string, target string) string {
	return mapping[target]
}
