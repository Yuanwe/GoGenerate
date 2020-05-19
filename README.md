
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
## 2. Getting started
```
- golang version >= v1.11
- IDE recommendation: Goland
```
## 3. Technical selection

- DB: `MySql`，using `gorm` to implement data manipulation, added support for SQLite databases.
- Config: using `fsnotify` and `viper` to implement `yaml` config file。
- Log: using `go-logging` record logs。

## 4. Project Architecture
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
## 5. Deploy





