package main

import (
	_ "GoGenerate/core"
	"GoGenerate/global"
	"GoGenerate/handler"
	"GoGenerate/initialize"
)

func main() {
	initialize.Mysql()
	gen := global.GG_CONFIG.Template
	global.GG_LOG.Info("开始自动生成代码,tableName:" + gen.TableName)
	handler.CreateTemp()
	global.GG_LOG.Info("自动生成代码结束,tableName:" + gen.TableName)
	// 程序结束前关闭数据库链接
	defer global.GG_DB.Close()
}
