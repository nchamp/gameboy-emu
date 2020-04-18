package gameboy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd8(t *testing.T) {
	new_value, is_overflow := Add8(1, 1)

	assert.Equal(t, uint8(2), new_value, "new value should be 2")
	assert.Equal(t, false, is_overflow, "is overflow flag should be false")

	new_value, is_overflow = Add8(254, 10)

	assert.Equal(t, uint8(8), new_value, "new value should be 8")
	assert.Equal(t, true, is_overflow, "is overflow flag should be true")
}

func TestAdd(t *testing.T) {
	cpu := CPU{registers: Registers{a: 10}}
	expected_registers := Registers{a: 10, f: FlagsRegister{half_carry: true}}

	result := cpu.add(10)

	assert.Equal(t, uint8(20), result, "result should return 10")
	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")

	cpu = CPU{registers: Registers{a: 255}}
	expected_registers = Registers{a: 255, f: FlagsRegister{zero: true, carry: true, half_carry: true}}

	result = cpu.add(1)

	assert.Equal(t, uint8(0), result, "result should overflow and return 0")
	assert.Equal(t, expected_registers, cpu.registers, "register state shold be as expected")
}

// TODO: Add coverage for execute