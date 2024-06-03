{{define "insert"}}
{{ $modelName := .TableName | singular | upCamel }}
// Insert 插入，支持批量插入，单个插入，支持分批插入，支持冲突更新或者报错
// batchSize 如果数据过多，比如几万个，则需要拆分多个，等于小于0则不分批
func (r *{{$modelName}}Repo) Insert(batchSize int, insertSlice ...*{{$modelName}})(int64, error) {
	db := r.db
if batchSize > 0 {
db = db.CreateInBatches(insertSlice, batchSize)
} else {
db = db.Create(insertSlice)
}
return db.RowsAffected, db.Error
}

func (r *{{$modelName}}Repo) DB()*gorm.DB {
return r.db.Table("{{.TableName}}")
}
{{end}}