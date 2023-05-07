# go-web-generator
基于 go 模板引擎的 web开发 代码生成器 code generator

使用说明：

```
func main() {

	// 代码生成
	handle.Generate("123:456@tcp(127.0.0.1)/databaseName?tls=true",
	"databaseName", 
	[]string{"tableName1","tableName2"})
}