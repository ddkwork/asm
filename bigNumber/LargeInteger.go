package main

import (
	"encoding/binary"
	"slices"

	"github.com/ddkwork/golibrary/mylog"
)

type LargeInteger struct {
	Low  uint32
	High uint32
}

func FromUint64(u uint64) LargeInteger {
	return LargeInteger{
		Low:  uint32(u & 0xFFFFFFFF),
		High: uint32(u >> 32),
	}
}

func (l *LargeInteger) Debug() {
	mylog.Struct(l)
}

func (l *LargeInteger) Self() uint64 {
	return uint64(l.High)<<32 | uint64(l.Low)
}

func (l *LargeInteger) Bytes() []byte {
	return slices.Concat(binary.LittleEndian.AppendUint32(nil, l.High), binary.LittleEndian.AppendUint32(nil, l.Low))
}

func mul(low, high, low2, high2 uint32) LargeInteger {
	//if (high2 | high) == 0 {
	//	v := uint64(low) * uint64(low2)
	//	return FromUint64(v)
	//}
	x := LargeInteger{Low: low, High: high}
	y := LargeInteger{Low: low2, High: high2}
	return FromUint64(x.Self() * y.Self())
}

func div(low, high, low2, high2 uint32) LargeInteger { // todo add remainder ?
	x := LargeInteger{Low: low, High: high}
	y := LargeInteger{Low: low2, High: high2}
	return FromUint64(x.Self() / y.Self())

	var v LargeInteger
	var uVar1 uint64
	var temp int64
	uVar3 := low
	uVar8 := high2
	uVar6 := high
	uVar9 := low2
	if high2 == 0 {
		uVar3 = high / low2
		temp = (int64(high)%int64(low2)<<32 | int64(low)) / int64(low2)
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
	v.High = uint32(temp >> 32)
	v.Low = uint32(temp)
	return v
}
