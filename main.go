package main

import "reflect"

func main() {
	// 测试字符串索引字符类型
	s := "Hello, World"
	s2 := "你好，世界1234567"
	println((reflect.TypeOf(s[0])).String())
	println((reflect.TypeOf(s2[0])).String())
	println(len(s2))        // 计算字节数
	println(s2[:5])         // 依照byte截取
	println(string(s2[20])) // 依照byte索引
}
