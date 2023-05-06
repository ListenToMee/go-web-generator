package main

import "gin-web-generator/handle"

func main() {

	// 获取连接
	connect := handle.Connect("w0lx4v7uyhb8ioi6130d:pscale_pw_VkBjB61Rf9AbRNQuGcM9kQmvpCOdcm2sBtaTdcjNJA4@tcp(aws.connect.psdb.cloud)/test?tls=true")

	handle.Generate(connect, "test", []string{"user"})
}
