{{define "addDto"}}package dto

import "time"

// {{.Table.TableName}}AddDTO {{.Table.TableComment}}
type {{.Table.TableName}}AddDTO struct {
    {{range .Fields}}
    {{.CamelField}} {{.RealType}} `json:"{{.ColumnName}}"{{.KeyStr }}` // {{.ColumnComment}} {{end}}
}

{{end}}
