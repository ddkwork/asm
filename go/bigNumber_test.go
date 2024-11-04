package main

import (
	"encoding/hex"
	"testing"

	"github.com/ddkwork/golibrary/mylog"

	"github.com/stretchr/testify/assert"
)

func Test_bigNumber(t *testing.T) {
	b := bigNumber()
	mylog.Struct(b)
	assert.Equal(t, uint64(0xCD2F21A91272EE20), b.QuadPart())
	assert.Equal(t, mylog.Check2(hex.DecodeString("CD2F21A91272EE20")), b.Bytes())
}

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
