package main

func main() {
	// 测试切片
	s := "Hello,World"
	s2 := s[0:3]
	print(s2)
	s2[0] = 'X'
	print(s2)
}
