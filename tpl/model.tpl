//go:gen DON'T EDIT !
package model

import (
{{if .HasNull}}"database/sql"{{end}}
{{if .HasTime}}"time"{{end}}

	"gorm.io/gorm"
    "gorm.io/gorm/clause"
)

{{ $modelName :=.TableName | singular | upCamel }}
type {{$modelName}} struct {
    {{range $k,$v := .Fields}}{{template "field" $v}}{{end}}
}

func ({{$modelName}}) TableName() string {
	return "{{.TableName}}"
}

type {{$modelName}}Repo struct{
    db *gorm.DB
}

func New{{$modelName}}Repo(db *gorm.DB)*{{$modelName}}Repo{
    return &{{$modelName}}Repo{db: db}
}

{{template "insert" .}}