package main

import (
	"encoding/hex"
	"testing"

	"github.com/ddkwork/golibrary/stream/binary"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/stretchr/testify/assert"
)

func Test_demo(t *testing.T) {
	data := []byte{0x9, 0x99, 0x8a, 0x7b, 0xfe, 0x46, 0xc2, 0xf0}
	b := demo(data, t)
	assert.Equal(t, uint64(0xCD2F21A91272EE20), b.Uint64())
	assert.Equal(t, mylog.Check2(hex.DecodeString("CD2F21A91272EE20")), b.Bytes())
	assert.Equal(t, mylog.Check2(hex.DecodeString("CD2F21A91272EE20")), asm(data))
}

func asm(data []byte) []byte {
	v5 := 0
	v6 := 0
	v7 := 0
	v8 := 0
	v9 := 0
	index := 0
	i := 0
	v17 := 0
	v18 := 0
	v20 := 0
	out := 0

	out = 0x71B793*(int(data[1])+0x1302*(int(data[1])+0xF0)) + 0x7CFF86
	v5 = (int(data[2])) + 0x1634*(int(data[2])+0xF0) + 1
	v6 = (v5*(out) + 2*((out)/0x6381BE9A+(out>>0x12)) + 0x2D1F65) >> 0x20
	out = v5*(out) + 2*((out)/0x6381BE9A+(out>>0x12)) + 0x2D1F65
	v6 = out
	v7 = ((int(data[3]))+0x1968*(int(data[3])+0xF0)+1)*(v6) + 3*((v6)/0x6381BE9A+(v6>>0x12)+0x21D78D)
	v17 = v7/0x6381BE9A + (v7 >> 0x12)
	v8 = ((int(data[4])) + 0x1C9E*(int(data[4])+0xF0) + 1) * v7
	v9 = 4*(v17) + v8 + 0xB47D9D
	index = 7
	out = (4*(v17)+v8+0xB47D9D)&0xFFF0 + (v9 >> 0x13) + v9&0xFFFFFF + v9&0xFFFFFF
	i = 7
	v20 = 0x18F - (int(data[7]))
	for {
		v18 = int(data[index])
		v17 = (out / 0x6A) + (((v6>>32)^((v6>>32)>>7))<<32|((v6)^(v6>>7)))>>0x19
		out = v18*v18*v18 + ((v18)+1)*(out) + (index * index * (v17))
		index = i - 1
		out += v17/0x14C9 + (i)*(int(data[i]))*(int(data[i]))*(i+v20)
		i--
		if i-1 < 0 {
			break
		}
	}
	return binary.LittleEndian.PutUint64(uint64(out))
}
