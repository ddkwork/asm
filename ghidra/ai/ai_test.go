package main

import (
	"math/big"
	"testing"

	"github.com/ddkwork/golibrary/mylog"

	"github.com/ddkwork/golibrary/assert"
)

func Test_asm1(t *testing.T) {
	data := []byte{0x9, 0x99, 0x8a, 0x7b, 0xfe, 0x46, 0xc2, 0xf0}
	b := asm1(data, t)
	mylog.Hex("", b)
	assert.Equal(t, 0xCD2F21A91272EE20, b)
	// assert.Equal(t, mylog.Check2(hex.DecodeString("CD2F21A91272EE20")), b.Bytes())
}

func mul2(param1, param2, param3, param4 uint64) uint64 {
	x := param1 | param2<<32
	y := param3 | param4<<32
	z := x * y
	return z
}

// $ ==>     a:=0x15282D4B
// $+4       a:=0x00000000
// $+8       a:=0x02F4963B
// $+C       a:=0x00080C29
func TestMul(t *testing.T) {
	a := uint64(0x15282D4B)
	b := uint64(0x00000000)
	c := uint64(0x02F4963B)
	d := uint64(0x00080C29)
	mul2(a, b, c, d)
	// a|b*c|d
	// ret := uint64(0x0000000015282D4B) * uint64(0x00080C2902F4963B)
	// mylog.Hex("v", ret)

	x := new(big.Int).SetUint64(0x0000000015282D4B)
	y := new(big.Int).SetUint64(0x00080C2902F4963B)
	v := new(big.Int).Mul(x, y)
	mylog.Hex("v", v.Uint64())
}
