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
	x := LargeInteger{Low: low, High: high}
	y := LargeInteger{Low: low2, High: high2}
	return FromUint64(x.Self() * y.Self())
}

func div(low, high, low2, high2 uint32) LargeInteger { // todo add remainder ?
	x := LargeInteger{Low: low, High: high}
	y := LargeInteger{Low: low2, High: high2}
	return FromUint64(x.Self() / y.Self())
}
