package gameboy

type Registers struct {
	a uint8
	b uint8
	c uint8
	d uint8
	e uint8
	f FlagsRegister
	h uint8
	l uint8
}

func get_16bit_registers(register_a uint8, register_b uint8) uint16 {
	a := uint16(register_a) << 8
	b := uint16(register_b)

	return a | b
}

func (reg *Registers) get_bc() uint16 {
	return get_16bit_registers(reg.b, reg.c)
}

func (reg *Registers) set_bc(value uint16) {
	reg.b = uint8((value & 0xFF00) >> 8)
	reg.c = uint8((value & 0xFF))
}

func (reg *Registers) get_hl() uint16 {
	return get_16bit_registers(reg.h, reg.l)
}

func (reg *Registers) set_hl(value uint16) {
	reg.h = uint8((value & 0xFF00) >> 8)
	reg.l = uint8((value & 0xFF))
}

// 'f' register is a special register
// The lower four bits are ALWAYS 0s when certain events occur
// The CPU flags certain states. Let's model these:
type FlagsRegister struct {
	zero       bool // set to true, result of operation is 0
	subtract   bool // set to true, operation was subsctraction
	half_carry bool // set to true, the operation resulted in an overflow
	carry      bool // if overflow occurs from lower four bits
}

const ZERO_FLAG_BYTE_POSITION uint8 = 7
const SUBTRACT_FLAG_BYTE_POSITION uint8 = 6
const HALF_CARRY_FLAG_BYTE_POSITION uint8 = 5
const CARRY_FLAG_BYTE_POSITION uint8 = 4

func bool_to_int(b bool) uint8 {
	var int_value uint8 = 0
	if b {
		int_value = 1
	}

	return int_value
}

func (flag *FlagsRegister) to_uint8() uint8 {
	zero := bool_to_int(flag.zero)
	subtract := bool_to_int(flag.subtract)
	half_carry := bool_to_int(flag.half_carry)
	carry := bool_to_int(flag.carry)

	return zero<<ZERO_FLAG_BYTE_POSITION |
		subtract<<SUBTRACT_FLAG_BYTE_POSITION |
		half_carry<<HALF_CARRY_FLAG_BYTE_POSITION |
		carry<<CARRY_FLAG_BYTE_POSITION
}

func byte_to_bool(b uint8, position uint8) bool {
	return ((b >> position) & 0b1) != 0
}

func to_flagregister(bits uint8) FlagsRegister {
	zero := byte_to_bool(bits, ZERO_FLAG_BYTE_POSITION)
	subtract := byte_to_bool(bits, SUBTRACT_FLAG_BYTE_POSITION)
	half_carry := byte_to_bool(bits, HALF_CARRY_FLAG_BYTE_POSITION)
	carry := byte_to_bool(bits, CARRY_FLAG_BYTE_POSITION)

	return FlagsRegister{zero, subtract, half_carry, carry}
}
