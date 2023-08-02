package pkg

import "strings"

type Mysql struct {
}

func NewMysql() *Mysql {
	return &Mysql{}
}

// TransType translate database types
func (mysql *Mysql) TransType(dbType string, nullable string) string {
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
