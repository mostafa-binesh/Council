package utils

import (
	"strings"
	// "unicode"
)

// ToSnake convert the given string to snake case following the Golang format:
// acronyms are converted to lower-case and preceded by an underscore.
// func ToSnake(in string) string {
// 	runes := []rune(in)
// 	length := len(runes)

// 	var out []rune
// 	for i := 0; i < length; i++ {
// 		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
// 			out = append(out, '_')
// 		}
// 		out = append(out, unicode.ToLower(runes[i]))
// 	}

// 	return string(out)
// }

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
func ToSnake(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "_")
	s = strings.ReplaceAll(s, "-", "_")
	return s
}