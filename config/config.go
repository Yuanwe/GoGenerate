package config

import "GoGenerate/model"

type Server struct {
	Mysql    Mysql    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Log      Log      `mapstructure:"log" json:"log" yaml:"log"`
	Template Template `mapstructure:"template" json:"template" yaml:"template"`
}
type Mysql struct {
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
}

type Log struct {
	Prefix  string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	LogFile bool   `mapstructure:"log-file" json:"logFile" yaml:"log-file"`
	Stdout  string `mapstructure:"stdout" json:"stdout" yaml:"stdout"`
	File    string `mapstructure:"file" json:"file" yaml:"file"`
}

type Template struct {
	Project   string        `mapstructure:"project" json:"project" yaml:"project"`
	TableName string        `mapstructure:"TableName" json:"TableName" yaml:"TableName"`
	GoTpl     []model.GoTpl `mapstructure:"gotpl" json:"gotpl" yaml:"gotpl"`
}
