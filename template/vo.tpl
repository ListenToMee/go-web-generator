{{define "vo"}}package vo

import "time"

// {{.Table.TableName}}VO {{.Table.TableComment}}
type {{.Table.TableName}}VO struct {
    {{range .Fields}}
    {{.CamelField}} {{.RealType}} `json:"{{.ColumnName}}"{{.KeyStr }}` // {{.ColumnComment}} {{end}}
    Current int // 当前页码
    PageSize int // 一页数据量
    PageCount int // 页总数
}

{{end}}
