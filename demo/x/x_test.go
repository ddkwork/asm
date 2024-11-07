package main

import (
	"encoding/hex"
	"testing"

	"github.com/ddkwork/golibrary/assert"
	"github.com/ddkwork/golibrary/mylog"
)

func Test_asm(t *testing.T) {
	assert.Equal(t, mylog.Check2(hex.DecodeString("CD2F21A91272EE20")), asm())
}
