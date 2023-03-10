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

// TransType translate database types
func TransType(dbType string) string {
	var dateType string
	dateType = strings.ToLower(dateType)
	switch dbType {
	case "smallint", "integer", "int", "bigint", "serial", "bigserial", "smallserial", "tinyint", "mediumint":
		dateType = "int64"
	case "decimal", "numeric", "real", "double precision", "money", "float", "double":
		dateType = "float64"
	case "text", "varchar", "character varying", "character", "char":
		dateType = "string"
	case "boolean":
		dateType = "bool"
	default:
		dateType = "interface{}"
	}
	return dateType
}
