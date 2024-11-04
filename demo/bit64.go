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

func (l *Bit64) ShiftLeft(n uint) Bit64 {
	v := l.Self() << n
	*l = FromUint64(v)
	l.Debug("ShiftLeft")
	return *l
}

func (l *Bit64) ShiftRight(n uint) Bit64 {
	v := l.Self() >> n
	*l = FromUint64(v)
	l.Debug("ShiftRight")
	return *l
}

func (l *Bit64) AddInt(u uint64) Bit64 {
	v := l.Self() + u
	*l = FromUint64(v)
	return *l
}

func (l *Bit64) Add(r Bit64) Bit64 {
	v := l.Self() + r.Self()
	*l = FromUint64(v)
	l.Debug("add")
	return *l
	return Bit64{
		Low:  l.Low + r.Low,
		High: l.High + r.High,
	}
}

func (l *Bit64) Sub(r Bit64) Bit64 {
	v := l.Self() - r.Self()
	*l = FromUint64(v)
	l.Debug("sub")
	return *l
	return Bit64{
		Low:  l.Low - r.Low,
		High: l.High - r.High,
	}
}

func (l *Bit64) And(r Bit64) Bit64 {
	return Bit64{
		Low:  l.Low & r.Low,
		High: l.High & r.High,
	}
}

func (l *Bit64) Or(r Bit64) Bit64 {
	return Bit64{
		Low:  l.Low | r.Low,
		High: l.High | r.High,
	}
}

func (l *Bit64) Xor(r Bit64) Bit64 {
	return Bit64{
		Low:  l.Low ^ r.Low,
		High: l.High ^ r.High,
	}
}

func (l *Bit64) Not() Bit64 {
	return Bit64{
		Low:  ^l.Low,
		High: ^l.High,
	}
}

func (l *Bit64) Debug(title string) {
	mylog.Struct(title, l)
}

func (l *Bit64) Self() uint64 {
	return uint64(l.High)<<32 | uint64(l.Low)
}

func (l *Bit64) Bytes() []byte {
	return slices.Concat(binary.LittleEndian.AppendUint32(nil, l.High), binary.LittleEndian.AppendUint32(nil, l.Low))
}

func mul(xLow, xHigh, yLow, yHigh uint32) (b Bit64) {
	defer b.Debug("mul")
	x := Bit64{Low: xLow, High: xHigh}
	y := Bit64{Low: yLow, High: yHigh}
	return FromUint64(x.Self() * y.Self())
}

func div(xLow, xHigh, yLow, yHigh uint32) (b Bit64) {
	defer b.Debug("div")
	x := Bit64{Low: xLow, High: xHigh}
	y := Bit64{Low: yLow, High: yHigh}
	return FromUint64(x.Self() / y.Self())
}
