{{define "po"}}package po

import "time"

// {{.Table.TableName}}PO {{.Table.TableComment}}
type {{.Table.TableName}}PO struct {
    {{range .Fields}}
    {{.CamelField}} {{.RealType}} `json:"{{.ColumnName}}"{{.KeyStr }}` // {{.ColumnComment}} {{end}}
}

{{end}}
