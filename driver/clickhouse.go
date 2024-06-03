package driver

import (
	_ "embed"
	"fmt"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"log"
	"regexp"
	"strings"

	"github.com/xiaoshouchen/go-model-gen/vars"
)

type Clickhouse struct {
	PublicFunc
}

func NewClickhouse(config vars.DatabaseConfig) *Clickhouse {
	return &Clickhouse{PublicFunc{db: InitClickhouse(config)}}
}

func extractValue(input string) string {
	re := regexp.MustCompile(`Nullable\(([^)]+)\)`)
	match := re.FindStringSubmatch(input)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

// TransType 翻译数据类型，从clickhouse的格式到golang的格式
func (m *Clickhouse) TransType(filedType string, nullable string, name string) string {
	if strings.Contains(filedType, "Decimal") {
		filedType = "Decimal"
	}
	if strings.Contains(filedType, "FixedString") {
		filedType = "FixedString"
	}
	if v := extractValue(filedType); v != "" {
		filedType = v
	}
	switch filedType {
	case "UInt8", "UInt16", "UInt32", "UInt64", "Int8", "Int16", "Int32", "Int64":
		if nullable == "1" {
			return "sql.NullInt64"
		}
		return "int64"
	case "Float32", "Float64":
		if nullable == "1" {
			return "sql.NullFloat64"
		}
		return "float64"
	case "String", "FixedString":
		if nullable == "1" {
			return "sql.NullString"
		}
		return "string"
	case "Date", "DateTime", "timestamp":
		if nullable == "1" {
			return "sql.NullTime" // Assuming nullable date/time is represented as string
		}
		return "time.Time"
	case "Boolean":
		if nullable == "1" {
			return "sql.NullBool"
		}
		return "bool"
	default:
		return "interface{}"
	}
}

func (m *Clickhouse) HasSpecialType(fields []vars.Field) (hasNull bool, hasTime bool) {
	for _, v := range fields {
		if strings.Index(v.DataType, "Nullable") == 0 {
			v.IsNullable = "1"
		}
		res := m.TransType(v.DataType, v.IsNullable, v.ColumnName)
		if strings.Contains(res, "time") {
			hasTime = true
		} else if strings.Contains(res, "sql") {
			hasNull = true
		}
	}
	return
}

func (m *Clickhouse) GetTableStructure(schema string, tables []string) []vars.Structure {
	var schemas []vars.Structure
	sql := "SELECT column_name as column_name,data_type as data_type,is_nullable as is_nullable," +
		"column_comment as column_comment,column_default as column_default " +
		"FROM information_schema.columns WHERE table_name = ? and table_schema = ?;"
	for _, v := range tables {
		var res []vars.Field
		m.db.Raw(sql, v, schema).Scan(&res)
		schemas = append(schemas, vars.Structure{
			TableName: v,
			Fields:    res,
		})
	}
	for k, v := range schemas {
		hasNull, hasTime := m.HasSpecialType(v.Fields)
		v.HasNull = hasNull
		v.HasTime = hasTime
		schemas[k] = v
	}
	return schemas
}

func InitClickhouse(config vars.DatabaseConfig) *gorm.DB {
	con := config.Connect
	dnsStr := "clickhouse://%s:%s@%s:%s/%s?dial_timeout=10s&read_timeout=20s"
	dsn := fmt.Sprintf(dnsStr, con.User, con.Password, con.Host, con.Port, config.Scheme)
	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//go:embed tpl/clickhouse/model.tpl
var clickhouseModel string

//go:embed tpl/clickhouse/field.tpl
var clickhouseField string

//go:embed tpl/clickhouse/insert.tpl
var clickhouseInsert string

//go:embed tpl/clickhouse/omit.tpl
var clickhouseOmit string

func (m *Clickhouse) GetTpl() string {
	return clickhouseModel + clickhouseField + clickhouseInsert + clickhouseOmit
}
