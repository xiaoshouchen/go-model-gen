{{define "insert"}}
{{ $modelName := .TableName | singular | upCamel }}
// Insert 插入，支持批量插入，单个插入，支持分批插入，支持冲突更新或者报错
// allowUpdate 允许报错更新
// batchSize 如果数据过多，比如几万个，则需要拆分多个，等于小于0则不分批
func (r *{{$modelName}}Repo) Insert(allowUpdate bool, batchSize int, insertSlice ...*{{$modelName}}) {
	db := r.db
	if allowUpdate {
		db = db.Clauses(
			clause.OnConflict{
				Columns: []clause.Column{
{{- range $index,$value:= .Fields}}
	{{- if eq $value.ColumnKey "PRI"}}
					{Name: "{{$value.ColumnName}}"},
	{{- end}}
{{- end}}
				},
				DoUpdates: clause.Assignments(map[string]interface{}{
{{- range $index,$value:= .Fields}}
	{{- if ne $value.ColumnKey "PRI"}}
		"{{$value.ColumnName}}": db.Raw("values({{$value.ColumnName}})"),
	{{- end}}
{{- end}}
				}),
			},
		)
	}
	if batchSize > 0 {
		db.CreateInBatches(insertSlice, batchSize)
	} else {
		db.Create(insertSlice)
	}
}
{{end}}