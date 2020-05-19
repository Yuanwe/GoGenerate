package global

import (
	"GoGenerate/config"
	"github.com/jinzhu/gorm"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	GG_DB     *gorm.DB
	GG_LOG    *oplogging.Logger
	GG_CONFIG config.Server
	GG_VP     *viper.Viper
)
