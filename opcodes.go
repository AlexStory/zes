package main

type OpCode struct {
	Code   byte
	Name   string
	Length byte
	Cycles byte
	Mode   AddressingMode
}

func new_opcode(
	code byte,
	name string,
	length byte,
	cycles byte,
	mode AddressingMode,
) *OpCode {
	return &OpCode{
		Code:   code,
		Name:   name,
		Length: length,
		Cycles: cycles,
		Mode:   mode,
	}
}

var OpCodes = [...]*OpCode{
	new_opcode(0x00, "BRK", 1, 7, NONE),
	new_opcode(0xEA, "NOP", 1, 2, NONE),
	new_opcode(0xAA, "TAX", 1, 2, NONE),
	new_opcode(0xA8, "TAY", 1, 2, NONE),
	new_opcode(0xE8, "INX", 1, 2, NONE),
	new_opcode(0xC8, "INY", 1, 2, NONE),
	new_opcode(0xCA, "DEX", 1, 2, NONE),
	new_opcode(0x88, "DEY", 1, 2, NONE),
	new_opcode(0x48, "PHA", 1, 3, NONE),
	new_opcode(0x08, "PHP", 1, 3, NONE),
	new_opcode(0x68, "PLA", 1, 4, NONE),
	new_opcode(0x28, "PLP", 1, 4, NONE),

	new_opcode(0x69, "ADC", 2, 2, IMMEDIATE),
	new_opcode(0x65, "ADC", 2, 3, ZEROPAGE),
	new_opcode(0x75, "ADC", 2, 4, ZEROPAGE_X),
	new_opcode(0x6D, "ADC", 3, 4, ABSOLUTE),
	new_opcode(0x7D, "ADC", 3, 4, ABSOLUTE_X),
	new_opcode(0x79, "ADC", 3, 4, ABSOLUTE_Y),
	new_opcode(0x61, "ADC", 2, 6, INDIRECT_X),
	new_opcode(0x71, "ADC", 2, 5, INDIRECT_Y),

	new_opcode(0x29, "AND", 2, 2, IMMEDIATE),
	new_opcode(0x25, "AND", 2, 3, ZEROPAGE),
	new_opcode(0x35, "AND", 2, 4, ZEROPAGE_X),
	new_opcode(0x2D, "AND", 3, 4, ABSOLUTE),
	new_opcode(0x3D, "AND", 3, 4, ABSOLUTE_X),
	new_opcode(0x39, "AND", 3, 4, ABSOLUTE_Y),
	new_opcode(0x21, "AND", 2, 6, INDIRECT_X),
	new_opcode(0x31, "AND", 2, 5, INDIRECT_Y),

	new_opcode(0x0A, "ASL", 1, 2, NONE),
	new_opcode(0x06, "ASL", 2, 5, ZEROPAGE),
	new_opcode(0x16, "ASL", 2, 6, ZEROPAGE_X),
	new_opcode(0x0E, "ASL", 3, 6, ABSOLUTE),
	new_opcode(0x1E, "ASL", 3, 7, ABSOLUTE_X),

	new_opcode(0xC6, "DEC", 2, 5, ZEROPAGE),
	new_opcode(0xD6, "DEC", 2, 6, ZEROPAGE_X),
	new_opcode(0xCE, "DEC", 3, 6, ABSOLUTE),
	new_opcode(0xDE, "DEC", 3, 7, ABSOLUTE_X),

	new_opcode(0xE6, "INC", 2, 5, ZEROPAGE),
	new_opcode(0xF6, "INC", 2, 6, ZEROPAGE_X),
	new_opcode(0xEE, "INC", 3, 6, ABSOLUTE),
	new_opcode(0xFE, "INC", 3, 7, ABSOLUTE_X),

	new_opcode(0x4C, "JMP", 3, 3, ABSOLUTE),
	new_opcode(0x6C, "JMP", 3, 5, NONE),

	new_opcode(0xA9, "LDA", 2, 2, IMMEDIATE),
	new_opcode(0xA5, "LDA", 2, 3, ZEROPAGE),
	new_opcode(0xB5, "LDA", 2, 4, ZEROPAGE_X),
	new_opcode(0xAD, "LDA", 3, 4, ABSOLUTE),
	new_opcode(0xBD, "LDA", 3, 4, ABSOLUTE_X),
	new_opcode(0xB9, "LDA", 3, 4, ABSOLUTE_Y),
	new_opcode(0xA1, "LDA", 2, 6, INDIRECT_X),
	new_opcode(0xB1, "LDA", 2, 5, INDIRECT_Y),

	new_opcode(0xA2, "LDX", 2, 2, IMMEDIATE),
	new_opcode(0xA6, "LDX", 2, 3, ZEROPAGE),
	new_opcode(0xB6, "LDX", 2, 4, ZEROPAGE_Y),
	new_opcode(0xAE, "LDX", 3, 4, ABSOLUTE),
	new_opcode(0xBE, "LDX", 3, 4, ABSOLUTE_Y),

	new_opcode(0xA0, "LDY", 2, 2, IMMEDIATE),
	new_opcode(0xA4, "LDY", 2, 3, ZEROPAGE),
	new_opcode(0xB4, "LDY", 2, 4, ZEROPAGE_X),
	new_opcode(0xAC, "LDY", 3, 4, ABSOLUTE),
	new_opcode(0xBC, "LDY", 3, 4, ABSOLUTE_X),

	new_opcode(0x85, "STA", 2, 3, ZEROPAGE),
	new_opcode(0x95, "STA", 2, 4, ZEROPAGE_X),
	new_opcode(0x8D, "STA", 3, 4, ABSOLUTE),
	new_opcode(0x9D, "STA", 3, 5, ABSOLUTE_X),
	new_opcode(0x99, "STA", 3, 5, ABSOLUTE_Y),
	new_opcode(0x81, "STA", 2, 6, INDIRECT_X),
	new_opcode(0x91, "STA", 2, 6, INDIRECT_Y),

	new_opcode(0x86, "STX", 2, 3, ZEROPAGE),
	new_opcode(0x96, "STX", 2, 4, ZEROPAGE_Y),
	new_opcode(0x8E, "STX", 3, 4, ABSOLUTE),

	new_opcode(0x84, "STY", 2, 3, ZEROPAGE),
	new_opcode(0x94, "STY", 2, 4, ZEROPAGE_X),
	new_opcode(0x8C, "STY", 3, 4, ABSOLUTE),
}

var OpCodeMap = map[byte]*OpCode{}

func init() {
	for _, opcode := range OpCodes {
		OpCodeMap[opcode.Code] = opcode
	}
}
