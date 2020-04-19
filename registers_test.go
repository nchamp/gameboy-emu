package gameboy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet_bc(t *testing.T) {
	registers := Registers{b: 10, c: 10}

	result := registers.get_bc()

	assert.Equal(t, uint16(2570), result, "they should match")
}

func TestSet_bc(t *testing.T) {
	registers := Registers{}

	registers.set_bc(10)

	assert.Equal(t, uint8(0), registers.b, "they should match")
	assert.Equal(t, uint8(10), registers.c, "they should match")
}

func TestBool_to_int(t *testing.T) {
	assert.Equal(t, uint8(0), bool_to_int(false), "they should match")
	assert.Equal(t, uint8(1), bool_to_int(true), "they should match")
}

func TestByte_to_bool(t *testing.T) {
	assert.Equal(t, false, byte_to_bool(1, 1), "they should match")
	assert.Equal(t, true, byte_to_bool(1, 0), "they should match")
}

// TODO: to_flagregister && to_uint8
