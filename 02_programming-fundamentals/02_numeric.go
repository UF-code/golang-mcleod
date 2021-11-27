package main

import (
	"fmt"
	"runtime"
)

// 8 bits integer value range
var minint8 int8 = -128
var maxint8 int8 = 127

// 16 bits integer value range
var minint16 int16 = -32768
var maxint16 int16 = 32767

// 32 bits integer value range
// rune which is allias for int32
var minint32 int32 = -2147483648
var maxint32 int32 = 2147483647

// 64 bits integer value range
var minint64 int64 = -9223372036854775808
var maxint64 int64 = 9223372036854775807

// 8 bits unsigned integer value range
// bytes which is allias for uint8
var minuint8 uint8 = 0
var maxuint8 uint8 = 255

// 16 bits unsigned integer value range
var minuint16 uint16 = 0
var maxuint16 uint16 = 65535

// 32 bits unsigned integer value range
var minuint32 uint32 = 0
var maxuint32 uint32 = 4294967295

// 64 bits unsigned integer value range
var minuint64 uint64 = 0
var maxuint64 uint64 = 18446744073709551615

func main() {
	fmt.Println("Integer and Unsigned Integer Value Ranges")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
