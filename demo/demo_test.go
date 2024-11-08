package main

import (
	"encoding/hex"
	"testing"

	"github.com/ddkwork/golibrary/assert"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream/binary"
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
	v5 := 0
	v6 := 0
	v7 := 0
	v8 := 0
	v9 := 0
	index := 0
	i := 0
	v17 := 0
	elem := 0
	v20 := 0
	out := 0

	out = 0x71B793*(int(data[1])+0x1302*(int(data[1])+0xf366)) + 0x7CFF86
	v5 = (int(data[2])) + 0x1634*(int(data[2])+0xf366) + 1
	out = v5*(out) + 2*((out)/0x6381BE9A+(out>>0x12)) + 0x2D1F65
	v6 = out

	// todo 下面的有bug，ida伪代码调试的时候变量的值刷新很诡异,无法对比
	v7 = ((int(data[3]))+0x1968*(int(data[3])+0xf366)+1)*v6 + 3*((v6)/0x6381BE9A+(v6>>0x12)+0x21D78D)
	// 0x423FBDA7D589F776 todo bug

	v17 = v7/0x6381BE9A + (v7 >> 0x12)                            // 0x109099DA1514
	v8 = ((int(data[4])) + 0x1C9E*(int(data[4])+0xf366) + 1) * v7 // 0x5DE07FE3195B295A
	v9 = 4*(v17) + v8 + 0xB47D9D
	index = 7
	out = (4*(v17)+v8+0xB47D9D)&0xFFF0 + (v9 >> 0x13) + v9&0xFFFFFF + v9&0xFFFFFF // 0xBBC1935A1FC
	i = 7
	v20 = 0x18F
	v20 = 0xC90B9788
	out = 0x13530849000010B9
	v6 = 0x1353084910B9
	for { // 先测试第一轮的计算
		elem = int(data[index])                                                  // f0 c2 46 fe 7b 8a 99 09 逆序输入的8字节
		v17 = (out / 0x6A) + (((v6>>32)^((v6>>32)>>7))<<32|((v6)^(v6>>7)))>>0x19 // 0x43F9B1AC1A
		out = elem*elem*elem + ((elem)+1)*(out) + (index * index * (v17))        // 0xB19168E94E0DC
		index = i - 1                                                            // 6
		out += v17/0x14C9 + i*(int(data[i]))*(int(data[i]))*v20                  // 0x3BF80DB8 000FC5F6 todo must
		i--
		v20--
		if i-1 < 0 {
			break
		}
	}
	return binary.LittleEndian.PutUint64(uint64(out))
}
