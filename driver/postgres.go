package driver

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/xiaoshouchen/go-model-gen/vars"
)

func InitPostgres(tables []string) []vars.Schema {
	dsn := os.ExpandEnv("user=$DB_USER password=$DB_PASSWORD host=$DB_HOST port=$DB_PORT dbname=$DB_DATABASE sslmode=$DB_SSL_MODE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var schemas []vars.Schema
	sql := `SELECT table_name,column_name,data_type FROM information_schema.columns WHERE table_name = '%s';`

	for _, v := range tables {
		var res []vars.Field
		db.Raw(fmt.Sprintf(sql, v)).Scan(&res)

		schemas = append(schemas, vars.Schema{
			TableName: v,
			Fields:    res,
		})
	}

	return schemas

}
