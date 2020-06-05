package model

// 初始版本自动化代码工具
type AutoCodeStruct struct {
	StructName   string  `json:"structName"`
	Project      string  `json:"project"`
	Abbreviation string  `json:"abbreviation"`
	TableName    string  `json:"TableName"`
	Fields       []Field `json:"fields"`
}

type Field struct {
	FieldName  string `json:"fieldName"`
	FieldType  string `json:"fieldType"`
	FieldJson  string `json:"fieldJson"`
	ColumnName string `json:"columnName"`
	FieldDesc  string `json:"fieldDesc"`
}

type TableField struct {
	ColumnName string `gorm:"column:COLUMN_NAME"`
	DataType   string `gorm:"column:DATA_TYPE"`
}

type GoTpl struct {
	Package string `json:"package"`
	TplPath string `json:"tplPath"`
}
