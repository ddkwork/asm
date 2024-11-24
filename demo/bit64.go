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

func (l *Bit64) Uint64() uint64 {
	return uint64(l.High)<<32 | uint64(l.Low)
}

func (l *Bit64) Bytes() []byte {
	return slices.Concat(binary.LittleEndian.AppendUint32(nil, l.High), binary.LittleEndian.AppendUint32(nil, l.Low))
}

func (l *Bit64) ShiftLeft(n uint) Bit64 {
	v := l.Uint64() << n
	b := FromUint64(v)
	b.Debug("<<")
	return b
}

func (l *Bit64) ShiftRight(n uint) Bit64 {
	v := l.Uint64() >> n
	b := FromUint64(v)
	b.Debug(">>")
	return b
}

func (l *Bit64) AddInt(u uint64) Bit64 {
	v := l.Uint64() + u
	*l = FromUint64(v)
	return *l
}

func (l *Bit64) Add(r Bit64) Bit64 {
	v := l.Uint64() + r.Uint64()
	*l = FromUint64(v)
	l.Debug("+")
	return *l
	return Bit64{
		Low:  l.Low + r.Low,
		High: l.High + r.High,
	}
}

func (l *Bit64) Sub(r Bit64) Bit64 {
	v := l.Uint64() - r.Uint64()
	*l = FromUint64(v)
	l.Debug("-")
	return *l
	return Bit64{
		Low:  l.Low - r.Low,
		High: l.High - r.High,
	}
}

func (l *Bit64) And(r Bit64) (b Bit64) {
	defer b.Debug("&")
	return Bit64{
		Low:  l.Low & r.Low,
		High: l.High & r.High,
	}
}

func (l *Bit64) Or(r Bit64) (b Bit64) {
	defer b.Debug("|")
	return Bit64{
		Low:  l.Low | r.Low,
		High: l.High | r.High,
	}
}

func (l *Bit64) Xor(r Bit64) (b Bit64) {
	defer b.Debug("^")
	return Bit64{
		Low:  l.Low ^ r.Low,
		High: l.High ^ r.High,
	}
}

func (l *Bit64) Not() (b Bit64) {
	defer b.Debug("~")
	return Bit64{
		Low:  ^l.Low,
		High: ^l.High,
	}
}

func (l *Bit64) Rem(r Bit64) (b Bit64) {
	defer b.Debug("%")
	return Bit64{
		Low:  l.Low % r.Low,
		High: l.High % r.High,
	}
}

func (l *Bit64) Debug(title string) {
	mylog.Struct(title, l)
}

var (
	mulCount byte
	divCount byte
)

func mul(xLow, xHigh, yLow, yHigh uint32) (b Bit64) {
	mulCount++
	x := Bit64{Low: xLow, High: xHigh}
	y := Bit64{Low: yLow, High: yHigh}
	type MulInfo struct {
		Index byte
		XLow  uint32
		XHigh uint32
		YLow  uint32
		YHigh uint32
		X     uint64
		Y     uint64
		Z     uint64
	}
	defer func() {
		mylog.Struct("*", MulInfo{
			Index: mulCount,
			XLow:  x.Low,
			XHigh: x.High,
			YLow:  y.Low,
			X:     x.Uint64(),
			Y:     y.Uint64(),
			Z:     b.Uint64(),
		})
	}()
	return FromUint64(x.Uint64() * y.Uint64())
}

func div(xLow, xHigh, yLow, yHigh uint32) (b Bit64) {
	divCount++
	x := Bit64{Low: xLow, High: xHigh}
	y := Bit64{Low: yLow, High: yHigh}
	type DivInfo struct {
		Index byte
		XLow  uint32
		XHigh uint32
		YLow  uint32
		YHigh uint32
		X     uint64
		Y     uint64
		Z     uint64
	}
	defer func() {
		mylog.Struct("/", DivInfo{
			Index: divCount,
			XLow:  x.Low,
			XHigh: x.High,
			YLow:  y.Low,
			YHigh: y.High,
			X:     x.Uint64(),
			Y:     y.Uint64(),
			Z:     b.Uint64(),
		})
	}()
	return FromUint64(x.Uint64() / y.Uint64())
}
