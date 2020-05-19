
<div align=center>
<h1>Go Generate</h1>
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.14-blue"/>
<img src="https://img.shields.io/badge/gorm-1.9.12-red"/>
<img src="https://img.shields.io/badge/viper-1.7.0-green"/>
</div>

# 项目文档
## 1. 基本介绍
### 1.1 项目介绍
> Go Generate是一个基于Go"html/template" 的代码自动生成项目，只需要简单的配置即可根据模板自动生成CURD
> 模板文件在/resource目录下
> 配置文件 gen_config.yaml  
## 2. 使用说明
```
- golang version >= v1.11
- IDE recommendation: Goland
```
## 3. 技术选型

- 数据库：，使用`gorm`对数据库的基本操作。
- 配置文件：使用`fsnotify`和`viper`实现`yaml`格式的配置文件。
- 日志：使用`go-logging`实现日志记录。

## 4. 目录结构
```
    ├─server  	     （backend）
      ├─config         （config file）
      ├─core  	        （core code）
      ├─global         （global objet）
      ├─handler         （handler）
      ├─initialiaze    （initialiazation）
      ├─model          （model and services）
      ├─service         (services)
      └─utils	        （common utilities）
```
## 5. 配置说明
- project：项目名、（项目目录）导包时需要用到
- tableName：需要生成CURD的数据库表名，支持多个表同时生成，用英文逗号分隔
- gotpl.package:代码存放路径，可以直接定义到对应的文件夹下
- gotpl.tplpath:文件模板对应目录，目前仅支持外挂目录（修改灵活）
- 文件名会根据表名自动生成，格式为tableName.type
```yaml
 project: "GoGenerate"
  tableName: "t_campaigns,t_delearship_enrolls"
  gotpl:
    - {package: autoCode/api/,tplpath: ./resource/template/te/api.go.tpl}
    - {package: autoCode/model/,tplpath: ./resource/template/te/model.go.tpl}
    - {package: autoCode/router/,tplpath: ./resource/template/te/router.go.tpl}
    - {package: autoCode/service/,tplpath: ./resource/template/te/service.go.tpl}
    - {package: autoCode/fe/,tplpath: ./resource/template/fe/api.js.tpl}
    - {package: autoCode/fe/,tplpath: ./resource/template/fe/table.vue.tpl}
```




