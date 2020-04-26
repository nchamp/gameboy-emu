package gameboy

import (
	"math"
)

type Instruction ArithmeticTarget

const (
	ADD   Instruction = iota
	ADDHL Instruction = iota
	ADC   Instruction = iota
	SUB   Instruction = iota
	SBC   Instruction = iota
	AND   Instruction = iota
	OR    Instruction = iota
	XOR   Instruction = iota
	CP    Instruction = iota
)

type ArithmeticTarget int

const (
	// 8bit targets (single register)
	A ArithmeticTarget = iota
	B ArithmeticTarget = iota
	C ArithmeticTarget = iota
	E ArithmeticTarget = iota
	H ArithmeticTarget = iota
	L ArithmeticTarget = iota
	// 16bit targets (multiple registers)
	HL ArithmeticTarget = iota
)

type CPU struct {
	registers Registers
}

func (cpu *CPU) get_target_arithmetic_value(target ArithmeticTarget) uint8 {
	switch target {
	case A:
		return cpu.registers.a
	case B:
		return cpu.registers.b
	case C:
		return cpu.registers.c
	case E:
		return cpu.registers.e
	case H:
		return cpu.registers.h
	case L:
		return cpu.registers.l
	}

	// TODO: Return err if not recognised
	return 0
}

func (cpu *CPU) get_target_virtual_16bit_arithmetic_value(target ArithmeticTarget) uint16 {
	switch target {
	case HL:
		return cpu.registers.get_hl()
	}

	// TODO: Return err if not recognised
	return 0
}

func (cpu *CPU) execute8(instruction Instruction, value uint8) {
	switch instruction {
	case ADD:
		cpu.registers.a = cpu.add(value)
	case ADC:
		cpu.registers.a = cpu.adc(value)
	case SUB:
		cpu.registers.a = cpu.sub(value)
	case SBC:
		cpu.registers.a = cpu.subc(value)
	case AND:
		cpu.registers.a = cpu.and(value)
	case OR:
		cpu.registers.a = cpu.or(value)
	case XOR:
		cpu.registers.a = cpu.xor(value)
	case CP:
		cpu.compare(value)
	}
	// todo: Support more instructions
}

func (cpu *CPU) execute16(instruction Instruction, value uint16) {
	switch instruction {
	case ADDHL:
		cpu.registers.set_hl(cpu.addhl(value))
	}
}

func (cpu *CPU) execute(instruction Instruction, target ArithmeticTarget) {
	switch target {
	case A:
	case B:
	case C:
	case E:
	case H:
	case L:
		cpu.execute8(instruction, cpu.get_target_arithmetic_value(target))
	case HL:
		cpu.execute16(instruction, cpu.get_target_virtual_16bit_arithmetic_value(target))
	}
}

// go allows integer overflow for performance sake, we need to manually detect it
func add8(left uint8, right uint8) (uint8, bool) {
	overflow := false

	if right > 0 && left > math.MaxInt8-right {
		overflow = true
	}

	return left + right, overflow
}

func add16(left uint16, right uint16) (uint16, bool) {
	overflow := false

	if right > 0 && left > math.MaxInt16-right {
		overflow = true
	}

	return left + right, overflow
}

func sub8(left uint8, right uint8) (uint8, bool) {
	overflow := false

	if right > left {
		overflow = true
	}

	return left - right, overflow
}

func (cpu *CPU) add(value uint8) uint8 {
	new_value, is_overflow := add8(cpu.registers.a, value)

	cpu.registers.f.zero = new_value == 0
	cpu.registers.f.subtract = false
	cpu.registers.f.carry = is_overflow

	cpu.registers.f.half_carry = (cpu.registers.a&0xF)+(value&0xF) > 0xF
	return new_value
}

func (cpu *CPU) addhl(value uint16) uint16 {
	hl := cpu.registers.get_hl()
	new_value, is_overflow := add16(hl, value)

	cpu.registers.f.zero = new_value == 0
	cpu.registers.f.subtract = false
	cpu.registers.f.carry = is_overflow

	// test if we flow over 11th bit. does adding the two numbers cause the 11th bit to flip?
	mask := uint16(0b111_1111_1111) // mask out bits 11-15
	cpu.registers.f.half_carry = (value&mask)+(hl&mask) > mask

	return new_value
}

func (cpu *CPU) adc(value uint8) uint8 {
	additional := 0
	if cpu.registers.f.carry {
		additional = 1
	}

	to_add, is_overflow := add8(value, uint8(additional))

	new_value := cpu.add(to_add)

	cpu.registers.f.carry = cpu.registers.f.carry || is_overflow

	return new_value
}

func (cpu *CPU) sub(value uint8) uint8 {
	new_value, is_overflow := sub8(cpu.registers.a, value)

	cpu.registers.f.zero = new_value == 0
	cpu.registers.f.subtract = true
	cpu.registers.f.carry = is_overflow

	cpu.registers.f.half_carry = (cpu.registers.a & 0xF) < (value & 0xF)
	return new_value
}

func (cpu *CPU) subc(value uint8) uint8 {
	additional := 0
	if cpu.registers.f.carry {
		additional = 1
	}

	to_sub, is_overflow := add8(value, uint8(additional))

	new_value := cpu.sub(to_sub)

	cpu.registers.f.carry = cpu.registers.f.carry || is_overflow

	return new_value
}

func (cpu *CPU) and(value uint8) uint8 {
	result := cpu.registers.a & value

	cpu.registers.f.zero = result == 0
	cpu.registers.f.subtract = false
	cpu.registers.f.carry = false
	cpu.registers.f.half_carry = true
	return result
}

func (cpu *CPU) or(value uint8) uint8 {
	result := cpu.registers.a | value

	cpu.registers.f.zero = result == 0
	cpu.registers.f.subtract = false
	cpu.registers.f.carry = false
	cpu.registers.f.half_carry = false
	return result
}

func (cpu *CPU) xor(value uint8) uint8 {
	result := cpu.registers.a ^ value

	cpu.registers.f.zero = result == 0
	cpu.registers.f.subtract = false
	cpu.registers.f.carry = false
	cpu.registers.f.half_carry = false
	return result
}

func (cpu *CPU) compare(value uint8) {
	cpu.registers.f.zero = cpu.registers.a == value
	cpu.registers.f.subtract = true // compare is considered a subtraction
	cpu.registers.f.carry = cpu.registers.a < value
	cpu.registers.f.half_carry = (cpu.registers.a & 0xF) < (value & 0xF)
}
