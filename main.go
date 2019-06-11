package main

<<<<<<< HEAD
import (
	"fmt"
	"reflect"
)

func test(name string, obj interface{}) { // 要点：iterface{} 相当于java中的object, 能接受任何类型的数据
	rt := reflect.TypeOf(obj)
	switch rt.Kind() {
	case reflect.Slice:
		arr := obj.([]int)
		fmt.Printf("%s is a slice, len=%d, cap=%d, pointer=%p\n", name, len(arr), cap(arr), obj)
	case reflect.Array:
		arr := obj.([6]int)
		fmt.Printf("%s is an array, len=%d, pointer=%p\n", name, len(arr), obj)
	default:
		fmt.Printf(name, "is neither a slice or a array")
	}
}

func main() {
	// 测试slice
	A1 := [6]int{1, 2, 3, 4, 5, 6}
	test("A1", A1)
	// 输出：A1 is an array, len=6, pointer=%!p([6]int=[1 2 3 4 5 6])
	// 解释：A1是一个大小为6的数组
	A2 := A1[0:3] // 从数组创建切片
	test("A2", A2)
	// 输出：A2 is a slice, len=3, cap=6, pointer=0xc00000a2a0
	A3 := append(A2, 1) // 将内容加入到切片中
	test("A3", A3)
	// 输出：A3 is a slice, len=4, cap=6, pointer=0xc000082030
	test("A2", A2)
	// 输出：A2 is a slice, len=3, cap=6, pointer=0xc000082030
	//
	test("A1", A1)
	// 输出：A1 is an array, len=6, pointer=%!p([6]int=[1 2 3 1 5 6]
	// 解释：对数组切片元素的修改直接影响到了原数组的元素

	// 从输出可以看到修改了A1的内容
	A4 := append(A2, -1, -1, -1, -1) // 新增4个元素
	test("A4", A4)                   //
	// A4 is a slice, len=7, cap=12, pointer=0xc000040060
	test("A1", A1)
	// A1 is an array, len=6, pointer=%!p([6]int=[1 2 3 1 5 6])
	// 新增4个元素，可以看到元素组扩容成原来的两倍(cap从6变成了12)，因此自动重新开辟空间，原来的数组没被修改

	println()

	B1 := make([]int, 5) // 首先创建了大小为5的数组，然后将指针
	test("B1", B1)       // 查看元素
	B2 := append(B1, 1)  // 新增一个元素自动扩容到原来的2倍，指针地址改变 test("B2", B2)
	test("B2", B2)

	println()

	C1 := make([]int, 2, 3)
	test("C1", C1)
	C2 := append(C1, 1) // 新增第一个元素，未超过最大容量，指针地址不变
	test("C2", C2)
	test("C1", C1)      // 用作对比：可以看到，尽管C1和C2的指针地址相同，但是这里C1和C2的大小是不一样的
	C3 := append(C2, 1) // 再次新增第二个元素，
	test("C3", C3)
	test("C2", C2)
	test("C1", C1)
=======
func main() {
	arr := buildIntArray("[-1, 2, 1, -4]")
	println(threeSumClosest(arr, 1))
>>>>>>> 8cebf84ab2402b818677e9a70bfdc43858dd9ff8
}
