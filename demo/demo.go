package main

import (
	"testing"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/stretchr/testify/assert"
)

func demo(data []byte, t *testing.T) Bit64 {
	var a, result, x, c, z Bit64
	mylog.HexDump("input data", data)
	// 0x0000F3FF
	// 0x00000000
	// 0x00001302
	// 0x00000000
	assert.Equal(t, uint32(0x0000F3FF), uint32(data[1])+0xf366)
	x = mul(uint32(data[1])+0xf366, 0, 0x1302, 0)
	assert.Equal(t, uint32(0x121DD4FE), x.Low)  // EAX
	assert.Equal(t, uint32(0x00000000), x.High) // EDX

	y := x                   //.high
	x.Low += uint32(data[1]) // 0x121DD597

	// 00EFF959  121DD597
	// 00EFF95D  00000000
	// 00EFF961  0071B793
	// 00EFF965  00000000
	assert.Equal(t, uint32(0x121DD597), x.Low)
	assert.Equal(t, uint32(0x00000000), x.High)
	x = mul(x.Low, x.High, 0x71b793, 0)
	assert.Equal(t, uint32(0x27796b5), x.Low)   // EAX
	assert.Equal(t, uint32(0x00080C29), x.High) // EDX

	z = x // 0x27796B5
	z.Low += 0x7CFF86

	elem2 := uint32(data[2])
	y.Low = uint32(z.Uint64() >> 18 & 0xFFFFFFFF)
	// stream.GenMask()
	// data := []byte{0x9, 0x99, 0x8a, 0x7b, 0xfe, 0x46, 0xc2, 0xf0}

	local10f := z.ShiftRight(18)
	// assert.Equal(t, uint32(0x0000F3FF), x.Low)  // EAX
	// assert.Equal(t, uint32(0x00000000), x.High) // EDX

	// 0x02F4963B
	// 0x00080C29
	// 0x6381BE9A
	// 0x00000000
	assert.Equal(t, uint32(0x02F4963B), z.Low)
	assert.Equal(t, uint32(0x00080C29), z.High)

	x = div(z.Low, z.High, 0x6381be9a, 0)       // 0x14B426 感觉是少了一半的值没有存储
	assert.Equal(t, uint32(0x14b426), x.Low)    // EAX
	assert.Equal(t, uint32(0x00000000), x.High) // EDX

	x.Low += y.Low

	//$ ==>     031EF4E3
	//$+4       00000002
	//$+8       00000002
	//$+C       00000000
	assert.Equal(t, uint32(0x031EF4E3), x.Low)
	assert.Equal(t, uint32(0x00000002), local10f.High)

	x = mul(x.Low, local10f.High, 2, 0)       // 0x63DE9C6 passed
	assert.Equal(t, uint32(0x63de9c6), x.Low) // EAX
	assert.Equal(t, uint32(0x4), x.High)      // EDX

	// 006FFA75  0000F3F0
	// 006FFA79  00000000
	// 006FFA7D  00001634
	// 006FFA81  00000000
	assert.Equal(t, uint32(0x0000F3F0), elem2+0xf366) // EAX
	c = mul(elem2+0xf366, 0, 0x1634, 0)               // 0x15282CC0 passed
	assert.Equal(t, uint32(0x15282cc0), c.Low)        // EAX
	assert.Equal(t, uint32(0x00000000), c.High)       // EDX

	var5 := c.Low + elem2
	// 0x15282D4B
	// 0x00000000
	// 0x02F4963B
	// 0x00080C29
	assert.Equal(t, uint32(0x15282D4B), var5+1)
	assert.Equal(t, uint32(0x00000000), y.High+c.High)
	assert.Equal(t, uint32(0x02F4963B), z.Low)
	assert.Equal(t, uint32(0x00080C29), z.High)

	c = mul(var5+1, y.High+c.High, z.Low, z.High)
	assert.Equal(t, uint32(0x2c486249), c.Low)  // EAX
	assert.Equal(t, uint32(0xb0254c13), c.High) // EDX

	c.High += 4 //?? why // todo bug

	z.Low = c.Low
	z.High = c.High
	z.Low += x.Low + 0x2d1f65
	uVar11 := z
	// uVar11.High = uVar11.Low //

	var7 := uint32(data[3])
	r := uVar11.ShiftRight(18)
	// 0x32B36B74
	// 0xB0254C17
	// 0x6381BE9A
	// 0x00000000
	assert.Equal(t, uint32(0x32B36B74), uVar11.Low)  // EAX
	assert.Equal(t, uint32(0xB0254C17), uVar11.High) // EDX
	x = div(uVar11.Low, uVar11.High, 0x6381be9a, 0)
	assert.Equal(t, uint32(0xc52afd7b), x.Low) // EAX
	assert.Equal(t, uint32(0x1), x.High)       // EDX

	// x.Low += (uVar10 << 32) | (var5 | elem2) + 0x21d78d
	x.Add(r)
	x.AddInt(0x21d78d)
	// 006FFA75  1852A1B4
	// 006FFA79  00002C0B
	// 006FFA7D  00000003
	// 006FFA81  00000000
	assert.Equal(t, uint32(0x1852A1B4), x.Low)  // EAX
	assert.Equal(t, uint32(0x00002C0B), x.High) // EDX
	x = mul(x.Low, x.High, 3, 0)
	assert.Equal(t, uint32(0x48f7e51c), x.Low) // EAX
	assert.Equal(t, uint32(0x8421), x.High)    // EDX

	//$ ==>     0000F3E1
	//$+4       00000000
	//$+8       00001968
	//$+C       00000000
	assert.Equal(t, uint32(0x0000F3E1), var7+0xf366)
	c = mul(var7+0xf366, 0, 0x1968, 0)
	assert.Equal(t, uint32(0x18340c68), c.Low)  // EAX
	assert.Equal(t, uint32(0x00000000), c.High) // EDX

	y.Low = c.Low
	var5 = c.Low + var7
	//$ ==>     18340CE4
	//$+4       00000000
	//$+8       32B36B74
	//$+C       B0254C17
	assert.Equal(t, uint32(0x18340CE4), var5+1)
	assert.Equal(t, uint32(0x00000000), y.High) // todo test more data
	assert.Equal(t, uint32(0x32B36B74), z.Low)
	assert.Equal(t, uint32(0xB0254C17), z.High)
	// 00A21895   | E8  | call <asm.__allmul>                                   | main.c:193
	c = mul(var5+1, y.High, z.Low, z.High)
	assert.Equal(t, uint32(0x44652350), c.Low)  // EAX
	assert.Equal(t, uint32(0x423ff98e), c.High) // EDX

	z.Low = x.Low + c.Low
	z.High = c.High + x.High
	//$ ==>     8D5D086C
	//$+4       42407DAF
	//$+8       6381BE9A
	//$+C       00000000
	assert.Equal(t, uint32(0x8D5D086C), z.Low)
	assert.Equal(t, uint32(0x42407DAF), z.High)
	x = div(z.Low, z.High, 0x6381BE9A, 0)
	// 00A218C7   | E8  | call <asm.__alldiv>                                   | main.c:210
	assert.Equal(t, uint32(0xaa720dbb), x.Low)  // EAX
	assert.Equal(t, uint32(0x00000000), x.High) // EDX

	back := z.Low
	z.Low = x.Low + (z.High | a.High)
	high := uint32(data[4])
	//$ ==>     0000F464
	//$+4       00000000
	//$+8       00001C9E
	//$+C       00000000
	assert.Equal(t, uint32(0x0000F464), high+0xf366)
	x = mul(high+0xf366, 0, 0x1c9e, 0)
	assert.Equal(t, uint32(0x1b51c5b8), x.Low)  // EAX
	assert.Equal(t, uint32(0x00000000), x.High) // EDX

	var5 = x.Low + high
	y = Bit64{Low: x.High, High: high}
	//$ ==>     1B51C6B7
	//$+4       00000000
	//$+8       8D5D086C
	//$+C       42407DAF
	assert.Equal(t, uint32(0x1B51C6B7), var5+1)
	assert.Equal(t, uint32(0x00000000), y.Low)
	assert.Equal(t, uint32(0x8D5D086C), back)
	assert.Equal(t, uint32(0x42407DAF), z.High) //?

	x = mul(var5+1, y.Low, back, z.High)
	assert.Equal(t, uint32(0x10308d34), x.Low)  // EAX
	assert.Equal(t, uint32(0x85c42a27), x.High) // EDX
	return x

	//todo 00A2192D   | 6A  | push 0                                                | main.c:242
	//这里开始还有个乘法猜到循环的

	// push 0
	// x.High
	// push 1090?
	// push 4
	// x.Low

	// 00B2FB0D  C9DDF112
	// 00B2FB11  00001090
	// 00B2FB15  00000004
	// 00B2FB19  00000000

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
