package main

import (
	"testing"

	"github.com/ddkwork/golibrary/mylog"
)

func TestLargeInteger(t *testing.T) {
	var div LargeInteger
	div.LowPart = 0x0000091416154ED7 % 0xFFFFFFFF // 模拟低位
	div.HighPart = div.LowPart / 0x6A             // 高位计算
	mylog.Struct(div)
	mylog.Hex("div", div.QuadPart())

	var mul LargeInteger
	mul.LowPart = 0x00000015ECE7BB69 % 0xFFFFFFFF // 模拟低位
	mul.HighPart = mul.LowPart * 0x31             // 高位计算

	mylog.Struct(div)
	mylog.Hex("div", div.QuadPart())
}
