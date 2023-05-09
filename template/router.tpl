{{define "router"}}package router

import (
	"github.com/gin-gonic/gin"
    "{{.ModuleName}}/controller"
)

// {{.Table.TableName}}Router 路由
func {{.Table.TableName}}Router(engin *gin.Engine) {

	// 定义路由组
	group := engin.Group("/{{.Table.Uri}}")

	// 绑定分页接口
	group.POST("/queryPage", controller.{{.Table.TableName}}Controller{}.{{.Table.TableName}}QueryPage)

	// 绑定新增接口
    group.PUT("/add", controller.{{.Table.TableName}}Controller{}.{{.Table.TableName}}Add)

    // 绑定更新接口
    group.PUT("/update", controller.{{.Table.TableName}}Controller{}.{{.Table.TableName}}Update)

    // 绑定查询接口
    group.GET("/queryDetails/:id", controller.{{.Table.TableName}}Controller{}.{{.Table.TableName}}QueryDetails)

    // 绑定删除接口
    group.DELETE("/delete/:id", controller.{{.Table.TableName}}Controller{}.{{.Table.TableName}}Delete)

    // 绑定批量删除接口
    group.DELETE("/deleteBatch", controller.{{.Table.TableName}}Controller{}.{{.Table.TableName}}DeleteBatch)
}{{end}}