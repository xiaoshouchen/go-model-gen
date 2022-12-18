package vars

type Schema struct {
	TableName string  `json:"table_name"`
	Fields    []Field `json:"fields"`
}

type Field struct {
	ColumnName string `json:"column_name"`
	DataType   string `json:"data_type"`
}

type Result struct {
	TableName  string `json:"table_name"`
	ColumnName string `json:"column_name"`
	DataType   string `json:"data_type"`
}
