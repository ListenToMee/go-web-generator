# go-web-generator
基于 go 模板引擎的 web开发 代码生成器 code generator 😁

<h1>怎么用</h1>

```
func main() {

	// 代码生成
	handle.Generate("123:456@tcp(127.0.0.1)/databaseName?tls=true",
	"databaseName", 
	[]string{"tableName1","tableName2"},
	"moduleName")
}
```
<h1>说明</h1>
目前版本只支持 mysql 数据库，未完待续～～～
