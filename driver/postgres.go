package driver

import (
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/xiaoshouchen/go-model-gen/vars"
)

type Postgres struct {
	PublicFunc
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (p *Postgres) TransType(filedType string, nullable string, name string) string {
	var dateType string
	dateType = strings.ToLower(dateType)
	switch filedType {
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

func (p *Postgres) GetTableStructure(schema string, tables []string) []vars.Structure {
	return p.getTableStructure(schema, tables)
}

func (*Postgres) Init(dbc vars.DatabaseConfig) []vars.Structure {
	dsn := os.ExpandEnv("user=$DB_USER password=$DB_PASSWORD host=$DB_HOST port=$DB_PORT dbname=$DB_DATABASE sslmode=$DB_SSL_MODE")
	_, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return nil
}

func (p *Postgres) GetTpl() string {
	return ""
}
