package gameboy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd8(t *testing.T) {
	new_value, is_overflow := add8(1, 1)

	assert.Equal(t, uint8(2), new_value, "new value should be 2")
	assert.Equal(t, false, is_overflow, "is overflow flag should be false")

	new_value, is_overflow = add8(254, 10)

	assert.Equal(t, uint8(8), new_value, "new value should be 8")
	assert.Equal(t, true, is_overflow, "is overflow flag should be true")
}

func TestAdd(t *testing.T) {
	cpu := CPU{registers: Registers{a: 10}}
	expected_registers := Registers{a: 10, f: FlagsRegister{half_carry: true}}

	result := cpu.add(10)

	assert.Equal(t, uint8(20), result, "result should return 20")
	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")

	cpu = CPU{registers: Registers{a: 255}}
	expected_registers = Registers{a: 255, f: FlagsRegister{zero: true, carry: true, half_carry: true}}

	result = cpu.add(1)

	assert.Equal(t, uint8(0), result, "result should overflow and return 0")
	assert.Equal(t, expected_registers, cpu.registers, "register state shold be as expected")
}

func TestAdd16(t *testing.T) {
	new_value, is_overflow := add16(1, 1)

	assert.Equal(t, uint16(2), new_value, "new value should be 2")
	assert.Equal(t, false, is_overflow, "is overflow flag should be false")

	new_value, is_overflow = add16(65534, 10)

	assert.Equal(t, uint16(8), new_value, "new value should be 8")
	assert.Equal(t, true, is_overflow, "is overflow flag should be true")
}

func TestAddhl(t *testing.T) {
	cpu := CPU{registers: Registers{h: 1, l: 0}}
	expected_registers := Registers{h: 1, l: 0, f: FlagsRegister{}}

	result := cpu.addhl(1)

	assert.Equal(t, uint16(0x101), result, "result should return unisgned int (16 bit) 0x101")
	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")
}

func TestAdc(t *testing.T) {
	cpu := CPU{registers: Registers{a: 10}}
	expected_registers := Registers{a: 10, f: FlagsRegister{half_carry: true}}

	result := cpu.adc(10)

	assert.Equal(t, uint8(20), result, "result should return 20")
	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")

	cpu = CPU{registers: Registers{a: 255}}
	expected_registers = Registers{a: 255, f: FlagsRegister{zero: true, carry: true, half_carry: true}}

	result = cpu.adc(1)

	assert.Equal(t, uint8(0), result, "result should overflow and return 0")
	assert.Equal(t, expected_registers, cpu.registers, "register state shold be as expected")
}
