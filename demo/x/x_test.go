package main

import (
	"encoding/hex"
	"testing"

	"github.com/ddkwork/golibrary/assert"
	"github.com/ddkwork/golibrary/mylog"
)

func Test_asm(t *testing.T) {
	data := []byte{9, 0x99, 0x8A, 0x7B, 0xFE, 0x46, 0xC2, 0xF0}
	assert.Equal(t, mylog.Check2(hex.DecodeString("CD2F21A91272EE20")), asm(data, t))
}
