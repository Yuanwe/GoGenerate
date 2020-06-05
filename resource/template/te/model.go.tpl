// 自动生成模板{{.StructName}}
package model


type {{.StructName}} struct {
      {{range .Fields}}
      {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}"{{if .ColumnName}} gorm:"column:{{.ColumnName}}"{{end}}`{{ end }}
}
//设置表名
func (v {{.StructName}}) TableName() string {
	return "{{.TableName}}"
}