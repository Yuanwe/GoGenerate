package handler

import (
	"GoGenerate/global"
	"GoGenerate/model"
	"GoGenerate/service"
	"GoGenerate/utils"
	"encoding/json"
	"html/template"
	"os"
	"strings"
)

// @title    CreateTemp
// @description   函数的详细描述
// @auth                     （2020/04/05  20:22）
// @param     autoCode        model.AutoCodeStruct
// @return    err             error

func CreateTemp() (err error) {
	gen := global.GG_CONFIG.Template
	goTpls := gen.GoTpl
	tableList := strings.Split(gen.TableName, ",")
	for i := 0; i < len(tableList); i++ {
		tableName := tableList[i]
		_, autoCode := buildAutoCodeStruct(tableName)
		var data, _ = json.Marshal(autoCode)
		global.GG_LOG.Info("数据组织完毕,准备开始生成:" + string(data))

		for _, tpl := range goTpls {
			global.GG_LOG.Info("开始生成:" + tableName)
			modelTmpl, err := template.ParseFiles(tpl.TplPath)
			if err != nil {
				global.GG_LOG.Info("模板文件读取失败:"+tpl.TplPath, err)
				continue
			}
			err = utils.CreateMutiDir(tpl.Package)

			model, err := os.OpenFile(tpl.Package+tableName+convertFileName(tpl.TplPath), os.O_CREATE|os.O_WRONLY, 0755)
			if err != nil {
				global.GG_LOG.Info("GO文件打开失败:"+tpl.Package+tableName, err)
				continue
			}
			err = modelTmpl.Execute(model, autoCode)
			if err != nil {
				global.GG_LOG.Info("GO文件生成失败:"+tpl.Package+tableName, err)
				continue
			}
			_ = model.Close()
			global.GG_LOG.Info("成功创建:" + tableName)
		}
	}
	return nil
}

func buildAutoCodeStruct(tableName string) (err error, autoCode model.AutoCodeStruct) {
	_, datas := service.GetAllFields(tableName)

	var fields []model.Field

	for i := range datas {
		data := datas[i]
		field := model.Field{
			FieldName:  convertName(data.ColumnName),
			ColumnName: data.ColumnName,
			FieldType:  convertType(data.DataType),
			FieldJson:  convertJSON(data.ColumnName),
		}
		fields = append(fields, field)
	}
	autoCode.Project = global.GG_CONFIG.Template.Project
	autoCode.StructName = convertName(tableName)
	autoCode.Abbreviation = autoCode.StructName
	autoCode.Fields = fields
	return nil, autoCode
}

func convertType(fieldType string) string {
	var key string
	if fieldType == "" {
		return key
	}

	switch fieldType {
	case "varchar":
		key = "string"
	case "double":
		key = "float32"
	case "float":
		key = "float64"
	case "datetime":
		key = "time.Time"
	case "bigint":
		key = "int64"
	case "int":
		key = "uint"
	default:
		key = fieldType
	}

	return key
}
func convertName(fieldName string) string {
	temp := strings.Split(fieldName, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				upperStr += strings.ToUpper(string(vv[i]))
			} else {
				upperStr += string(vv[i])
			}
		}

	}
	return upperStr
}
func convertJSON(fieldName string) string {
	temp := strings.Split(fieldName, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		if y != 0 {
			for i := 0; i < len(vv); i++ {
				if i == 0 {
					vv[i] -= 32
					upperStr += string(vv[i]) // + string(vv[i+1])
				} else {
					upperStr += string(vv[i])
				}
			}
		}
	}
	return temp[0] + upperStr
}

func convertFileName(tpl string) string {
	temp := strings.Split(tpl, ".")
	return "." + temp[len(temp)-2]
}
