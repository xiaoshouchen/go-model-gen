{{- define "field"}}
    {{.ColumnName | upCamel}} {{transType .DataType .IsNullable  }} `json:"{{.ColumnName}}" {{if eq .ColumnKey "PRI"}}gorm:"primaryKey"{{end}}` //{{.ColumnComment|inline}}
{{- end}}