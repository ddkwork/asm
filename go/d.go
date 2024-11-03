package main

import (
	"fmt"
)

type LargeInteger struct {
	LowPart  uint32
	HighPart uint32
}

func (li LargeInteger) QuadPart() uint64 {
	return uint64(li.HighPart)<<32 | uint64(li.LowPart)
}

// makeLargeInteger 创建一个 LargeInteger
func makeLargeInteger(high, low uint32) LargeInteger {
	return LargeInteger{LowPart: low, HighPart: high}
}

// mul 进行大整数的乘法运算
func mul(param1, param2, param3, param4 uint32) LargeInteger {
	var v LargeInteger

	if (param4 | param2) == 0 {
		v.LowPart = param1 * param3
		return v
	}
	v.HighPart = param1*param3>>32 + param2*param3 + param1*param4<<32
	v.LowPart = (param1 * param3) & 0xFFFFFFFF
	return v
}

// div 进行大整数的除法运算
func div(param1, param2, param3, param4 uint32) LargeInteger {
	var v LargeInteger
	var uVar1 uint64
	var temp int64
	uVar3 := param1
	uVar8 := param4
	uVar6 := param2
	uVar9 := param3

	if param4 == 0 {
		uVar3 = param2 / param3
		temp = (int64(param2)%int64(param3)<<32 | int64(param1)) / int64(param3)
	} else {
		for uVar8 != 0 {
			uVar5 := uVar8 >> 1
			uVar9 = uVar9>>1 | (uVar8&1)<<31
			uVar7 := uVar6 >> 1
			uVar3 = uVar3>>1 | (uVar6&1)<<31
			uVar8 = uVar5
			uVar6 = uVar7
		}
		uVar7 := 0
		uVar1 = (uint64(uVar7)<<32 | uint64(uVar3)) / uint64(uVar9)
		temp = int64(uVar1)
	}
	v.HighPart = uint32(temp >> 32)
	v.LowPart = uint32(temp)
	return v
}

// ghidraDecode 主要功能
func ghidraDecode() {
	var uVar1, puVar2, lVar11, lVar12, uVar13 LargeInteger
	var local167 [4]byte

	// 初始化数据
	local167[0] = '\t'
	local167[1] = 0x99
	local167[2] = 0x8a
	local167[3] = 0x7b

	lVar11 = makeLargeInteger(0, 0)
	lVar12 = makeLargeInteger(0, 0)

	lVar11 = mul(uint32(local167[1])+0xf366, 0, 0x1302, 0)
	local13f := lVar11.HighPart
	lVar11.LowPart += uint32(local167[1])
	lVar11 = mul(lVar11.HighPart, lVar11.LowPart>>32, 0x71b793, 0)
	uVar13 = lVar11

	uVar9 := uint32(local167[2])
	local13f = uVar13.HighPart >> 18
	local10f := (uVar13.HighPart >> 18) >> 32
	lVar11 = div(uVar13.LowPart, uVar13.LowPart>>32, 0x6381be9a, 0)
	lVar11.LowPart += (local10f << 32) | local13f
	lVar11 = mul(lVar11.HighPart, lVar11.LowPart>>32, 2, 0)

	lVar12 = mul(uVar9+0xf366, 0, 0x1634, 0)
	local13f = lVar12.HighPart
	uVar5 := lVar12.HighPart + uVar9
	lVar12 = mul(uVar5+1, local13f+(lVar12.HighPart<<16|uVar9)+1, uVar13.HighPart, uVar13.LowPart)

	// 进行后续的数学运算
	uVar13 = lVar12
	uVar13.LowPart += lVar11.LowPart + 0x2d1f65
	uVar11 := uVar13
	uVar11.HighPart = uVar11.LowPart >> 32

	uVar7 := uint32(local167[3])
	uVar5 = uVar11.LowPart >> 18
	uVar9 = uVar11.LowPart << 14
	uVar10 := uVar11.HighPart >> 18
	lVar11 = div(uVar11.HighPart, uVar11.LowPart, 0x6381be9a, 0)
	lVar11.LowPart += (uVar10 << 32) | (uVar5 | uVar9) + 0x21d78d
	lVar11 = mul(lVar11.HighPart, lVar11.LowPart>>32, 3, 0)

	lVar12 = mul(uVar7+0xf366, 0, 0x1968, 0)
	local13f = lVar12.HighPart
	uVar5 = lVar12.HighPart + uVar7
	lVar12 = mul(uVar5+1, local13f+lVar12.HighPart|(uVar7+1), uVar13.HighPart, uVar13.HighPart)

	puVar2 = uVar11
	uVar13 = lVar12
	uVar13.LowPart += lVar11.LowPart
	uVar7 = uVar13.HighPart

	local13b := puVar2.HighPart >> 32
	uVar1 = uVar13
	lVar11 = div(uVar13.LowPart, uVar13.LowPart>>32, 0x6A, 0)
	uVar13.LowPart = lVar11.LowPart + (local13b<<32 | uVar1.HighPart)

	local9 := uVar1.HighPart
	lVar11 = mul(local9+0xf366, 0, 0x1c9e, 0)
	uVar5 = lVar11.LowPart + local9
	local13f = lVar11.HighPart>>32 + makeLargeInteger(lVar11.HighPart, local9).LowPart + 1
	lVar11 = mul(uVar5+1, local13f, local13f, uVar7)

	// 循环部分
	i := 7
	for i >= 0 {
		iVar4 := i
		uVar5 := puVar2.HighPart + 4 // 假设此处为真实地址
		local117 := (uVar5<<25^puVar2.HighPart)>>25 | (uVar5>>7^uVar5)<<7
		local113 := uVar5 >> 25
		lVar11 = div(puVar2.HighPart, uVar5, 0x6a, 0)
		local127 := lVar11
		local127.LowPart += local113<<32 | local117
		lVar12 = lVar11
		local11b := local167[iVar4]
		lVar12 = mul(uint32(iVar4*iVar4), uint32(iVar4*iVar4>>31), local127.HighPart, local127.HighPart>>32)
		local13f = lVar12.LowPart >> 32
		lVar11 = mul(uint32((local11b)+1), uint32(int(local11b)>>31+i), puVar2.HighPart, uVar5)
		i--
	}

	fmt.Println("Large integer result use ghidra")
	fmt.Printf("%08x  ", puVar2.HighPart)
	fmt.Printf("%08x  ", puVar2.LowPart)
}

// test 测试大整数的基本运算
func test() {
	var div LargeInteger
	div.LowPart = 0x0000091416154ED7 % 0xFFFFFFFF // 模拟低位
	div.HighPart = 0x0000091416154ED7 / 0x6A      // 高位计算

	var mul LargeInteger
	mul.LowPart = 0x00000015ECE7BB69 % 0xFFFFFFFF // 模拟低位
	mul.HighPart = 0x00000015ECE7BB69 * 0x31      // 高位计算

	fmt.Printf("div QuadPart: 0x%016x\n", (uint64(div.HighPart)<<32)|uint64(div.LowPart))
	fmt.Printf("div HighPart: 0x%016x\n", div.HighPart)
	fmt.Printf("div LowPart: 0x%016x\n", div.LowPart)

	fmt.Printf("mul QuadPart: 0x%016x\n", (uint64(mul.HighPart)<<32)|uint64(mul.LowPart))
	fmt.Printf("mul HighPart: 0x%016x\n", mul.HighPart)
	fmt.Printf("mul LowPart: 0x%016x\n", mul.LowPart)
}

// main 函数
func main() {
	test()         // 进行测试
	ghidraDecode() // 调用主要解码函数
}
