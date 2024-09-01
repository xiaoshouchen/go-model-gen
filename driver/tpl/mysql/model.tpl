//go:gen DON'T EDIT !
package model

import (
{{if .HasNull}}"database/sql"{{end}}
{{if .HasTime}}"time"{{end}}

	"gorm.io/gorm"
    "gorm.io/gorm/clause"
{{if eq .SoftDelete 1}}"gorm.io/plugin/soft_delete"{{end}}
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

func (r *{{$modelName}}Repo) DB()*gorm.DB{
return r.db
}

{{template "insert" .}}
{{template "omit" .}}