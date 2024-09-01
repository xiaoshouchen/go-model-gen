package vars

type Structure struct {
	TableName  string  `json:"table_name"`
	Fields     []Field `json:"fields"`
	HasNull    bool    `json:"has_null"`
	HasTime    bool    `json:"has_time"`
	SoftDelete int     `json:"soft_delete"`
}

type Field struct {
	ColumnName    string `json:"column_name"`
	DataType      string `json:"data_type"`
	IsNullable    string `json:"is_nullable"`
	ColumnComment string `json:"column_comment"`
	ColumnKey     string `json:"column_key"`
	ColumnDefault string `json:"column_default"`
}

type Result struct {
	TableName     string      `json:"table_name"`
	ColumnName    string      `json:"column_name"`
	DataType      string      `json:"data_type"`
	IsNullable    string      `json:"is_nullable"`
	ColumnComment string      `json:"column_comment"`
	ColumnKey     string      `json:"column_key"`
	ColumnDefault interface{} `json:"column_default"`
}

type DatabaseConfig struct {
	Scheme      string `json:"scheme"`
	PackageName string `json:"package_name"`
	Connect     struct {
		Type     string `json:"type"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"connect"`
	Tables            []string `json:"tables"`
	TableFilterOption string   `json:"table_filter_option"`
}
