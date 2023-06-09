{{define "pageDto"}}package dto

import "time"

// {{.Table.TableName}}PageDTO {{.Table.TableComment}}
type {{.Table.TableName}}PageDTO struct {
    {{range .Fields}}
    {{.CamelField}} {{.RealType}} `json:"{{.ColumnName}}"{{.KeyStr }}` {{.ColumnComment}} {{end}}
    Current int // 当前页码
    PageSize int // 一页数据量
}

{{end}}