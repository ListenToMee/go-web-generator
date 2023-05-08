{{define "router"}}package router

import (
	"github.com/gin-gonic/gin"
     "{{.ModuleName}}/controller"
)

// {{.Table.TableName}}Router 路由
func {{.Table.TableName}}Router(engin *gin.Engine) {

	// 定义路由组
	group := engin.Group("/article")

	// 绑定分页接口
	group.POST("/queryPage", {{.Table.TableName}}Controller.{{.Table.TableName}}Controller{}.QueryPage)

	// 绑定新增接口
    group.PUT("/add", {{.Table.TableName}}Controller.{{.Table.TableName}}Controller{}.Add)

    // 绑定更新接口
    group.PUT("/update", {{.Table.TableName}}Controller.{{.Table.TableName}}Controller{}.Update)

    // 绑定查询接口
    group.GET("/queryDetails", {{.Table.TableName}}Controller.{{.Table.TableName}}Controller{}.QueryDetails)

    // 绑定删除接口
    group.DELETE("/delete", {{.Table.TableName}}Controller.{{.Table.TableName}}Controller{}.Delete)

    // 绑定批量删除接口
    group.DELETE("/deleteBatch", {{.Table.TableName}}Controller.{{.Table.TableName}}Controller{}.DeleteBatch)
}{{end}}