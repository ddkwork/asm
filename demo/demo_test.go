package main

import (
	"encoding/hex"
	"testing"

	"encoding/binary"
	"github.com/ddkwork/golibrary/assert"
	"github.com/ddkwork/golibrary/mylog"
)

func Test_demo(t *testing.T) {
	data := []byte{0x9, 0x99, 0x8a, 0x7b, 0xfe, 0x46, 0xc2, 0xf0}
	b := demo(data, t)
	assert.Equal(t, 0xCD2F21A91272EE20, b.Uint64())
	assert.Equal(t, mylog.Check2(hex.DecodeString("CD2F21A91272EE20")), b.Bytes())
}

func TestName(t *testing.T) {
	data := []byte{0x9, 0x99, 0x8a, 0x7b, 0xfe, 0x46, 0xc2, 0xf0}
	assert.Equal(t, mylog.Check2(hex.DecodeString("CD2F21A91272EE20")), asm(data))
}

func asm(data []byte) []byte {
	var (
		v5  uint64
		v6  uint64
		v7  uint64
		v8  uint64
		v9  uint64
		v17 uint64
		v20 uint64
		out uint64
	)
	elem := uint64(data[1])
	out = 0x71B793*(elem+0x1302*(elem+0xf366)) + 0x7CFF86
	elem = uint64(data[2])
	v5 = elem + 0x1634*(elem+0xf366) + 1
	out = v5*out + 2*(out/0x6381BE9A+(out>>0x12)) + 0x2D1F65
	v6 = out
	elem = uint64(data[3])
	v7 = (elem+
		0x1968*(elem+0xf366)+1)*v6 +
		3*(v6/0x6381BE9A+(v6>>0x12)+0x21D78D)
	v17 = v7/0x6381BE9A + (v7 >> 0x12) // 0x109099DA1514
	elem = uint64(data[4])
	v8 = (elem + 0x1C9E*(elem+0xf366) + 1) * v7 // 0x5DE07FE3195B295A
	v9 = 4*(v17) + v8 + 0xB47D9D
	index := uint64(7)
	out = (4*v17+v8+0xB47D9D)&0xFFF0 + (v9 >> 0x13) + v9&0xFFFFFF + v9&0xFFFFFF // 0xBBC1935A1FC
	i := uint64(7)
	v20 = 0x18F
	out = 0x000010B913530849
	v6 = 0x000010B913530849
	for { // 先测试第一轮的计算
		if i < 1 {
			break
		}
		println(index)
		elem = uint64(data[index])                                             // f0 c2 46 fe 7b 8a 99 09 逆序输入的8字节
		v17 = (out / 0x6A) + (((v6>>32)^((v6>>32)>>7))<<32|(v6^(v6>>7)))>>0x19 // 0x43F9B1AC1A
		// 16FBA94F 第二轮就错了
		out = elem*elem*elem + (elem+1)*out + (index * index * v17)   // 0xB19168E94E0DC
		index = i - 1                                                 // 6
		out += v17/0x14C9 + i*(uint64(data[i]))*(uint64(data[i]))*v20 // 0x3BF80DB8 000FC5F6 todo must
		i--
		v20--
	}
	return binary.LittleEndian.PutUint64(uint64(out))
}
