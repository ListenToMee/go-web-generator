{{define "controller"}}package controller

import (
    "github.com/gin-gonic/gin"
	"{{.ModuleName}}/vo"
)

type {{.Table.TableName}}Controller struct {
}

// QueryPage {{.Table.TableComment}}分页查询
func ({{.Table.TableName}}Controller {{.Table.TableName}}Controller) QueryPage(ctx *gin.Context) vo.{{.Table.TableName}}VO {

	// todo
    return vo.{{.Table.TableName}}VO{}
}

// Add 新增{{.Table.TableComment}}
func ({{.Table.TableName}}Controller {{.Table.TableName}}Controller) Add(ctx *gin.Context) bool {

	// todo
	return true
}

// Update 更新{{.Table.TableComment}}
func ({{.Table.TableName}}Controller {{.Table.TableName}}Controller) Update(ctx *gin.Context) bool {

	// todo
	return true
}

// QueryDetails 查询{{.Table.TableComment}}
func ({{.Table.TableName}}Controller {{.Table.TableName}}Controller) QueryDetails(ctx *gin.Context) vo.{{.Table.TableName}}VO {

	// todo
	return vo.{{.Table.TableName}}VO{}
}

// Delete 删除{{.Table.TableComment}}
func ({{.Table.TableName}}Controller {{.Table.TableName}}Controller) Delete(ctx *gin.Context) bool {

	// todo
	return true
}

// DeleteBatch 批量删除{{.Table.TableComment}}
func ({{.Table.TableName}}Controller {{.Table.TableName}}Controller) DeleteBatch(ctx *gin.Context) bool {

	// todo
	return true
}{{end}}