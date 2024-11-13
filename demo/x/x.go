package main

import (
	"testing"

	"github.com/ddkwork/golibrary/assert"
)

func mul(a1, a2 uint64) uint64 {
	return a1 * a2
}

func div(a1, a2 uint64) uint64 {
	return a1 / a2
}

func asm(data []byte, t *testing.T) []byte {
	out := make([]byte, 8)
	var v62 uint64 = 0x71B793
	v5 := uint64(data[1])

	v6 := mul(v5+0xF366, 0x1302)
	v62 = mul(v5+v6, 0x71B793) + 0x7CFF86
	v7 := uint64(data[2])
	v50 := v62

	// 000000000014B426 000000001FA6035F
	assert.Unsigned(t, uint64(0), 0x80C2902F4963B/0x6381BE9A) // 0x14B426 030A40BD

	v8 := div(v62, 0x6381BE9A)
	v9 := mul(v8+v62>>18, 2) // 确保计算的正确性
	v10 := mul(v7+0xF366, 0x1634)
	v11 := mul(v7+v10+1, v50)

	v20 := make([]uint64, 2)
	v20[0] = v11 + v9 + 0x2D1F65

	v14 := uint64(data[3])
	// v41 := v62
	v15 := v62 >> 32

	v16 := div(v11, 0x6381BE9A)
	v17 := mul(v16+v15+0x21D78D, 3)
	v18 := mul(v14+0xF366, 0x1968)
	v19 := mul(v14+v18+1, v20[0])

	//v19&0xffffffff + 0x48F7E51C = 0x8D5D086C
	//v19&0xffffffff00000000 - 0x8421=42407DAF>>18=1F6BE357
	//
	//42407DAF>>18=00001090
	v20[1] = (v19 + v17) >> 32

	v42 := (v19 + v17) << 32
	v20[0] = v19 + v17

	// 012FFA3D  8D5D086C
	// 012FFA41  42407DAF
	// 012FFA45  6381BE9A
	// 012FFA49  00000000
	v22 := div(v42, 0x6381BE9A) // todo bug 002D18C7
	v50 = v22 + v62

	v25 := uint64(data[4]) + 0xF366*0x1C9E + 1
	v26 := mul(v25, uint64(data[4]))
	v27 := mul(v50, 4) + v26 + 0xB47D9D
	v20[0] = v27

	index := 7
	num := 7
	v53 := 0x18F - int(data[7]) // 确保 v53 的值在范围内

	for {
		if num < 0 {
			break
		}
		v29 := v20[1]
		v43 := v20[0]
		v52 := v29 ^ (v29 >> 7)
		v30 := div(v43, 0x6A)

		v51 := uint64(data[index]) // 直接从数据中读取
		v50 = v30 + v52
		v31 := mul(uint64(index)*uint64(index), v30+v52)

		v32 := mul(v51+1, v43) + v31
		i := num
		v20[0] = v51*v51*v51 + v32

		// 使用更新后的计算逻辑
		v34 := i * (int(data[i])) * (int(data[i])) * (int(data[i]) + v53)
		v20[0] += uint64(v34)
		num = i - 1
	}

	// 将填充的数据转换为字节
	for i := 0; i < 2; i++ {
		if i < len(v20) {
			out[i*4+0] = byte(v20[i] & 0xFF)
			out[i*4+1] = byte((v20[i] >> 8) & 0xFF)
			out[i*4+2] = byte((v20[i] >> 16) & 0xFF)
			out[i*4+3] = byte((v20[i] >> 24) & 0xFF)
		}
	}

	return out
}
