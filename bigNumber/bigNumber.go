package main

import (
	"github.com/ddkwork/golibrary/mylog"
)

func bigNumber(data []byte) LargeInteger {
	var uVar1, result, x, lVar12, uVar13 LargeInteger
	mylog.HexDump("input data", data)
	x = mul(uint32(data[1])+0xf366, 0, 0x1302, 0)
	y := x                   //.high
	x.Low += uint32(data[1]) // 0x121DD597
	x = mul(x.Low, x.Low>>32, 0x71b793, 0)
	uVar13 = x // 0x27796B5
	uVar13.Low += 0x7CFF86
	x.Debug()

	uVar9 := uint32(data[2])
	y.Low = uint32(uVar13.Self() >> 18 & 0xFFFFFFFF)
	// stream.GenMask()

	local10f := LargeInteger{}
	local10f.Low = uVar13.High >> 18
	// 012FFAC5     02F4963B
	// 012FFAC9     00080C29
	// 012FFACD     6381BE9A
	// 012FFAD1     00000000
	x = div(uVar13.Low, uVar13.High, 0x6381be9a, 0) // 0x14B426 感觉是少了一半的值没有存储
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
	lVar12 = mul(uVar9+0xf366, 0, 0x1634, 0) // 0x15282CC0 passed

	uVar5 := lVar12.Low + uVar9
	// 00DCFCA1     15282D4B
	// 00DCFCA5     00000000
	// 00DCFCA9     02F4963B
	// 00DCFCAD     00080C29
	lVar12 = mul(uVar5+1, y.High+lVar12.High, uVar13.Low, uVar13.High)

	uVar13.Low = lVar12.Low
	uVar13.High = lVar12.High
	uVar13.Low += x.Low // + 0x2d1f65
	uVar11 := uVar13
	uVar11.High = uVar11.Low >> 32

	uVar7 := uint32(data[3])
	uVar5 = uVar11.Low >> 18
	uVar9 = uVar11.Low << 14
	uVar10 := uVar11.High >> 18
	// 012FFBC9     32B36B74
	// 012FFBCD     B0254C17 bug
	// 012FFBD1     6381BE9A
	// 012FFBD5     00000000
	x = div(uVar11.Low, uVar11.High, 0x6381be9a, 0) // todo bug

	x.Low += (uVar10 << 32) | (uVar5 | uVar9) + 0x21d78d
	x = mul(x.High, x.Low>>32, 3, 0)

	lVar12 = mul(uVar7+0xf366, 0, 0x1968, 0)
	y.Low = lVar12.High
	uVar5 = lVar12.High + uVar7
	lVar12 = mul(uVar5+1, y.Low+lVar12.High|(uVar7+1), uVar13.High, uVar13.High)

	result = uVar11
	uVar13 = lVar12
	uVar13.Low += x.Low
	uVar7 = uVar13.High

	local13b := result.High >> 32
	uVar1 = uVar13
	x = div(uVar13.Low, uVar13.Low>>32, 0x6A, 0)
	uVar13.Low = x.Low + (local13b<<32 | uVar1.High)

	local9 := uVar1.High
	x = mul(local9+0xf366, 0, 0x1c9e, 0)
	uVar5 = x.Low + local9
	y = LargeInteger{Low: x.High, High: local9}
	x = mul(uVar5+1, y.Low, y.Low, uVar7)

	i := 7
	for i >= 0 {
		iVar4 := i
		uVar5 := result.High + 4
		local117 := (uVar5<<25^result.High)>>25 | (uVar5>>7^uVar5)<<7
		local113 := uVar5 >> 25
		x = div(result.High, uVar5, 0x6a, 0)
		local127 := x
		local127.Low += local113<<32 | local117
		lVar12 = x
		local11b := data[iVar4]
		lVar12 = mul(uint32(iVar4*iVar4), uint32(iVar4*iVar4>>31), local127.High, local127.High>>32)
		y.Low = lVar12.High // low >> 32
		x = mul(uint32((local11b)+1), uint32(int(local11b)>>31+i), result.High, uVar5)
		i--
	}
	return result
}
