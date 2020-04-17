package gameboy

import (
	"math"
)

type Instruction ArithmeticTarget

const (
	ADD   Instruction = iota
	ADDHL Instruction = iota
)

type ArithmeticTarget int

const (
	A ArithmeticTarget = iota
	B ArithmeticTarget = iota
	C ArithmeticTarget = iota
	E ArithmeticTarget = iota
	H ArithmeticTarget = iota
	L ArithmeticTarget = iota
)

type CPU struct {
	registers Registers
}

func (cpu *CPU) execute(instruction Instruction, target ArithmeticTarget) {
	switch instruction {
	case ADD:
		switch target {
		case C:
			cpu.registers.a = cpu.add(cpu.registers.c)
		}
		// todo: Support more targets
	case ADDHL:
		switch target {

		}
	}
	// todo: Support more instructions
}

// go allows integer overflow for performance sake, we need to manually detect it
func Add8(left uint8, right uint8) (uint8, bool) {
	overflow := false

	if right > 0 && left > math.MaxInt8-right {
		overflow = true
	}

	return left + right, overflow
}

func (cpu *CPU) add(value uint8) uint8 {
	new_value, is_overflow := Add8(cpu.registers.a, value)

	cpu.registers.f.zero = new_value == 0
	cpu.registers.f.subtract = false
	cpu.registers.f.carry = is_overflow

	cpu.registers.f.half_carry = (cpu.registers.a&0xF)+(value&0xF) > 0xF
	return new_value
}
