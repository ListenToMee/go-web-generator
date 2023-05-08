package main

import "go-web-generator/handle"

func main() {

	// 代码生成 // 123:456@tcp(127.0.0.1)/databaseName?tls=true
	handle.Generate("123:456@tcp(127.0.0.1)/databaseName?tls=true",
		"databaseName",
		[]string{"tableName1", "tableName2"},
		"moduleName")
}
