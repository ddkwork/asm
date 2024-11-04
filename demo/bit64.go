package main

import (
	"encoding/binary"
	"slices"

	"github.com/ddkwork/golibrary/mylog"
)

type Bit64 struct {
	Low  uint32
	High uint32
}

func FromUint64(u uint64) Bit64 {
	return Bit64{
		Low:  uint32(u & 0xFFFFFFFF),
		High: uint32(u >> 32),
	}
}

func (l *Bit64) Debug() {
	mylog.Struct(l)
}

func (l *Bit64) Self() uint64 {
	return uint64(l.High)<<32 | uint64(l.Low)
}

func (l *Bit64) Bytes() []byte {
	return slices.Concat(binary.LittleEndian.AppendUint32(nil, l.High), binary.LittleEndian.AppendUint32(nil, l.Low))
}

func mul(xLow, xHigh, yLow, yHigh uint32) Bit64 {
	x := Bit64{Low: xLow, High: xHigh}
	y := Bit64{Low: yLow, High: yHigh}
	return FromUint64(x.Self() * y.Self())
}

func div(xLow, xHigh, yLow, yHigh uint32) Bit64 {
	x := Bit64{Low: xLow, High: xHigh}
	y := Bit64{Low: yLow, High: yHigh}
	return FromUint64(x.Self() / y.Self())
}
