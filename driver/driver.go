package driver

import (
	"github.com/xiaoshouchen/go-model-gen/vars"
	"gorm.io/gorm"
)

// PublicFunc 公共方法，可复用
type PublicFunc struct {
	db *gorm.DB
}

type I interface {
	TransType(filedType string, nullable string, name string) string
	GetTableStructure(schema string, tables []string) []vars.Structure
	GetTpl() string
}

func GetInstance(config vars.DatabaseConfig) I {
	switch config.Connect.Type {
	case "mysql":
		return NewMysql(config)
	case "postgres":
		return NewPostgres()
	case "clickhouse":
		return NewClickhouse(config)
	}
	return nil
}

func (d *PublicFunc) getTables(schema string) []string {
	var tables []string
	sql := "SELECT table_name FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA= ?"
	d.db.Raw(sql, schema).Pluck("table_name", &tables)
	return tables
}

func (d *PublicFunc) getTableStructure(schema string, tables []string) []vars.Structure {
	var schemas []vars.Structure
	sql := "SELECT column_name as column_name,data_type as data_type,is_nullable as is_nullable," +
		"column_comment as column_comment,column_key as column_key,column_default as column_default " +
		"FROM information_schema.columns WHERE table_name = ? and table_schema = ? order by ordinal_position;"
	for _, v := range tables {
		var res []vars.Field
		d.db.Raw(sql, v, schema).Scan(&res)
		schemas = append(schemas, vars.Structure{
			TableName: v,
			Fields:    res,
		})
	}
	return schemas
}
