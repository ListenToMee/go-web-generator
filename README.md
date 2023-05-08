# go-web-generator
基于 go 模板引擎的 web开发 代码生成器 code generator 😁

<h1>怎么用</h1>

```
func main() {

	// 数据库连接url
	handle.Generate("123:456@tcp(127.0.0.1)/databaseName?tls=true",
	// 数据库名称
	"databaseName", 
	// 表名，支持多表同时生成
	[]string{"tableName1","tableName2"},
	// 模块名
	"moduleName")
}
```
<h1>说明</h1>
目前版本只支持 mysql 数据库，未完待续～～～
