package main

import (
	"testing"

	"github.com/ddkwork/golibrary/mylog"

	"github.com/ddkwork/golibrary/assert"
)

func mul(param1, param2, param3, param4 uint64) uint64 {
	x := param1 | param2<<32
	y := param3 | param4<<32
	z := x * y
	return z
}

func div(param1, param2, param3, param4 uint64) uint64 {
	x := param1 | param2<<32
	y := param3 | param4<<32
	return x / y
}

func asm1(data []byte, t *testing.T) uint64 {
	var (
		uVar1    uint64
		lVar2    uint64
		highXor  uint64
		v        uint64
		elem     uint64
		uVar6    uint64
		uVar10   uint64
		iVar11   uint64
		uVar12   uint64
		out      uint64
		uStack30 uint64
	)

	out = 0
	uVar10 = mul((uint64(data[1]))+0xf366, 0, 0x1302, 0)
	lVar2 = uVar10 + uint64(data[1])
	uVar10 = mul(lVar2, lVar2>>32, 0x71b793, 0)
	out = uVar10 + 0x7cff86
	highXor = uint64(data[2])
	uVar1 = out >> 18
	uVar6 = out
	elem = out >> 32
	iVar11 = div(uVar6, elem, 0x6381be9a, 0)
	lVar2 = uVar1 + iVar11
	uVar10 = mul(lVar2, lVar2>>32, 2, 0)
	uVar12 = mul(highXor+0xf366, 0, 0x1634, 0)
	v = uVar12 + highXor

	fnBool2Uint64 := func(b bool) uint64 {
		if b {
			return 1
		}
		return 0
	}

	uVar12 = mul(v+1,
		//(uVar12>>32)+(uVar12<<32|highXor)+(v),
		//(uVar12>>32)+(uVar12<<32|highXor)+fnBool2Uint64(0xfffffffe < v),
		fnBool2Uint64(0xfffffffe < v),
		uVar6&0xffffffff,
		// uVar6&0xffffffff|uVar1<<18,
		elem)
	lVar2 = uVar12 + uVar10 + 0x2d1f65
	highXor = lVar2 >> 32
	puVar9 := out
	uVar6 = uint64(data[3])
	v = out >> 18
	//$ ==>     32B36B74
	//$+4       B0254C17
	//$+8       6381BE9A
	//$+C       00000000
	puVar9 = lVar2 & 0xffffffff
	iVar11 = div(puVar9, highXor, 0x6381be9a, 0)
	lVar2 >>= 18
	lVar2 += iVar11 + 0x21d78d // & 0xffffffff

	// lVar2 = iVar11 | (lVar2>>32)<<32 | (highXor << 18) + ((highXor>>18)<<32 | v) | (highXor << 14) + 0x21d78d

	//$ ==>     1852A1B4
	//$+4       00002C0B
	//$+8       00000003
	//$+C       00000000
	uVar10 = mul(lVar2&0xffffffff, lVar2>>32, 3, 0)
	uVar12 = mul(uVar6+0xf366, 0, 0x1968, 0)

	v = uVar12 + uVar6
	//$ ==>     18340CE4
	//$+4       00000000
	//$+8       32B36B74
	//$+C       B0254C17
	uVar12 = mul(v+1,
		//(uVar12>>32)+(uVar12<<32|uVar6)+(v),
		fnBool2Uint64(0xfffffffe < v),
		//(uVar12>>32)+(uVar12<<32|uVar6)+(0xfffffffe < v),
		puVar9, highXor)
	uVar1 = uVar12 + uVar10
	puVar9 = uVar1
	v = uVar1 >> 32
	uStack30 = v
	iVar11 = div(puVar9, v, 0x6381be9a, 0)
	lVar2 = (uVar1 >> 18) + iVar11
	highXor = uint64(data[4])
	uVar10 = mul(highXor+0xf366, 0, 0x1c9e, 0)

	v = uVar10 + highXor
	//$ ==>     1B51C6B7
	//$+4       00000000
	//$+8       8D5D086C
	//$+C       42407DAF
	puVar9 = uVar1 & 0xffffffff
	uVar10 = mul(v+1,
		//(uVar10>>32)+(uVar10<<32|highXor)+(0xfffffffe < v),
		fnBool2Uint64(0xfffffffe < v),
		//(uVar10>>32)+(uVar10<<32|highXor)+(v),
		puVar9, uStack30)
	uVar12 = mul(lVar2, lVar2>>32, 4, 0)

	// todo 检查下面的计算即可成功
	uVar1 = uVar12 + uVar10 + 0xb47d9d
	v = uVar1 >> 32
	out = (uVar1 >> 19) + ((v>>16)<<32|v)<<32 | ((v >> 16) + v) + (uVar1 & 0xfff0)

	out = 0x000010B913530849
	uStack30 = 0x000010B913530849
	v = uStack30
	j := uint64(0x18f)
	for i := uint64(7); i >= 0; i-- {
		highXor = uStack30>>7 ^ uStack30
		elem = uStack30 ^ out
		uVar6 = elem >> 25
		//$ ==>     13530849
		//$+4       000010B9
		//$+8       0000006A
		//$+C       00000000

		//$ ==>     3BF80DB8
		//$+4       000FC5F6
		//$+8       0000006A
		//$+C       00000000
		iVar11 = div(v&0xffffffff, v>>32, 0x6a, 0)

		// lVar2 = iVar11 + uVar6<<32 |
		//	elem>>25 | highXor<<7
		//edx=D78FFDA3
		//ecx=FDA7D
		//.text:00BC19C0 asm.exe:$19C0 #DC0

		//ecx=7ED3EEB
		//eax=F0E6A64
		//.text:00BC19DF asm.exe:$19DF #DDF
		highXor >>= 25                 // todo bug 第二轮错了
		highXor += iVar11 & 0xffffffff //+ 7ED3EEB ? 应该和 uVar6接近
		elem = uint64(data[i])
		//$ ==>     00000031
		//$+4       00000000
		//$+8       633BCC44
		//$+C       00000028

		//ecx=7ED3EEB
		//eax=F0E6A64
		//.text:00BC19DF asm.exe:$19DF #DDF
		//$ ==>     00000024
		//$+4       00000000
		//$+8       16FBA94F
		//$+C       00002618
		//uVar11 = mul(pvVar4,i * i,i * i >> 0x1f,(uint)pvVar4,highXor);
		uVar10 = mul(i*i, i*i>>31, highXor, iVar11>>32)
		// 3B63CF1C
		// 0x55B62 1E06F610
		// uVar10 = mul(i*i, i*i>>31, uVar6, highXor)

		//$ ==>     000000F1
		//$+4       00000000
		//$+8       13530849
		//$+C       000010B9
		out = v & 0xffffffff
		uVar12 = mul(elem+1, fnBool2Uint64(0xfffffffe < elem), out, v>>32)

		// uVar12 = mul(elem+1, 0xfffffffe < elem, out, v)
		out = uVar10 + uVar12 + (elem * elem * elem)
		v = uint64(data[i])

		//$ ==>     633BCC44
		//$+4       00000028
		//$+8       000014C9
		//$+C       00000000
		iVar11 = div(highXor, iVar11>>32, 0x14c9, 0)
		uStack30 = iVar11
		// out = iVar11 + (uint64(data[i]) + ((0x18f)-uint64(data[7]))*v*v*i) + out
		mylog.Hex(j, j)
		// out = iVar11 + (uint64(data[i]) + j*v*v*i) + out
		out += j*v*v*i + iVar11
		v = out
		j--
		if i == 7 {
			assert.UnsignedInteger(t, 0x01F16EFB, iVar11)
			assert.UnsignedInteger(t, 0x000FC5F63BF80DB8, out)
			assert.Equal(t, 0x000FC5F63BF80DB8, out)
		}
		if i == 0 {
			break
		}
	}
	return out // 0xba0529b4000535f8
}
