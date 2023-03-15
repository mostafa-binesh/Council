package utils

import "strings"

func ToCamelCase(s string) string {
	var result string
	capitalize := false
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			result += string(c)
		} else if c >= 'a' && c <= 'z' {
			if capitalize {
				result += strings.ToUpper(string(c))
				capitalize = false
			} else {
				result += string(c)
			}
		} else {
			capitalize = true
		}
	}
	return result
}
