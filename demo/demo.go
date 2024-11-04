package main

import (
	"github.com/ddkwork/golibrary/mylog"
)

func demo(data []byte) Bit64 {
	var a, result, x, c, z Bit64
	mylog.HexDump("input data", data)
	x = mul(uint32(data[1])+0xf366, 0, 0x1302, 0)
	y := x                   //.high
	x.Low += uint32(data[1]) // 0x121DD597
	x = mul(x.Low, x.Low, 0x71b793, 0)
	z = x // 0x27796B5
	z.Low += 0x7CFF86

	elem2 := uint32(data[2])
	y.Low = uint32(z.Self() >> 18 & 0xFFFFFFFF)
	// stream.GenMask()

	local10f := Bit64{}
	local10f.Low = z.High >> 18
	// 012FFAC5     02F4963B
	// 012FFAC9     00080C29
	// 012FFACD     6381BE9A
	// 012FFAD1     00000000
	x = div(z.Low, z.High, 0x6381be9a, 0) // 0x14B426 感觉是少了一半的值没有存储
	x.Low += y.Low
	// 0096F7CD     031EF4E3
	// 0096F7D1     00000002
	// 0096F7D5     00000002
	// 0096F7D9     00000000
	x = mul(x.Low, local10f.Low, 2, 0) // 0x63DE9C6 passed

	// 006FF901     0000F3F0
	// 006FF905     00000000
	// 006FF909     00001634
	// 006FF90D     00000000
	c = mul(elem2+0xf366, 0, 0x1634, 0) // 0x15282CC0 passed

	var5 := c.Low + elem2
	// 00DCFCA1     15282D4B
	// 00DCFCA5     00000000
	// 00DCFCA9     02F4963B
	// 00DCFCAD     00080C29
	c = mul(var5+1, y.High+c.High, z.Low, z.High)
	c.High += 4 //?? why // todo bug

	z.Low = c.Low
	z.High = c.High
	z.Low += x.Low + 0x2d1f65
	uVar11 := z
	// uVar11.High = uVar11.Low //

	var7 := uint32(data[3])
	r := uVar11.ShiftRight(18)
	// elem2 = uVar11.Low << 14
	// uVar10 := uVar11.High >> 18
	mylog.Hex("bug", uVar11.Self()>>18)
	// 012FFBC9     32B36B74
	// 012FFBCD     B0254C17 bug
	// 012FFBD1     6381BE9A
	// 012FFBD5     00000000
	x = div(uVar11.Low, uVar11.High, 0x6381be9a, 0)

	// x.Low += (uVar10 << 32) | (var5 | elem2) + 0x21d78d
	x.Add(r)
	x.AddInt(0x21d78d)

	// 008FFAED     1852A1B4
	// 008FFAF1     00002C0B
	// 008FFAF5     00000003
	// 008FFAF9     00000000
	x = mul(x.Low, x.Low, 3, 0) // todo bug

	c = mul(var7+0xf366, 0, 0x1968, 0)
	y.Low = c.High
	var5 = c.High + var7
	c = mul(var5+1, y.Low+c.High|(var7+1), z.High, z.High)

	result = uVar11
	z = c
	z.Low += x.Low
	var7 = z.High

	local13b := result.High
	a = z
	x = div(z.Low, z.Low, 0x6A, 0)
	z.Low = x.Low + (local13b | a.High)

	high := a.High
	x = mul(high+0xf366, 0, 0x1c9e, 0)
	var5 = x.Low + high
	y = Bit64{Low: x.High, High: high}
	x = mul(var5+1, y.Low, y.Low, var7)

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
		i--
	}
	return result
}
