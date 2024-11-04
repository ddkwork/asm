package main

import (
	"github.com/ddkwork/golibrary/mylog"
)

func bigNumber(data []byte) LargeInteger {
	var uVar1, result, x, lVar12, uVar13 LargeInteger
	mylog.HexDump("input data", data)
	x = mul(uint32(data[1])+0xf366, 0, 0x1302, 0)
	mylog.Hex("", 0xF366+0x99)
	mylog.Hex("", 0xF366+0x99*0x1302)
	mylog.Hex("", 0x1302*(0x99+0xF366))
	mylog.Hex("", 0x1302*(0x99+0xF366)+0x99)
	v := 0x1302*(0x99+0xF366) + 0x99
	v *= 0x71B793
	v += 0x7CFF86
	v >>= 0x12
	v /= 0x6381BE9A
	mylog.Hex("", v)

	local13f := x.HighPart
	x.LowPart += uint32(data[1]) // 0x121DD597
	x = mul(x.LowPart, x.LowPart>>32, 0x71b793, 0)
	uVar13 = x // 0x27796B5
	uVar13.LowPart += 0x7CFF86
	mylog.Hex("mul1 bug", 0x121DD597*0x71b793)
	mylog.Struct(x)

	uVar9 := uint32(data[2])
	local13f = uVar13.LowPart >> 18 // todo bug   shrd edx, ecx, 0x12
	local10f := (uVar13.LowPart >> 18) >> 32
	x = div(uVar13.LowPart, uVar13.LowPart>>32, 0x6381be9a, 0) // todo bug hight is not right
	x.LowPart += (local10f << 32) | local13f
	x = mul(x.HighPart, x.LowPart>>32, 2, 0)

	lVar12 = mul(uVar9+0xf366, 0, 0x1634, 0)
	local13f = lVar12.HighPart
	uVar5 := lVar12.HighPart + uVar9
	lVar12 = mul(uVar5+1, local13f+(lVar12.HighPart<<16|uVar9)+1, uVar13.HighPart, uVar13.LowPart)

	uVar13 = lVar12
	uVar13.LowPart += x.LowPart + 0x2d1f65
	uVar11 := uVar13
	uVar11.HighPart = uVar11.LowPart >> 32

	uVar7 := uint32(data[3])
	uVar5 = uVar11.LowPart >> 18
	uVar9 = uVar11.LowPart << 14
	uVar10 := uVar11.HighPart >> 18
	x = div(uVar11.HighPart, uVar11.LowPart, 0x6381be9a, 0)
	x.LowPart += (uVar10 << 32) | (uVar5 | uVar9) + 0x21d78d
	x = mul(x.HighPart, x.LowPart>>32, 3, 0)

	lVar12 = mul(uVar7+0xf366, 0, 0x1968, 0)
	local13f = lVar12.HighPart
	uVar5 = lVar12.HighPart + uVar7
	lVar12 = mul(uVar5+1, local13f+lVar12.HighPart|(uVar7+1), uVar13.HighPart, uVar13.HighPart)

	result = uVar11
	uVar13 = lVar12
	uVar13.LowPart += x.LowPart
	uVar7 = uVar13.HighPart

	local13b := result.HighPart >> 32
	uVar1 = uVar13
	x = div(uVar13.LowPart, uVar13.LowPart>>32, 0x6A, 0)
	uVar13.LowPart = x.LowPart + (local13b<<32 | uVar1.HighPart)

	local9 := uVar1.HighPart
	x = mul(local9+0xf366, 0, 0x1c9e, 0)
	uVar5 = x.LowPart + local9
	local13f = x.HighPart>>32 + makeLargeInteger(x.HighPart, local9).LowPart + 1
	x = mul(uVar5+1, local13f, local13f, uVar7)

	i := 7
	for i >= 0 {
		iVar4 := i
		uVar5 := result.HighPart + 4
		local117 := (uVar5<<25^result.HighPart)>>25 | (uVar5>>7^uVar5)<<7
		local113 := uVar5 >> 25
		x = div(result.HighPart, uVar5, 0x6a, 0)
		local127 := x
		local127.LowPart += local113<<32 | local117
		lVar12 = x
		local11b := data[iVar4]
		lVar12 = mul(uint32(iVar4*iVar4), uint32(iVar4*iVar4>>31), local127.HighPart, local127.HighPart>>32)
		local13f = lVar12.LowPart >> 32
		x = mul(uint32((local11b)+1), uint32(int(local11b)>>31+i), result.HighPart, uVar5)
		i--
	}
	return result
}
