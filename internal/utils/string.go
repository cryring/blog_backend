// Package utils 通用工具类包
package utils

import (
	"runtime"
	"unsafe"

	"github.com/spf13/cast"
)

// Bytes2String force casts []byte to string.
//
//go:nosplit
func Bytes2String(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

// String2Bytes force casts string to []byte.
//
//go:nosplit
func String2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	d := *(*[]byte)(unsafe.Pointer(&h))
	runtime.KeepAlive(s)
	return d
}

func FormatFloat(val float64) string {
	return cast.ToString(val)
}
