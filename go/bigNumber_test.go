package main

import (
	"encoding/hex"
	"testing"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/stretchr/testify/assert"
)

func Test_bigNumber(t *testing.T) {
	data := []byte{0x9, 0x99, 0x8a, 0x7b, 0xfe, 0x46, 0xc2, 0xf0}
	b := bigNumber(data)
	mylog.Struct(b)
	assert.Equal(t, uint64(0xCD2F21A91272EE20), b.QuadPart())
	assert.Equal(t, mylog.Check2(hex.DecodeString("CD2F21A91272EE20")), b.Bytes())
}

func TestLargeInteger(t *testing.T) {
	var div LargeInteger
	div.LowPart = 0x0000091416154ED7 % 0xFFFFFFFF
	div.HighPart = div.LowPart / 0x6A
	mylog.Struct(div)
	mylog.Hex("div", div.QuadPart())

	var mul LargeInteger
	mul.LowPart = 0x00000015ECE7BB69 % 0xFFFFFFFF
	mul.HighPart = mul.LowPart * 0x31

	mylog.Struct(div)
	mylog.Hex("div", div.QuadPart())
}
