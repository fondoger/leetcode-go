package main

import (
	"reflect"
	"unsafe"
)

func main() {
	/* 非常重要：在我的64位win10上，int是64位的 */

	const MaxUint32 = ^uint32(0)
	const MaxInt32 = int32(MaxUint32 >> 1)
	const MinInt32 = -int32(MaxUint32>>1) - 1

	println("test:", -23%10)

	println(reverse(0))
	println(reverse(123))
	println(reverse((int(MaxInt32)))) //  2147483647 => overflow
	println(reverse(int(MinInt32)))   // -2147483648 => overflow

	println("测试windows64上的int类型字节数")
	println(reflect.TypeOf(reverse(int(MaxInt32))).String())
	println("int竟然")
	println(unsafe.Sizeof(int(32))) // 8个字节：8*8=64位

}
