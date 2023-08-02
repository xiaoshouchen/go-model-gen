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
	TransType(filedType string, nullable string) string
	GetTableStructure(schema string, tables []string) []vars.Structure
}

func GetInstance(config vars.DatabaseConfig) I {
	switch config.Connect.Type {
	case "mysql":
		return NewMysql(config)
	case "postgres":
		return NewPostgres()
	}
	return nil
}

func (d *PublicFunc) getTableStructure(schema string, tables []string) []vars.Structure {
	var schemas []vars.Structure
	sql := `SELECT table_name,column_name,data_type,is_nullable FROM information_schema.columns WHERE table_name = ? and table_schema = ?;`

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
