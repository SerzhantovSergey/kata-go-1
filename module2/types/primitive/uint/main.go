package main

import (
	"fmt"
	"unsafe"
)

func main() {
	n := 112358132134
	fmt.Println("size is:", unsafe.Sizeof(n), "bytes")

	typeUint()

}

func typeUint() {
	fmt.Println("===START type uint.===")
	var numberUint8 uint8 = 1 << 1
	fmt.Println("left shift uint8:", numberUint8)
	numberUint8 = (1 << 8) - 1
	fmt.Println("uint8 max value:", numberUint8)
	var numberUint16 uint16 = (1 << 16) - 1
	fmt.Println("uint16 max value:", numberUint16)
	var numberUint32 uint32 = (1 << 32) - 1
	fmt.Println("uint32 max value:", numberUint32)
	var numberUint64 uint64 = (1 << 64) - 1
	fmt.Println("uint64 max value:", numberUint64)
}
