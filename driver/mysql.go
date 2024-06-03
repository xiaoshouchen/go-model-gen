package driver

import (
	_ "embed"
	"fmt"
	"log"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/xiaoshouchen/go-model-gen/vars"
)

type Mysql struct {
	PublicFunc
}

func NewMysql(config vars.DatabaseConfig) *Mysql {
	return &Mysql{PublicFunc{db: InitMysql(config)}}
}

// TransType 翻译数据类型，从mysql的格式到golang的格式
func (m *Mysql) TransType(filedType string, nullable string, name string) string {
	var dateType string
	dateType = strings.ToLower(dateType)
	switch filedType {
	case "smallint", "integer", "int", "bigint", "serial", "bigserial", "smallserial", "tinyint", "mediumint":
		dateType = "int64"
	case "decimal", "numeric", "real", "double precision", "money", "float", "double":
		if strings.ToLower(nullable) == "yes" {
			dateType = "sql.NullFloat64"
		} else {
			dateType = "float64"
		}
	case "text", "varchar", "character varying", "character", "char", "mediumtext":
		if strings.ToLower(nullable) == "yes" {
			dateType = "sql.NullString"
		} else {
			dateType = "string"
		}
	case "boolean":
		if strings.ToLower(nullable) == "yes" {
			dateType = "sql.NullBool"
		} else {
			dateType = "bool"
		}
	case "date", "timestamp", "datetime":
		if strings.ToLower(nullable) == "yes" {
			dateType = "sql.NullTime"
		} else {
			dateType = "time.Time"
		}
	default:
		dateType = "interface{}"
	}
	if strings.Contains(name, "deleted_at") {
		dateType = "soft_delete.DeletedAt"
	}
	return dateType
}

func (m *Mysql) HasSpecialType(fields []vars.Field) (hasNull bool, hasTime bool, hasSoftDelete bool) {
	for _, v := range fields {
		res := m.TransType(v.DataType, v.IsNullable, v.ColumnName)
		if strings.Contains(res, "time") {
			hasTime = true
		} else if strings.Contains(res, "sql") {
			hasNull = true
		}
		if strings.Contains(res, "soft_delete") {
			hasSoftDelete = true
		}
	}
	return
}

func (m *Mysql) GetTableStructure(schema string, tables []string) []vars.Structure {
	res := m.getTableStructure(schema, tables)
	for k, v := range res {
		hasNull, hasTime, hasSoftDelete := m.HasSpecialType(v.Fields)
		v.HasNull = hasNull
		v.HasTime = hasTime
		v.SoftDelete = hasSoftDelete
		res[k] = v
	}
	return res
}

func InitMysql(config vars.DatabaseConfig) *gorm.DB {
	con := config.Connect
	dnsStr := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(dnsStr, con.User, con.Password, con.Host, con.Port, config.Scheme)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//go:embed tpl/mysql/model.tpl
var mysqlModel string

//go:embed tpl/mysql/field.tpl
var mysqlField string

//go:embed tpl/mysql/insert.tpl
var mysqlInsert string

//go:embed tpl/mysql/omit.tpl
var mysqlOmit string

func (m *Mysql) GetTpl() string {
	return mysqlModel + mysqlField + mysqlInsert + mysqlOmit
}
