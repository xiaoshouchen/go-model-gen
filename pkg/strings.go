package pkg

import (
	"strings"
)

func LineToCamel(str string) string {
	strSlice := strings.Split(strings.ToLower(str), "_")
	for k, s := range strSlice {
		strSlice[k] = strings.ToUpper(s[:1]) + s[1:]
	}
	return strings.Join(strSlice, "")
}
