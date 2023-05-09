{{define "service"}}package service

import (
    "{{.ModuleName}}/dto"
	"{{.ModuleName}}/vo"
)

// {{.Table.TableName}}QueryPage {{.Table.TableComment}}分页查询
func {{.Table.TableName}}QueryPage(param dto.{{.Table.TableName}}PageDTO) []vo.{{.Table.TableName}}VO{

	return []vo.{{.Table.TableName}}VO{}
}

// {{.Table.TableName}}Add 新增{{.Table.TableComment}}
func {{.Table.TableName}}Add(param dto.{{.Table.TableName}}AddDTO) bool {

	return true
}

// {{.Table.TableName}}Update 更新{{.Table.TableComment}}
func {{.Table.TableName}}Update(param dto.{{.Table.TableName}}AddDTO) bool{

	return true
}

// {{.Table.TableName}}QueryDetails 查询{{.Table.TableComment}}
func {{.Table.TableName}}QueryDetails(id int64) vo.{{.Table.TableName}}VO{

	return vo.{{.Table.TableName}}VO{}
}

// {{.Table.TableName}}Delete 删除{{.Table.TableComment}}
func {{.Table.TableName}}Delete(id int64) bool {

	return true
}

// {{.Table.TableName}}DeleteBatch 批量删除{{.Table.TableComment}}
func {{.Table.TableName}}DeleteBatch(id []int64) bool {

	return true
}{{end}}