package main

import (
	"encoding/hex"
	"testing"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/stretchr/testify/assert"
)

func Test_demo(t *testing.T) {
	data := []byte{0x9, 0x99, 0x8a, 0x7b, 0xfe, 0x46, 0xc2, 0xf0}
	b := demo(data)
	mylog.Struct("demo", b)
	assert.Equal(t, uint64(0xCD2F21A91272EE20), b.Self())
	assert.Equal(t, mylog.Check2(hex.DecodeString("CD2F21A91272EE20")), b.Bytes())
}
