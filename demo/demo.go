package main

import (
	"testing"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/stretchr/testify/assert"
)

func demo(data []byte, t *testing.T) Bit64 {
	var a, result, x, c, z Bit64
	mylog.HexDump("input data", data)
	// 00EFF959  0000F3FF
	// 00EFF95D  00000000
	// 00EFF961  00001302
	// 00EFF965  00000000
	assert.Equal(t, uint32(0x02F4963B), z.Low)
	assert.Equal(t, uint32(0x00080C29), z.High)
	x = mul(uint32(data[1])+0xf366, 0, 0x1302, 0)
	assert.Equal(t, uint32(0x0000F3FF), x.Low)  // EAX
	assert.Equal(t, uint32(0x00000000), x.High) // EDX

	y := x                   //.high
	x.Low += uint32(data[1]) // 0x121DD597

	// 00EFF959  121DD597
	// 00EFF95D  00000000
	// 00EFF961  0071B793
	// 00EFF965  00000000
	assert.Equal(t, uint32(0x121DD597), z.Low)
	assert.Equal(t, uint32(0x00000000), z.High)
	x = mul(x.Low, x.High, 0x71b793, 0)
	assert.Equal(t, uint32(0x0000F3FF), x.Low)  // EAX
	assert.Equal(t, uint32(0x00000000), x.High) // EDX

	z = x // 0x27796B5
	z.Low += 0x7CFF86

	elem2 := uint32(data[2])
	y.Low = uint32(z.Uint64() >> 18 & 0xFFFFFFFF)
	// stream.GenMask()
	// data := []byte{0x9, 0x99, 0x8a, 0x7b, 0xfe, 0x46, 0xc2, 0xf0}

	local10f := z.ShiftRight(18)
	assert.Equal(t, uint32(0x0000F3FF), x.Low)  // EAX
	assert.Equal(t, uint32(0x00000000), x.High) // EDX

	//$ ==>     02F4963B
	//$+4       00080C29
	//$+8       6381BE9A
	//$+C       00000000
	assert.Equal(t, uint32(0x02F4963B), z.Low)
	assert.Equal(t, uint32(0x00080C29), z.High)
	x = div(z.Low, z.High, 0x6381be9a, 0)       // 0x14B426 感觉是少了一半的值没有存储
	assert.Equal(t, uint32(0x0000F3FF), x.Low)  // EAX
	assert.Equal(t, uint32(0x00000000), x.High) // EDX

	x.Low += y.Low

	//$ ==>     031EF4E3
	//$+4       00000002
	//$+8       00000002
	//$+C       00000000
	assert.Equal(t, uint32(0x031EF4E3), x.Low)
	assert.Equal(t, uint32(0x00000002), local10f.Low)
	x = mul(x.Low, local10f.Low, 2, 0)          // 0x63DE9C6 passed
	assert.Equal(t, uint32(0x0000F3FF), x.Low)  // EAX
	assert.Equal(t, uint32(0x00000000), x.High) // EDX
	//$ ==>     0000F3F0
	//$+4       00000000
	//$+8       00001634
	//$+C       00000000
	assert.Equal(t, 0x0000F3F0, elem2+0xf366)
	c = mul(elem2+0xf366, 0, 0x1634, 0)         // 0x15282CC0 passed
	assert.Equal(t, uint32(0x0000F3FF), x.Low)  // EAX
	assert.Equal(t, uint32(0x00000000), x.High) // EDX

	var5 := c.Low + elem2
	//$ ==>     15282D4B
	//$+4       00000000
	//$+8       02F4963B
	//$+C       00080C29
	assert.Equal(t, uint32(0x15282D4B), var5+1)
	assert.Equal(t, uint32(0x00000000), y.High+c.High)
	assert.Equal(t, uint32(0x02F4963B), z.Low)
	assert.Equal(t, uint32(0x00080C29), z.High)
	c = mul(var5+1, y.High+c.High, z.Low, z.High)
	assert.Equal(t, uint32(0x0000F3FF), x.Low)  // EAX
	assert.Equal(t, uint32(0x00000000), x.High) // EDX

	c.High += 4 //?? why // todo bug

	z.Low = c.Low
	z.High = c.High
	z.Low += x.Low + 0x2d1f65
	uVar11 := z
	// uVar11.High = uVar11.Low //

	var7 := uint32(data[3])
	r := uVar11.ShiftRight(18)
	assert.Equal(t, uint32(0x0000F3FF), x.Low)  // EAX
	assert.Equal(t, uint32(0x00000000), x.High) // EDX
	// elem2 = uVar11.Low << 14
	// uVar10 := uVar11.High >> 18
	//$ ==>     32B36B74
	//$+4       B0254C17
	//$+8       6381BE9A
	//$+C       00000000
	assert.Equal(t, uint32(0x32B36B74), z.Low)
	assert.Equal(t, uint32(0xB0254C17), z.High)
	x = div(uVar11.Low, uVar11.High, 0x6381be9a, 0)
	assert.Equal(t, uint32(0x0000F3FF), x.Low)  // EAX
	assert.Equal(t, uint32(0x00000000), x.High) // EDX

	// x.Low += (uVar10 << 32) | (var5 | elem2) + 0x21d78d
	x.Add(r)
	x.AddInt(0x21d78d)
	assert.Equal(t, uint32(0x0000F3FF), x.Low)  // EAX
	assert.Equal(t, uint32(0x00000000), x.High) // EDX
	//$ ==>     1852A1B4
	//$+4       00002C0B
	//$+8       00000003
	//$+C       00000000
	assert.Equal(t, uint32(0x1852A1B4), z.Low)
	assert.Equal(t, uint32(0x00002C0B), z.High)
	x = mul(x.Low, x.High, 3, 0)
	assert.Equal(t, uint32(0x0000F3FF), x.Low)  // EAX
	assert.Equal(t, uint32(0x00000000), x.High) // EDX

	//$ ==>     0000F3E1
	//$+4       00000000
	//$+8       00001968
	//$+C       00000000
	assert.Equal(t, 0x0000F3E1, var7+0xf366)
	c = mul(var7+0xf366, 0, 0x1968, 0)
	assert.Equal(t, uint32(0x0000F3FF), x.Low)  // EAX
	assert.Equal(t, uint32(0x00000000), x.High) // EDX

	y.Low = c.High
	var5 = c.High + var7
	//$ ==>     18340CE4
	//$+4       00000000
	//$+8       32B36B74
	//$+C       B0254C17
	assert.Equal(t, uint32(0x18340CE4), var5+1)
	assert.Equal(t, uint32(0x00000000), y.Low+c.High|(var7+1))
	assert.Equal(t, uint32(0x32B36B74), z.Low)
	assert.Equal(t, uint32(0xB0254C17), z.High)
	c = mul(var5+1, y.Low+c.High|(var7+1), z.Low, z.High)
	assert.Equal(t, uint32(0x0000F3FF), x.Low)  // EAX
	assert.Equal(t, uint32(0x00000000), x.High) // EDX

	result = uVar11
	z = c
	z.Low += x.Low
	var7 = z.High

	local13b := result.High
	a = z
	//$ ==>     8D5D086C
	//$+4       42407DAF
	//$+8       6381BE9A
	//$+C       00000000
	//todo
	x = div(z.Low, z.Low, 0x6A, 0)
	assert.Equal(t, uint32(0x0000F3FF), x.Low)  // EAX
	assert.Equal(t, uint32(0x00000000), x.High) // EDX

	z.Low = x.Low + (local13b | a.High)
	high := a.High
	//$ ==>     0000F464
	//$+4       00000000
	//$+8       00001C9E
	//$+C       00000000
	assert.Equal(t, 0x0000F464, high+0xf366)
	x = mul(high+0xf366, 0, 0x1c9e, 0)
	assert.Equal(t, uint32(0x0000F3FF), x.Low)  // EAX
	assert.Equal(t, uint32(0x00000000), x.High) // EDX

	var5 = x.Low + high
	y = Bit64{Low: x.High, High: high}
	//$ ==>     1B51C6B7
	//$+4       00000000
	//$+8       8D5D086C
	//$+C       42407DAF
	assert.Equal(t, uint32(0x1B51C6B7), var5+1)
	assert.Equal(t, uint32(0x00000000), y.Low)
	assert.Equal(t, uint32(0x8D5D086C), z.Low)
	assert.Equal(t, uint32(0x42407DAF), z.High) //?
	x = mul(var5+1, y.Low, y.Low, var7)
	assert.Equal(t, uint32(0x0000F3FF), x.Low)  // EAX
	assert.Equal(t, uint32(0x00000000), x.High) // EDX

	x.Low += 0x2d1f65 // todo test
	x.High = 0
	result = x
	c = x

	i := 7
	for i >= 0 {
		j := i
		v := result.High + 4
		m := (v<<25^result.High)>>25 | (v>>7^v)<<7
		local113 := v >> 25
		x = div(result.High, v, 0x6a, 0)
		k := x
		k.Low += local113 | m
		c = x
		elem := data[j]
		c = mul(uint32(j*j), uint32(j*j>>31), k.High, k.High)
		y.Low = c.High // low
		x = mul(uint32((elem)+1), uint32(int(elem)>>31+i), result.High, v)
		// todo missing a div ?
		i--
	}
	return result
}

func TestRange() {
	x := Bit64{Low: 0x12345678, High: 0x9abcdef0}
	y := Bit64{Low: 0x12345678, High: 0x9abcdef0}
	c := Bit64{Low: 0x12345678, High: 0x9abcdef0}
	result := Bit64{Low: 0x12345678, High: 0x9abcdef0}
	data := []byte{0x9, 0x99, 0x8a, 0x7b, 0xfe, 0x46, 0xc2, 0xf0}
	i := 7
	for i >= 0 {
		j := i
		v := result.High + 4
		m := (v<<25^result.High)>>25 | (v>>7^v)<<7
		local113 := v >> 25
		x = div(result.High, v, 0x6a, 0)
		k := x
		k.Low += local113 | m
		c = x
		elem := data[j]
		c = mul(uint32(j*j), uint32(j*j>>31), k.High, k.High)
		y.Low = c.High // low
		x = mul(uint32((elem)+1), uint32(int(elem)>>31+i), result.High, v)
		// todo missing a div ?
		i--
	}
}
