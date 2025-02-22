package main

type OpCode struct {
	Code   byte
	Name   string
	Length byte
	Cycles byte
	Mode   AddressingMode
}

func newOpcode(
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
	newOpcode(0x00, "BRK", 1, 7, NONE),
	newOpcode(0xEA, "NOP", 1, 2, NONE),
	newOpcode(0xAA, "TAX", 1, 2, NONE),
	newOpcode(0xA8, "TAY", 1, 2, NONE),
	newOpcode(0xE8, "INX", 1, 2, NONE),
	newOpcode(0xC8, "INY", 1, 2, NONE),
	newOpcode(0xCA, "DEX", 1, 2, NONE),
	newOpcode(0x88, "DEY", 1, 2, NONE),
	newOpcode(0x48, "PHA", 1, 3, NONE),
	newOpcode(0x08, "PHP", 1, 3, NONE),
	newOpcode(0x68, "PLA", 1, 4, NONE),
	newOpcode(0x28, "PLP", 1, 4, NONE),
	newOpcode(0x18, "CLC", 1, 2, NONE),
	newOpcode(0xD8, "CLD", 1, 2, NONE),
	newOpcode(0x58, "CLI", 1, 2, NONE),
	newOpcode(0xB8, "CLV", 1, 2, NONE),

	newOpcode(0x69, "ADC", 2, 2, IMMEDIATE),
	newOpcode(0x65, "ADC", 2, 3, ZEROPAGE),
	newOpcode(0x75, "ADC", 2, 4, ZEROPAGE_X),
	newOpcode(0x6D, "ADC", 3, 4, ABSOLUTE),
	newOpcode(0x7D, "ADC", 3, 4, ABSOLUTE_X),
	newOpcode(0x79, "ADC", 3, 4, ABSOLUTE_Y),
	newOpcode(0x61, "ADC", 2, 6, INDIRECT_X),
	newOpcode(0x71, "ADC", 2, 5, INDIRECT_Y),

	newOpcode(0x29, "AND", 2, 2, IMMEDIATE),
	newOpcode(0x25, "AND", 2, 3, ZEROPAGE),
	newOpcode(0x35, "AND", 2, 4, ZEROPAGE_X),
	newOpcode(0x2D, "AND", 3, 4, ABSOLUTE),
	newOpcode(0x3D, "AND", 3, 4, ABSOLUTE_X),
	newOpcode(0x39, "AND", 3, 4, ABSOLUTE_Y),
	newOpcode(0x21, "AND", 2, 6, INDIRECT_X),
	newOpcode(0x31, "AND", 2, 5, INDIRECT_Y),

	newOpcode(0x0A, "ASL", 1, 2, NONE),
	newOpcode(0x06, "ASL", 2, 5, ZEROPAGE),
	newOpcode(0x16, "ASL", 2, 6, ZEROPAGE_X),
	newOpcode(0x0E, "ASL", 3, 6, ABSOLUTE),
	newOpcode(0x1E, "ASL", 3, 7, ABSOLUTE_X),

	newOpcode(0x90, "BCC", 2, 2, NONE),
	newOpcode(0xB0, "BCS", 2, 2, NONE),
	newOpcode(0xF0, "BEQ", 2, 2, NONE),
	newOpcode(0x30, "BMI", 2, 2, NONE),
	newOpcode(0xD0, "BNE", 2, 2, NONE),
	newOpcode(0x10, "BPL", 2, 2, NONE),
	newOpcode(0x50, "BVC", 2, 2, NONE),
	newOpcode(0x70, "BVS", 2, 2, NONE),

	newOpcode(0x24, "BIT", 2, 3, ZEROPAGE),
	newOpcode(0x2C, "BIT", 3, 4, ABSOLUTE),

	newOpcode(0xC9, "CMP", 2, 2, IMMEDIATE),
	newOpcode(0xC5, "CMP", 2, 3, ZEROPAGE),
	newOpcode(0xD5, "CMP", 2, 4, ZEROPAGE_X),
	newOpcode(0xCD, "CMP", 3, 4, ABSOLUTE),
	newOpcode(0xDD, "CMP", 3, 4, ABSOLUTE_X),
	newOpcode(0xD9, "CMP", 3, 4, ABSOLUTE_Y),
	newOpcode(0xC1, "CMP", 2, 6, INDIRECT_X),
	newOpcode(0xD1, "CMP", 2, 5, INDIRECT_Y),

	newOpcode(0xE0, "CPX", 2, 2, IMMEDIATE),
	newOpcode(0xE4, "CPX", 2, 3, ZEROPAGE),
	newOpcode(0xEC, "CPX", 3, 4, ABSOLUTE),

	newOpcode(0xC0, "CPY", 2, 2, IMMEDIATE),
	newOpcode(0xC4, "CPY", 2, 3, ZEROPAGE),
	newOpcode(0xCC, "CPY", 3, 4, ABSOLUTE),

	newOpcode(0xC6, "DEC", 2, 5, ZEROPAGE),
	newOpcode(0xD6, "DEC", 2, 6, ZEROPAGE_X),
	newOpcode(0xCE, "DEC", 3, 6, ABSOLUTE),
	newOpcode(0xDE, "DEC", 3, 7, ABSOLUTE_X),

	newOpcode(0x49, "EOR", 2, 2, IMMEDIATE),
	newOpcode(0x45, "EOR", 2, 3, ZEROPAGE),
	newOpcode(0x55, "EOR", 2, 4, ZEROPAGE_X),
	newOpcode(0x4D, "EOR", 3, 4, ABSOLUTE),
	newOpcode(0x5D, "EOR", 3, 4, ABSOLUTE_X),
	newOpcode(0x59, "EOR", 3, 4, ABSOLUTE_Y),
	newOpcode(0x41, "EOR", 2, 6, INDIRECT_X),
	newOpcode(0x51, "EOR", 2, 5, INDIRECT_Y),

	newOpcode(0xE6, "INC", 2, 5, ZEROPAGE),
	newOpcode(0xF6, "INC", 2, 6, ZEROPAGE_X),
	newOpcode(0xEE, "INC", 3, 6, ABSOLUTE),
	newOpcode(0xFE, "INC", 3, 7, ABSOLUTE_X),

	newOpcode(0x4C, "JMP", 3, 3, ABSOLUTE),
	newOpcode(0x6C, "JMP", 3, 5, NONE),

	newOpcode(0x20, "JSR", 3, 6, ABSOLUTE),

	newOpcode(0xA9, "LDA", 2, 2, IMMEDIATE),
	newOpcode(0xA5, "LDA", 2, 3, ZEROPAGE),
	newOpcode(0xB5, "LDA", 2, 4, ZEROPAGE_X),
	newOpcode(0xAD, "LDA", 3, 4, ABSOLUTE),
	newOpcode(0xBD, "LDA", 3, 4, ABSOLUTE_X),
	newOpcode(0xB9, "LDA", 3, 4, ABSOLUTE_Y),
	newOpcode(0xA1, "LDA", 2, 6, INDIRECT_X),
	newOpcode(0xB1, "LDA", 2, 5, INDIRECT_Y),

	newOpcode(0xA2, "LDX", 2, 2, IMMEDIATE),
	newOpcode(0xA6, "LDX", 2, 3, ZEROPAGE),
	newOpcode(0xB6, "LDX", 2, 4, ZEROPAGE_Y),
	newOpcode(0xAE, "LDX", 3, 4, ABSOLUTE),
	newOpcode(0xBE, "LDX", 3, 4, ABSOLUTE_Y),

	newOpcode(0xA0, "LDY", 2, 2, IMMEDIATE),
	newOpcode(0xA4, "LDY", 2, 3, ZEROPAGE),
	newOpcode(0xB4, "LDY", 2, 4, ZEROPAGE_X),
	newOpcode(0xAC, "LDY", 3, 4, ABSOLUTE),
	newOpcode(0xBC, "LDY", 3, 4, ABSOLUTE_X),

	newOpcode(0x4A, "LSR", 1, 2, NONE),
	newOpcode(0x46, "LSR", 2, 5, ZEROPAGE),
	newOpcode(0x56, "LSR", 2, 6, ZEROPAGE_X),
	newOpcode(0x4E, "LSR", 3, 6, ABSOLUTE),
	newOpcode(0x5E, "LSR", 3, 7, ABSOLUTE_X),

	newOpcode(0x09, "ORA", 2, 2, IMMEDIATE),
	newOpcode(0x05, "ORA", 2, 3, ZEROPAGE),
	newOpcode(0x15, "ORA", 2, 4, ZEROPAGE_X),
	newOpcode(0x0D, "ORA", 3, 4, ABSOLUTE),
	newOpcode(0x1D, "ORA", 3, 4, ABSOLUTE_X),
	newOpcode(0x19, "ORA", 3, 4, ABSOLUTE_Y),
	newOpcode(0x01, "ORA", 2, 6, INDIRECT_X),
	newOpcode(0x11, "ORA", 2, 5, INDIRECT_Y),

	newOpcode(0x2A, "ROL", 1, 2, NONE),
	newOpcode(0x26, "ROL", 2, 5, ZEROPAGE),
	newOpcode(0x36, "ROL", 2, 6, ZEROPAGE_X),
	newOpcode(0x2E, "ROL", 3, 6, ABSOLUTE),
	newOpcode(0x3E, "ROL", 3, 7, ABSOLUTE_X),

	newOpcode(0x6A, "ROR", 1, 2, NONE),
	newOpcode(0x66, "ROR", 2, 5, ZEROPAGE),
	newOpcode(0x76, "ROR", 2, 6, ZEROPAGE_X),
	newOpcode(0x6E, "ROR", 3, 6, ABSOLUTE),
	newOpcode(0x7E, "ROR", 3, 7, ABSOLUTE_X),

	newOpcode(0x40, "RTI", 1, 6, NONE),

	newOpcode(0x60, "RTS", 1, 6, NONE),

	newOpcode(0xE9, "SBC", 2, 2, IMMEDIATE),
	newOpcode(0xE5, "SBC", 2, 3, ZEROPAGE),
	newOpcode(0xF5, "SBC", 2, 4, ZEROPAGE_X),
	newOpcode(0xED, "SBC", 3, 4, ABSOLUTE),
	newOpcode(0xFD, "SBC", 3, 4, ABSOLUTE_X),
	newOpcode(0xF9, "SBC", 3, 4, ABSOLUTE_Y),
	newOpcode(0xE1, "SBC", 2, 6, INDIRECT_X),
	newOpcode(0xF1, "SBC", 2, 5, INDIRECT_Y),

	newOpcode(0x38, "SEC", 1, 2, NONE),
	newOpcode(0xF8, "SED", 1, 2, NONE),
	newOpcode(0x78, "SEI", 1, 2, NONE),

	newOpcode(0x85, "STA", 2, 3, ZEROPAGE),
	newOpcode(0x95, "STA", 2, 4, ZEROPAGE_X),
	newOpcode(0x8D, "STA", 3, 4, ABSOLUTE),
	newOpcode(0x9D, "STA", 3, 5, ABSOLUTE_X),
	newOpcode(0x99, "STA", 3, 5, ABSOLUTE_Y),
	newOpcode(0x81, "STA", 2, 6, INDIRECT_X),
	newOpcode(0x91, "STA", 2, 6, INDIRECT_Y),

	newOpcode(0x86, "STX", 2, 3, ZEROPAGE),
	newOpcode(0x96, "STX", 2, 4, ZEROPAGE_Y),
	newOpcode(0x8E, "STX", 3, 4, ABSOLUTE),

	newOpcode(0x84, "STY", 2, 3, ZEROPAGE),
	newOpcode(0x94, "STY", 2, 4, ZEROPAGE_X),
	newOpcode(0x8C, "STY", 3, 4, ABSOLUTE),

	newOpcode(0xBA, "TSX", 1, 2, NONE),
	newOpcode(0x8A, "TXA", 1, 2, NONE),
	newOpcode(0x9A, "TXS", 1, 2, NONE),
	newOpcode(0x98, "TYA", 1, 2, NONE),
}

var OpCodeMap = map[byte]*OpCode{}

func init() {
	for _, opcode := range OpCodes {
		OpCodeMap[opcode.Code] = opcode
	}
}
