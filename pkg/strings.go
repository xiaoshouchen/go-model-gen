package pkg

import (
	"regexp"
	"strings"
)

func LineToLowCamel(str string) string {
	strSlice := strings.Split(strings.ToLower(str), "_")
	for k, s := range strSlice {
		if k == 0 {
			continue
		}
		strSlice[k] = strings.ToUpper(s[:1]) + s[1:]
	}
	return strings.Join(strSlice, "")
}

func LineToUpCamel(str string) string {
	strSlice := strings.Split(strings.ToLower(str), "_")
	for k, s := range strSlice {
		strSlice[k] = strings.ToUpper(s[:1]) + s[1:]
	}
	return strings.Join(strSlice, "")
}

// Inline 备注变成一行
func Inline(str string) string {
	str = strings.Replace(str, "\n", " ", -1)
	str = strings.Replace(str, "\t", " ", -1)
	return str
}

// ContainsNumber 判断字符串是否包含数字
func ContainsNumber(s string) bool {
	pattern := regexp.MustCompile(`\d`)
	return pattern.MatchString(s)
}
