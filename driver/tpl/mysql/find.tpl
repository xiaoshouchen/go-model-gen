{{define "find"}}
    {{ $modelName := .TableName | singular | upCamel }}
    // 根据主键进行查询，最好有缓存
    func (r *{{$modelName}}Repo) FindByPk({{- range $index,$value:= .Fields}}{{- if ne $value.ColumnKey "PRI"}}{{- end}}{{- end}})(int64, error) {
    return 0, nil
    }
{{end}}