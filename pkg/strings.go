package pkg

import (
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
