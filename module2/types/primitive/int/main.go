package main

import (
	"fmt"
	"unsafe"
)

func main() {

	typeInt()
}

func typeInt() {
	fmt.Println("===START type uint.===")
	var uintNumber uint8 = 1 << 7
	var min = int8(uintNumber)
	uintNumber--
	var max = int8(uintNumber)
	fmt.Println(min, max)
	fmt.Println("int8 min value:", min, "int8 max value", max, "size:", unsafe.Sizeof(min), "bytes")
	var uint16Number uint16 = 1 << 15
	var min1 = int16(uint16Number)
	uint16Number--
	var max1 = int16(uint16Number)
	fmt.Println(min1, max1)
	fmt.Println("int16 min value:", min1, "int16 max value", max1, "size:", unsafe.Sizeof(min1), "bytes")
	var uint32Number uint32 = 1 << 31
	var min2 = int32(uint32Number)
	uint32Number--
	var max2 = int32(uint32Number)
	fmt.Println(min2, max2)
	fmt.Println("int32 min value:", min2, "int32 max value", max2, "size:", unsafe.Sizeof(min2), "bytes")
	var uint64Number uint64 = 1 << 63
	var min3 = int64(uint64Number)
	uint64Number--
	var max3 = int64(uint64Number)
	fmt.Println(min3, max3)
	fmt.Println("int64 min value:", min3, "int64 max value", max3, "size:", unsafe.Sizeof(min3), "bytes")
	fmt.Println("*** END type uint ***")

}
