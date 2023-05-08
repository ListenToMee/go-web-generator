{{define "service"}}package service

import (
    "{{.ModuleName}}/dto"
	"{{.ModuleName}}/vo"
)

// QueryPage {{.Table.TableComment}}分页查询
func QueryPage(param dto.{{.Table.TableName}}PageDTO) []vo.{{.Table.TableName}}VO{

	return []vo.{{.Table.TableName}}VO{}
}

// Add 新增{{.Table.TableComment}}
func Add(param dto.{{.Table.TableName}}AddDTO) bool {

	return true
}

// Update 更新{{.Table.TableComment}}
func Update(param dto.{{.Table.TableName}}AddDTO) bool{

	return true
}

// QueryDetails 查询{{.Table.TableComment}}
func QueryDetails(id int64) vo.{{.Table.TableName}}VO{

	return vo.{{.Table.TableName}}VO{}
}

// Delete 删除{{.Table.TableComment}}
func Delete(id int64) bool {

	return true
}

// DeleteBatch 批量删除{{.Table.TableComment}}
func DeleteBatch(id []int64) bool {

	return true
}{{end}}