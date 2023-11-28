{{- define "field"}}
    {{.ColumnName | upCamel}} {{transType .DataType .IsNullable  }} `json:"{{.ColumnName}}" {{template "gorm" .}}` //{{.ColumnComment|inline}}
{{- end}}

{{- define "gorm"}}
    {{- if  or (containsNumber .ColumnName)  (eq .ColumnKey "PRI")}}gorm:"{{if containsNumber .ColumnName}}column:{{.ColumnName}}{{end}}{{if eq .ColumnKey "PRI"}}primaryKey{{end}}"{{- end}}
{{- end}}