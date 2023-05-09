# go-web-generator

基于 go 模板引擎的 web开发 代码生成器 code generator 😁

基于数据库表信息自动生成 po,dto,vo,service,controller,router 👍

帮您快速开始 web 业务开发 👏👏👏

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

middleware 可以自己创建和绑定

如果目录不太符合您的习惯，可自行修改template目录下的模板～

❗❗️❗️️目前版本只支持 mysql 数据库，未完待续～～～
