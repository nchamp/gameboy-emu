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
	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")
}

func TestAdd16(t *testing.T) {
	new_value, is_overflow := add16(1, 1)

	assert.Equal(t, uint16(2), new_value, "new value should be 2")
	assert.Equal(t, false, is_overflow, "is overflow flag should be false")

	new_value, is_overflow = add16(65534, 10)

	assert.Equal(t, uint16(8), new_value, "new value should be 8")
	assert.Equal(t, true, is_overflow, "is overflow flag should be true")
}

func TestSub8(t *testing.T) {
	new_value, is_overflow := sub8(1, 1)

	assert.Equal(t, uint8(0), new_value, "new value should be 2")
	assert.Equal(t, false, is_overflow, "is overflow flag should be false")

	new_value, is_overflow = sub8(1, 2)

	assert.Equal(t, uint8(255), new_value, "new value should be 255")
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
	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")
}

func TestSub(t *testing.T) {
	cpu := CPU{registers: Registers{a: 20}}
	expected_registers := Registers{a: 20, f: FlagsRegister{subtract: true, half_carry: true}}

	result := cpu.sub(10)

	assert.Equal(t, uint8(10), result, "result should return 10")
	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")

	cpu = CPU{registers: Registers{a: 255}}
	expected_registers = Registers{a: 255, f: FlagsRegister{subtract: true, zero: true}}

	result = cpu.sub(255)

	assert.Equal(t, uint8(0), result, "result should return 0")
	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")

	cpu = CPU{registers: Registers{a: 2}}
	expected_registers = Registers{a: 2, f: FlagsRegister{subtract: true, carry: true, half_carry: true}}

	result = cpu.sub(5)

	assert.Equal(t, uint8(253), result, "result should return 253")
	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")
}

func TestSubc(t *testing.T) {
	cpu := CPU{registers: Registers{a: 20}}
	expected_registers := Registers{a: 20, f: FlagsRegister{subtract: true, half_carry: true}}

	result := cpu.subc(10)

	assert.Equal(t, uint8(10), result, "result should return 10")
	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")

	cpu = CPU{registers: Registers{a: 255}}
	expected_registers = Registers{a: 255, f: FlagsRegister{subtract: true, zero: true}}

	result = cpu.subc(255)

	assert.Equal(t, uint8(0), result, "result should return 0")
	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")

	cpu = CPU{registers: Registers{a: 255, f: FlagsRegister{carry: true}}}
	expected_registers = Registers{a: 255, f: FlagsRegister{subtract: true, zero: true, carry: true}}

	result = cpu.subc(254)

	assert.Equal(t, uint8(0), result, "result should return 0")
	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")

	cpu = CPU{registers: Registers{a: 2}}
	expected_registers = Registers{a: 2, f: FlagsRegister{subtract: true, carry: true, half_carry: true}}

	result = cpu.subc(5)

	assert.Equal(t, uint8(253), result, "result should return 253")
	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")
}

func TestAnd(t *testing.T) {
	cpu := CPU{registers: Registers{a: 20}}
	expected_registers := Registers{a: 20, f: FlagsRegister{half_carry: true}}

	result := cpu.and(20)

	assert.Equal(t, uint8(20), result, "result should return 20")
	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")

	cpu = CPU{registers: Registers{a: 20}}
	expected_registers = Registers{a: 20, f: FlagsRegister{zero: true, half_carry: true}}

	result = cpu.and(10)

	assert.Equal(t, uint8(0), result, "result should return 0")
	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")
}

func TestOr(t *testing.T) {
	cpu := CPU{registers: Registers{a: 20}}
	expected_registers := Registers{a: 20}

	result := cpu.or(20)

	assert.Equal(t, uint8(20), result, "result should return 20")
	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")

	cpu = CPU{registers: Registers{a: 20}}
	expected_registers = Registers{a: 20}

	result = cpu.or(10)

	assert.Equal(t, uint8(30), result, "result should return 30")
	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")
}

func TestXor(t *testing.T) {
	cpu := CPU{registers: Registers{a: 20}}
	expected_registers := Registers{a: 20, f: FlagsRegister{zero: true}}

	result := cpu.xor(20)

	assert.Equal(t, uint8(0), result, "result should return 0")
	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")

	cpu = CPU{registers: Registers{a: 20}}
	expected_registers = Registers{a: 20}

	result = cpu.xor(10)

	assert.Equal(t, uint8(30), result, "result should return 30")
	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")
}

func TestCompare(t *testing.T) {
	cpu := CPU{registers: Registers{a: 20}}
	expected_registers := Registers{a: 20, f: FlagsRegister{zero: true, subtract: true}}

	cpu.compare(20)

	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")

	cpu = CPU{registers: Registers{a: 20}}
	expected_registers = Registers{a: 20, f: FlagsRegister{subtract: true, half_carry: true}}

	cpu.compare(10)

	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")

	cpu = CPU{registers: Registers{a: 20}}
	expected_registers = Registers{a: 20, f: FlagsRegister{subtract: true, half_carry: true, carry: true}}

	cpu.compare(30)

	assert.Equal(t, expected_registers, cpu.registers, "register state should be as expected")
}

// TODO: Add tests for: INC, DEC and CCF
