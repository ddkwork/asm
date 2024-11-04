package main

import (
	"encoding/binary"
	"slices"
)

type LargeInteger struct {
	LowPart  uint32
	HighPart uint32
}

func FromUint64(u uint64) LargeInteger {
	return LargeInteger{
		LowPart:  uint32(u),
		HighPart: uint32(u >> 32),
	}
}

func (l *LargeInteger) QuadPart() uint64 {
	return uint64(l.HighPart)<<32 | uint64(l.LowPart)
}

func (l *LargeInteger) Bytes() []byte {
	return slices.Concat(binary.LittleEndian.AppendUint32(nil, l.HighPart), binary.LittleEndian.AppendUint32(nil, l.LowPart))
}

func mul(low, high, magic, param4 uint32) LargeInteger {
	var l LargeInteger
	if (param4 | high) == 0 {
		v := uint64(low) * uint64(magic)
		return FromUint64(v)
	}
	l.HighPart = low*magic>>32 + high*magic + low*param4<<32
	l.LowPart = (low * magic) & 0xFFFFFFFF
	return l
}

func div(low, high, magic, param4 uint32) LargeInteger {
	var v LargeInteger
	var uVar1 uint64
	var temp int64
	uVar3 := low
	uVar8 := param4
	uVar6 := high
	uVar9 := magic
	if param4 == 0 {
		uVar3 = high / magic
		temp = (int64(high)%int64(magic)<<32 | int64(low)) / int64(magic)
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
