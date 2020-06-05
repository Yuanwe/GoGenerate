package router

import (
	"{{.Project}}/api"
    "github.com/gin-gonic/gin"
)

func Init{{.StructName}}Router(Router *gin.RouterGroup) {
	{{.StructName}}Router := Router.Group("{{.Abbreviation}}")
	{
		{{.StructName}}Router.POST("{{.StructName}}", api.Create{{.StructName}})     // 新建{{.StructName}}
		{{.StructName}}Router.DELETE("{{.StructName}}", api.Delete{{.StructName}})   //删除{{.StructName}}
		{{.StructName}}Router.PUT("{{.StructName}}", api.Update{{.StructName}})   //更新{{.StructName}}
		{{.StructName}}Router.GET("{{.StructName}}", api.Find{{.StructName}})           // 根据ID获取{{.StructName}}
		{{.StructName}}Router.GET("{{.StructName}}List", api.Get{{.StructName}}List) //获取{{.StructName}}列表
}
}
