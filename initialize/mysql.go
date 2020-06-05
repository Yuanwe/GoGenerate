package initialize

import (
	"GoGenerate/global"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

//初始化数据库并产生数据库全局变量
func Mysql() {
	admin := global.GG_CONFIG.Mysql
	if db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config); err != nil {
		global.GG_LOG.Error("MySQL启动异常", err)
		os.Exit(0)
	} else {
		global.GG_DB = db
		global.GG_DB.DB().SetMaxIdleConns(admin.MaxIdleConns)
		global.GG_DB.DB().SetMaxOpenConns(admin.MaxOpenConns)
		global.GG_DB.LogMode(admin.LogMode)
	}
}
