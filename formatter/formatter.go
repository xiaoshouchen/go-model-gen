package formatter

import (
	_ "embed"
	"fmt"
	"github.com/gertd/go-pluralize"
	"github.com/xiaoshouchen/go-model-gen/pkg"
	"github.com/xiaoshouchen/go-model-gen/vars"
	"strings"
)

//go:embed tpl/filed.tpl
var field string

//go:embed tpl/model.tpl
var model string

//go:embed tpl/repo.tpl
var repo string

//go:embed tpl/table_name.tpl
var tableName string

type Formatter struct {
	pluralize *pluralize.Client
}

func NewFormatter() *Formatter {
	return &Formatter{
		pluralize: pluralize.NewClient(),
	}
}

func (formatter *Formatter) Generate(table vars.Structure) string {
	modelStr := formatter.generateModel(table.TableName, formatter.joinFields(table.Fields))
	repoStr := formatter.generateRepo(table.TableName)
	nameStr := formatter.generateTableName(table.TableName)
	return strings.Join([]string{modelStr, repoStr, nameStr}, "\n")
}

func (formatter *Formatter) generateField(f vars.Field) string {
	name := pkg.LineToCamel(f.ColumnName)
	jsonName := strings.ToLower(f.ColumnName)
	return fmt.Sprintf(field, name, f.DataType, jsonName)
}

func (formatter *Formatter) joinFields(fields []vars.Field) string {
	var s string
	for _, v := range fields {
		s += formatter.generateField(v) + "\n"
	}
	return s
}

func (formatter *Formatter) generateModel(table, field string) string {
	tableName := formatter.pluralize.Singular(pkg.LineToCamel(table))
	return fmt.Sprintf(model, tableName, field)
}

func (formatter *Formatter) generateRepo(table string) string {
	tableStr := formatter.pluralize.Singular(pkg.LineToCamel(table))
	return fmt.Sprintf(repo, tableStr, tableStr, tableStr, tableStr)
}

func (formatter *Formatter) generateTableName(table string) string {
	tableStr := formatter.pluralize.Singular(pkg.LineToCamel(table))
	return fmt.Sprintf(tableName, tableStr, table)
}
