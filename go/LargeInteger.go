package main

type LargeInteger struct {
	LowPart  uint32
	HighPart uint32
}

func (li LargeInteger) QuadPart() uint64 {
	return uint64(li.HighPart)<<32 | uint64(li.LowPart)
}

func makeLargeInteger(high, low uint32) LargeInteger {
	return LargeInteger{LowPart: low, HighPart: high}
}

func mul(param1, param2, param3, param4 uint32) LargeInteger {
	var v LargeInteger

	if (param4 | param2) == 0 {
		v.LowPart = param1 * param3
		return v
	}
	v.HighPart = param1*param3>>32 + param2*param3 + param1*param4<<32
	v.LowPart = (param1 * param3) & 0xFFFFFFFF
	return v
}

func div(param1, param2, param3, param4 uint32) LargeInteger {
	var v LargeInteger
	var uVar1 uint64
	var temp int64
	uVar3 := param1
	uVar8 := param4
	uVar6 := param2
	uVar9 := param3

	if param4 == 0 {
		uVar3 = param2 / param3
		temp = (int64(param2)%int64(param3)<<32 | int64(param1)) / int64(param3)
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
