{{define "controller"}}package controller

import (
    "github.com/gin-gonic/gin"
)

type {{.Table.TableName}}Controller struct {
}

// QueryPage {{.Table.TableComment}}分页查询
func ({{.Table.TableName}}Controller {{.Table.TableName}}Controller) QueryPage(ctx *gin.Context) {

}

// Add 新增{{.Table.TableComment}}
func ({{.Table.TableName}}Controller {{.Table.TableName}}Controller) Add(ctx *gin.Context) {

}

// Update 更新{{.Table.TableComment}}
func ({{.Table.TableName}}Controller {{.Table.TableName}}Controller) Update(ctx *gin.Context) {

}

// QueryDetails 查询{{.Table.TableComment}}
func ({{.Table.TableName}}Controller {{.Table.TableName}}Controller) QueryDetails(ctx *gin.Context) {

}

// Delete 删除{{.Table.TableComment}}
func ({{.Table.TableName}}Controller {{.Table.TableName}}Controller) Delete(ctx *gin.Context) {

}

// DeleteBatch 批量删除{{.Table.TableComment}}
func ({{.Table.TableName}}Controller {{.Table.TableName}}Controller) DeleteBatch(ctx *gin.Context) {

}{{end}}