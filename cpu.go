package main

import "fmt"

type Cpu struct {
	A      byte
	X      byte
	Y      byte
	PC     uint16
	SP     byte
	Status byte
	memory [0xFFFF]byte
}

type StatusFlag uint8

const STACK_START = 0xFD
const STACK = 0x100

const (
	CARRY StatusFlag = 1 << iota
	ZERO
	INTERRUPT
	DECIMAL
	BREAK
	UNUSED
	OVERFLOW
	NEGATIVE
)

type AddressingMode uint8

const (
	IMMEDIATE AddressingMode = iota
	ZEROPAGE
	ZEROPAGE_X
	ZEROPAGE_Y
	ABSOLUTE
	ABSOLUTE_X
	ABSOLUTE_Y
	INDIRECT_X
	INDIRECT_Y
	NONE
)

func new_cpu() *Cpu {
	return &Cpu{}
}

func (c *Cpu) getFlag(flag StatusFlag) bool {
	return c.Status&uint8(flag) > 0
}

func (c *Cpu) setFlag(flag StatusFlag, value bool) {
	if value {
		c.Status |= uint8(flag)
	} else {
		c.Status &= ^uint8(flag)
	}
}

func (c *Cpu) getOperandAddress(mode AddressingMode) uint16 {
	switch mode {
	case IMMEDIATE:
		return c.PC
	case ZEROPAGE:
		return uint16(c.read(c.PC))
	case ABSOLUTE:
		return c.read16(c.PC)
	case ZEROPAGE_X:
		pos := c.read(c.PC)
		addr := c.X + pos
		return uint16(addr)
	case ZEROPAGE_Y:
		pos := c.read(c.PC)
		addr := c.Y + pos
		return uint16(addr)
	case ABSOLUTE_X:
		addr := c.read16(c.PC)
		return addr + uint16(c.X)
	case ABSOLUTE_Y:
		addr := c.read16(c.PC)
		return addr + uint16(c.Y)
	case INDIRECT_X:
		pos := c.read(c.PC)
		ptr := c.X + pos
		lo := uint16(c.read(uint16(ptr)))
		hi := uint16(c.read(uint16(ptr + 1)))
		return (hi << 8) | lo
	case INDIRECT_Y:
		pos := c.read(c.PC)
		lo := c.read(uint16(pos))
		hi := c.read(uint16(pos + 1))
		return (uint16(hi) << 8) | uint16(lo) + uint16(c.Y)
	default:
		panic(fmt.Sprintf("Unknown addressing mode: %d", mode))
	}
}

func (c *Cpu) Run() {
	for {
		code := c.memory[c.PC]
		c.PC += 1
		opcode := OpCodeMap[code]
		programCounterState := c.PC

		switch code {
		case 0x69, 0x65, 0x75, 0x6D, 0x7D, 0x79, 0x61, 0x71:
			c.ADC(opcode.Mode)
		case 0x29, 0x25, 0x35, 0x2D, 0x3D, 0x39, 0x21, 0x31:
			c.AND(opcode.Mode)
		case 0xA9, 0xA5, 0xB5, 0xAD, 0xBD, 0xB9, 0xA1, 0xB1:
			c.LDA(opcode.Mode)
		case 0xA2, 0xA6, 0xB6, 0xAE, 0xBE:
			c.LDX(opcode.Mode)
		case 0xA0, 0xA4, 0xB4, 0xAC, 0xBC:
			c.LDY(opcode.Mode)
		case 0x85, 0x95, 0x8D, 0x9D, 0x99, 0x81, 0x91:
			c.STA(opcode.Mode)
		case 0x86, 0x96, 0x8E:
			c.STX(opcode.Mode)
		case 0x84, 0x94, 0x8C:
			c.STY(opcode.Mode)
		case 0xE6, 0xF6, 0xEE, 0xFE:
			c.INC(opcode.Mode)
		case 0xC6, 0xD6, 0xCE, 0xDE:
			c.DEC()
		case 0x0A:
			c.ASLAccumulator()
		case 0x06, 0x16, 0x0E, 0x1E:
			c.ASL(opcode.Mode)
		// JMP ABSOLUTe
		case 0x4C:
			addr := c.read16(c.PC)
			c.PC = addr
		// JMP INDIRECT
		case 0x6C:
			addr := c.read16(c.PC)
			var indirect_ref uint16
			if addr&0x00FF == 0x00FF {
				lo := c.read(addr)
				hi := c.read(addr & 0xFF00)
				indirect_ref = (uint16(hi) << 8) | uint16(lo)
			} else {
				indirect_ref = c.read16(addr)
			}
			c.PC = indirect_ref
		case 0x48:
			c.PHA()
		case 0x68:
			c.PLA()
		case 0x08:
			c.PHP()
		case 0x28:
			c.PLP()
		case 0xAA:
			c.TAX()
		case 0xA8:
			c.TAY()
		case 0xE8:
			c.INX()
		case 0xC8:
			c.INY()
		case 0xCA:
			c.DEX()
		case 0x88:
			c.DEY()
		case 0xEA:
			c.NOP()
		case 0x00:
			return
		default:
			panic(fmt.Sprintf("Unknown opcode: %02x", code))
		}

		if c.PC == programCounterState {
			c.PC += uint16(opcode.Length - 1)
		}
	}
}

func (c *Cpu) setZeroAndNegativeFlags(value byte) {
	if value == 0 {
		c.setFlag(ZERO, true)
	} else {
		c.setFlag(ZERO, false)
	}

	if value&0x80 > 0 {
		c.setFlag(NEGATIVE, true)
	} else {
		c.setFlag(NEGATIVE, false)
	}
}

func (c *Cpu) Load(program []byte) {
	copy(c.memory[0x8000:], program)
	c.PC = 0x8000
	c.write16(0xFFFC, 0x8000)
}

func (c *Cpu) LoadAndRun(program []byte) {
	c.Load(program)
	c.Reset()
	c.Run()
}

func (c *Cpu) Reset() {
	c.A = 0
	c.X = 0
	c.Y = 0
	c.Status = 0

	c.SP = STACK_START
	c.PC = c.read16(0xFFFC)
}

func (c *Cpu) ASLAccumulator() {
	value := c.A << 1
	c.setZeroAndNegativeFlags(value)
	if c.A&0x80 > 0 {
		c.setFlag(CARRY, true)
	} else {
		c.setFlag(CARRY, false)
	}
	c.A = value
}

func (c *Cpu) addToA(value byte) {
	result := c.A + value

	if (c.A^result)&(value^result)&0x80 > 0 {
		c.setFlag(OVERFLOW, true)
	} else {
		c.setFlag(OVERFLOW, false)
	}

	if result < c.A {
		c.setFlag(CARRY, true)
	} else {
		c.setFlag(CARRY, false)
	}
	c.A = result
}

// Read & Write

func (c *Cpu) read(addr uint16) byte {
	return c.memory[addr]
}

func (c *Cpu) read16(addr uint16) uint16 {
	lo := uint16(c.read(addr))
	hi := uint16(c.read(addr + 1))
	return (hi << 8) | lo
}

func (c *Cpu) write16(addr, data uint16) {
	hi := uint8(data >> 8)
	lo := uint8(data & 0xFF)
	c.write(addr, lo)
	c.write(addr+1, hi)
}

func (c *Cpu) write(addr uint16, data byte) {
	c.memory[addr] = data
}

// CPU Instructions
func (c *Cpu) ADC(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.read(addr)
	c.addToA(value)

	c.setZeroAndNegativeFlags(c.A)
}

func (c *Cpu) AND(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.read(addr)
	result := c.A & value
	c.setZeroAndNegativeFlags(result)
	c.A = result
}

func (c *Cpu) ASL(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.read(addr)
	result := value << 1

	if value&0x80 > 0 {
		c.setFlag(CARRY, true)
	} else {
		c.setFlag(CARRY, false)
	}

	c.setZeroAndNegativeFlags(result)
	c.write(addr, result)
}

func (c *Cpu) LDA(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.read(addr)
	c.A = value
	c.setZeroAndNegativeFlags(c.A)
}

func (c *Cpu) LDX(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.read(addr)
	c.X = value
	c.setZeroAndNegativeFlags(c.X)
}

func (c *Cpu) LDY(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.read(addr)
	c.Y = value
	c.setZeroAndNegativeFlags(c.Y)
}

func (c *Cpu) STA(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	c.write(addr, c.A)
}

func (c *Cpu) STX(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	c.write(addr, c.X)
}

func (c *Cpu) STY(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	c.write(addr, c.Y)
}

func (c *Cpu) PHA() {
	c.write(STACK+uint16(c.SP), c.A)
	c.SP -= 1
}

func (c *Cpu) PLA() {
	c.SP += 1
	c.A = c.read(STACK + uint16(c.SP))
	c.setZeroAndNegativeFlags(c.A)
}

func (c *Cpu) PHP() {
	c.write(STACK+uint16(c.SP), c.Status)
	c.SP -= 1
}

func (c *Cpu) PLP() {
	c.SP += 1
	c.Status = c.read(STACK + uint16(c.SP))
}

func (c *Cpu) TAX() {
	c.X = c.A
	c.setZeroAndNegativeFlags(c.X)
}

func (c *Cpu) TAY() {
	c.Y = c.A
	c.setZeroAndNegativeFlags(c.Y)
}

func (c *Cpu) DEC() {
	addr := c.getOperandAddress(ZEROPAGE)
	value := c.read(addr) - 1
	c.write(addr, value)
	c.setZeroAndNegativeFlags(value)
}

func (c *Cpu) DEX() {
	c.X -= 1
	c.setZeroAndNegativeFlags(c.X)
}

func (c *Cpu) DEY() {
	c.Y -= 1
	c.setZeroAndNegativeFlags(c.Y)
}

func (c *Cpu) NOP() {
	// No operation
}

func (c *Cpu) INC(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.read(addr) + 1
	c.write(addr, value)
	c.setZeroAndNegativeFlags(value)

}

func (c *Cpu) INX() {
	c.X += 1
	c.setZeroAndNegativeFlags(c.X)
}

func (c *Cpu) INY() {
	c.Y += 1
	c.setZeroAndNegativeFlags(c.Y)
}
