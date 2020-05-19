package v1

import (
	"fmt"
    "{{.Project}}/global/response"
    "{{.Project}}/model"
    "{{.Project}}/model/request"
    resp "{{.Project}}/model/response"
    "{{.Project}}/service"
    "github.com/gin-gonic/gin"
)


// @Tags {{.StructName}}
// @Summary 创建{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "创建{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /{{.Abbreviation}}/create{{.StructName}} [post]
func Create{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} model.{{.StructName}}
	_ = c.ShouldBindJSON(&{{.Abbreviation}})
	err := service.Create{{.StructName}}({{.Abbreviation}})
	if err != nil {
		response.Fail(c)
	} else {
		response.Ok(c)
	}
}


// @Tags {{.StructName}}
// @Summary 删除{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "删除{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /{{.Abbreviation}}/delete{{.StructName}} [delete]
func Delete{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} model.{{.StructName}}
	_ = c.ShouldBindJSON(&{{.Abbreviation}})
	err := service.Delete{{.StructName}}({{.Abbreviation}})
	if err != nil {
		response.Fail(c)
	} else {
		response.Ok(c)
	}
}


// @Tags {{.StructName}}
// @Summary 更新{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "更新{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /{{.Abbreviation}}/update{{.StructName}} [put]
func Update{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} model.{{.StructName}}
	_ = c.ShouldBindJSON(&{{.Abbreviation}})
	err := service.Update{{.StructName}}(&{{.Abbreviation}})
	if err != nil {
		response.Fail(c)
	} else {
		response.Ok(c)
	}
}


// @Tags {{.StructName}}
// @Summary 用id查询{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "用id查询{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /{{.Abbreviation}}/find{{.StructName}} [get]
func Find{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} model.{{.StructName}}
	_ = c.ShouldBindQuery(&{{.Abbreviation}})
	err,re{{.Abbreviation}} := service.Get{{.StructName}}({{.Abbreviation}}.ID)
	if err != nil {
		response.Fail(c)
	} else {
		response.OkWithData( gin.H{"re{{.Abbreviation}}":re{{.Abbreviation}},}, c)
	}
}


// @Tags {{.StructName}}
// @Summary 分页获取{{.StructName}}列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取{{.StructName}}列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /{{.Abbreviation}}/get{{.StructName}}List [get]
func Get{{.StructName}}List(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.Get{{.StructName}}InfoList(pageInfo)
	if err != nil {
		response.FailWithDetailed(err, resp.NOT_FOUND, "Bind Error", c)
	} else {
		totalPages := (total + pageInfo.Limit - 1) / pageInfo.Limit
        response.OkWithPageInfo(list, resp.PageResult{
        			Current:    pageInfo.Page,
        			Limit:      pageInfo.Limit,
        			TotalItems: total,
        			TotalPages: totalPages,
        		}, c)
	}
}