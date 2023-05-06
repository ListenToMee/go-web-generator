package main

import "gin-web-generator/handle"

func main() {

	// 获取连接
	connect := handle.Connect("p8xq8j57fre8rtkys9pv:pscale_pw_44H7Xl3b9X8kh3uN4eXLzI1ehve2lElRdMg8apVy6jg@tcp(aws.connect.psdb.cloud)/test?tls=true")

	// 代码生成
	handle.Generate(connect, "test", []string{"user"})
}
