{{define "controller"}}package controller

import (
    "github.com/gin-gonic/gin"
    "{{.ModuleName}}/dto"
    "{{.ModuleName}}/service"
    "net/http"
    "strconv"
)

type {{.Table.TableName}}Controller struct {
}

// {{.Table.TableName}}QueryPage {{.Table.TableComment}}分页查询
func ({{.Table.TableName}}Controller {{.Table.TableName}}Controller) {{.Table.TableName}}QueryPage(ctx *gin.Context) {

    var param dto.{{.Table.TableName}}PageDTO

    // 绑定参数
    err := ctx.ShouldBindJSON(&param)

    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": "请求参数错误"})
        return
    }

    // 查询新增
    ctx.JSON(http.StatusOK, gin.H{"data": service.{{.Table.TableName}}QueryPage(param)})
}

// {{.Table.TableName}}Add 新增{{.Table.TableComment}}
func ({{.Table.TableName}}Controller {{.Table.TableName}}Controller) {{.Table.TableName}}Add(ctx *gin.Context) {

    var param dto.{{.Table.TableName}}AddDTO

    // 绑定参数
    err := ctx.ShouldBindJSON(&param)

    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": "请求参数错误"})
        return
    }

    // 新增
    ctx.JSON(http.StatusOK, gin.H{"data": service.{{.Table.TableName}}Add(param)})
}

// {{.Table.TableName}}Update 更新{{.Table.TableComment}}
func ({{.Table.TableName}}Controller {{.Table.TableName}}Controller) {{.Table.TableName}}Update(ctx *gin.Context) {

    var param dto.{{.Table.TableName}}AddDTO

    // 绑定参数
    err := ctx.ShouldBindJSON(&param)

    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": "请求参数错误"})
        return
    }

    // 新增
    ctx.JSON(http.StatusOK, gin.H{"data": service.{{.Table.TableName}}Update(param)})
}

// {{.Table.TableName}}QueryDetails 查询{{.Table.TableComment}}
func ({{.Table.TableName}}Controller {{.Table.TableName}}Controller) {{.Table.TableName}}QueryDetails(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "请求参数错误"})
		return
	}

	// 查询详情
	ctx.JSON(http.StatusOK, gin.H{"data": service.{{.Table.TableName}}QueryDetails(id)})
}

// {{.Table.TableName}}Delete 删除{{.Table.TableComment}}
func ({{.Table.TableName}}Controller {{.Table.TableName}}Controller) {{.Table.TableName}}Delete(ctx *gin.Context) {

    id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

    if err != nil {
    	ctx.JSON(http.StatusBadRequest, gin.H{"message": "请求参数错误"})
    	return
    }

    // 删除
    ctx.JSON(http.StatusOK, gin.H{"data": service.{{.Table.TableName}}Delete(id)})
}

// {{.Table.TableName}}DeleteBatch 批量删除{{.Table.TableComment}}
func ({{.Table.TableName}}Controller {{.Table.TableName}}Controller) {{.Table.TableName}}DeleteBatch(ctx *gin.Context) {

	var param []int64

	// 绑定参数
	err := ctx.ShouldBindJSON(&param)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "请求参数错误"})
		return
	}

	// 批量删除
	ctx.JSON(http.StatusOK, gin.H{"data": service.{{.Table.TableName}}DeleteBatch(param)})
}{{end}}