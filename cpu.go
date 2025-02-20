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
		case 0xAA:
			c.TAX()
		case 0xA8:
			c.TAY()
		case 0xE8:
			c.INX()
		case 0xC8:
			c.INY()
		case 0xEA:
			// NOP
			continue
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

	c.PC = c.read16(0xFFFC)
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

func (c *Cpu) TAX() {
	c.X = c.A
	c.setZeroAndNegativeFlags(c.X)
}

func (c *Cpu) TAY() {
	c.Y = c.A
	c.setZeroAndNegativeFlags(c.Y)
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
