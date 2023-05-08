{{define "vo"}}package vo

import "time"

// {{.Table.TableName}}VO {{.Table.TableComment}}
type {{.Table.TableName}}VO struct {
    {{range .Fields}}
    {{.CamelField}} {{.RealType}} `json:"{{.ColumnName}}"{{.KeyStr }}` {{.ColumnComment}} {{end}}
}

{{end}}
