{{define "omit"}}
    {{ $modelName := .TableName | singular | upCamel }}
    func (r *{{$modelName}}Repo) getAllFields() []string {
    return []string{
    {{- range $index,$value:= .Fields}}
        "{{$value.ColumnName}}",
    {{- end}}
    }
    }
    // Omit 过滤自己不想要的字段
    func (r *{{$modelName}}Repo) Omit(filter []string) []string {
    fields := r.getAllFields()
    result := make([]string, 0, len(fields))
    filterSet := make(map[string]bool)

    // 将需要过滤的值添加到 filterSet 中
    for _, v := range filter {
    filterSet[v] = true
    }

    // 遍历原始切片，将不在 filterSet 中的值添加到结果切片中
    for _, v := range fields {
    if _, ok := filterSet[v]; !ok {
    result = append(result, v)
    }
    }

    return result
    }
{{end}}