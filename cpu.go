package main

import "fmt"

type Cpu struct {
	A      byte
	X      byte
	Y      byte
	PC     uint16
	SP     byte
	Status byte
	Bus    *Bus
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

func newCpu() *Cpu {
	return &Cpu{
		SP:     STACK_START,
		Bus:    newBus(),
		Status: 0b100100,
	}
}

func (c *Cpu) snakeLoad(program []byte) {
	for i, b := range program {
		c.Write(0x0600+uint16(i), b)
	}
	c.Write16(0xFFFC, 0x0600)
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
		return uint16(c.Read(c.PC))
	case ABSOLUTE:
		return c.Read16(c.PC)
	case ZEROPAGE_X:
		pos := c.Read(c.PC)
		addr := c.X + pos
		return uint16(addr)
	case ZEROPAGE_Y:
		pos := c.Read(c.PC)
		addr := c.Y + pos
		return uint16(addr)
	case ABSOLUTE_X:
		addr := c.Read16(c.PC)
		return addr + uint16(c.X)
	case ABSOLUTE_Y:
		addr := c.Read16(c.PC)
		return addr + uint16(c.Y)
	case INDIRECT_X:
		pos := c.Read(c.PC)
		ptr := c.X + pos
		lo := uint16(c.Read(uint16(ptr)))
		hi := uint16(c.Read(uint16(ptr + 1)))
		return (hi << 8) | lo
	case INDIRECT_Y:
		pos := c.Read(c.PC)
		lo := c.Read(uint16(pos))
		hi := c.Read(uint16(pos + 1))
		return ((uint16(hi) << 8) | uint16(lo)) + uint16(c.Y)
	default:
		panic(fmt.Sprintf("Unknown addressing mode: %d", mode))
	}
}

func (c *Cpu) Run() {
	c.RunWithCallback(func(c *Cpu) {})
}

func (c *Cpu) RunWithCallback(f func(*Cpu)) {
	for {
		f(c)

		code := c.Read(c.PC)
		c.PC += 1
		opcode := OpCodeMap[code]
		programCounterState := c.PC

		switch code {
		case 0x69, 0x65, 0x75, 0x6D, 0x7D, 0x79, 0x61, 0x71:
			c.ADC(opcode.Mode)
		case 0x29, 0x25, 0x35, 0x2D, 0x3D, 0x39, 0x21, 0x31:
			c.AND(opcode.Mode)
		case 0x24, 0x2C:
			c.BIT(opcode.Mode)
		case 0xC9, 0xC5, 0xD5, 0xCD, 0xDD, 0xD9, 0xC1, 0xD1:
			c.CMP(opcode.Mode)
		case 0xE0, 0xE4, 0xEC:
			c.CPX(opcode.Mode)
		case 0xC0, 0xC4, 0xCC:
			c.CPY(opcode.Mode)
		case 0x49, 0x45, 0x55, 0x4D, 0x5D, 0x59, 0x41, 0x51:
			c.EOR(opcode.Mode)
		case 0xA9, 0xA5, 0xB5, 0xAD, 0xBD, 0xB9, 0xA1, 0xB1:
			c.LDA(opcode.Mode)
		case 0xA2, 0xA6, 0xB6, 0xAE, 0xBE:
			c.LDX(opcode.Mode)
		case 0xA0, 0xA4, 0xB4, 0xAC, 0xBC:
			c.LDY(opcode.Mode)
		case 0xE9, 0xE5, 0xF5, 0xED, 0xFD, 0xF9, 0xE1, 0xF1:
			c.SBC(opcode.Mode)
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
		case 0x4A:
			c.LSRAccumulator()
		case 0x46, 0x56, 0x4E, 0x5E:
			c.LSR(opcode.Mode)
		case 0x09, 0x05, 0x15, 0x0D, 0x1D, 0x19, 0x01, 0x11:
			c.ORA(opcode.Mode)
		case 0x2A:
			c.ROLAccumulator()
		case 0x26, 0x36, 0x2E, 0x3E:
			c.ROL(opcode.Mode)
		case 0x6A:
			c.RORAccumulator()
		case 0x66, 0x76, 0x6E, 0x7E:
			c.ROR(opcode.Mode)
		// JMP ABSOLUTe
		case 0x4C:
			addr := c.Read16(c.PC)
			c.PC = addr
		// JMP INDIRECT
		case 0x6C:
			addr := c.Read16(c.PC)
			var indirect_ref uint16
			if addr&0x00FF == 0x00FF {
				lo := c.Read(addr)
				hi := c.Read(addr & 0xFF00)
				indirect_ref = (uint16(hi) << 8) | uint16(lo)
			} else {
				indirect_ref = c.Read16(addr)
			}
			c.PC = indirect_ref
		case 0x20:
			c.pushStack16(c.PC + 2 - 1)
			c.PC = c.Read16(c.PC)
		case 0x60:
			c.PC = c.popStack16() + 1
		case 0x40:
			c.RTI()
		case 0x90:
			c.branch(!c.getFlag(CARRY))
		case 0xB0:
			c.branch(c.getFlag(CARRY))
		case 0xF0:
			c.branch(c.getFlag(ZERO))
		case 0x30:
			c.branch(c.getFlag(NEGATIVE))
		case 0xD0:
			c.branch(!c.getFlag(ZERO))
		case 0x10:
			c.branch(!c.getFlag(NEGATIVE))
		case 0x50:
			c.branch(!c.getFlag(OVERFLOW))
		case 0x70:
			c.branch(c.getFlag(OVERFLOW))
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
		case 0xBA:
			c.TSX()
		case 0x8A:
			c.TXA()
		case 0x9A:
			c.TXS()
		case 0x98:
			c.TYA()
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
		case 0x18:
			c.setFlag(CARRY, false)
		case 0xD8:
			c.setFlag(DECIMAL, false)
		case 0x58:
			c.setFlag(INTERRUPT, false)
		case 0xB8:
			c.setFlag(OVERFLOW, false)
		case 0x38:
			c.setFlag(CARRY, true)
		case 0xF8:
			c.setFlag(DECIMAL, true)
		case 0x78:
			c.setFlag(INTERRUPT, true)
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
	for i, b := range program {
		c.Write(0x0000+uint16(i), b)
	}
	c.Write16(0xFFFC, 0x0000)
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
	c.PC = c.Read16(0xFFFC)
}

func (c *Cpu) ASLAccumulator() {
	if c.A&0x80 > 0 {
		c.setFlag(CARRY, true)
	} else {
		c.setFlag(CARRY, false)
	}

	value := c.A << 1
	c.A = value
	c.setZeroAndNegativeFlags(value)
}

func (c *Cpu) LSRAccumulator() {
	if c.A&0x01 > 0 {
		c.setFlag(CARRY, true)
	} else {
		c.setFlag(CARRY, false)
	}

	c.A = c.A >> 1
	c.setZeroAndNegativeFlags(c.A)
}

func (c *Cpu) ROLAccumulator() {
	carry := c.getFlag(CARRY)
	if c.A&0x80 > 0 {
		c.setFlag(CARRY, true)
	} else {
		c.setFlag(CARRY, false)
	}

	c.A = c.A << 1
	if carry {
		c.A |= 0x01
	}

	c.setZeroAndNegativeFlags(c.A)
}

func (c *Cpu) RORAccumulator() {
	carry := c.getFlag(CARRY)
	if c.A&0x01 > 0 {
		c.setFlag(CARRY, true)
	} else {
		c.setFlag(CARRY, false)
	}

	c.A = c.A >> 1
	if carry {
		c.A |= 0x80
	}

	c.setZeroAndNegativeFlags(c.A)
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

func (c *Cpu) pushStack(value byte) {
	c.Write(STACK+uint16(c.SP), value)
	c.SP -= 1
}

func (c *Cpu) popStack() byte {
	c.SP += 1
	return c.Read(STACK + uint16(c.SP))
}

func (c *Cpu) branch(condition bool) {
	if condition {
		jump := int8(c.Read(c.PC))
		addr := c.PC + 1 + uint16(jump)
		c.PC = addr
	}
}

func (c *Cpu) pushStack16(value uint16) {
	hi := byte(value >> 8)
	lo := byte(value & 0xFF)
	c.pushStack(hi)
	c.pushStack(lo)
}

func (c *Cpu) popStack16() uint16 {
	lo := uint16(c.popStack())
	hi := uint16(c.popStack())
	return (hi << 8) | lo
}

// Read & Write

func (c *Cpu) Read(addr uint16) byte {
	return c.Bus.Read(addr)
}

func (c *Cpu) Read16(addr uint16) uint16 {
	lo := uint16(c.Bus.Read(addr))
	hi := uint16(c.Bus.Read(addr + 1))
	return (hi << 8) | lo
}

func (c *Cpu) Write16(addr, data uint16) {
	hi := uint8(data >> 8)
	lo := uint8(data & 0xFF)
	c.Write(addr, lo)
	c.Write(addr+1, hi)
}

func (c *Cpu) Write(addr uint16, data byte) {
	c.Bus.Write(addr, data)
}

// CPU Instructions
func (c *Cpu) ADC(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.Read(addr)
	c.addToA(value)

	c.setZeroAndNegativeFlags(c.A)
}

func (c *Cpu) AND(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.Read(addr)
	result := c.A & value
	c.setZeroAndNegativeFlags(result)
	c.A = result
}

func (c *Cpu) ASL(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.Read(addr)
	result := value << 1

	if value&0x80 > 0 {
		c.setFlag(CARRY, true)
	} else {
		c.setFlag(CARRY, false)
	}

	c.setZeroAndNegativeFlags(result)
	c.Write(addr, result)
}

func (c *Cpu) BIT(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.Read(addr)
	result := c.A & value
	c.setFlag(ZERO, result == 0)
	c.setFlag(NEGATIVE, value&0x80 != 0)
	c.setFlag(OVERFLOW, value&0x40 != 0)
}

func (c *Cpu) CMP(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.Read(addr)
	result := c.A - value
	c.setZeroAndNegativeFlags(result)
	c.setFlag(CARRY, c.A >= value)
}

func (c *Cpu) CPX(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.Read(addr)
	result := c.X - value
	c.setZeroAndNegativeFlags(result)
	c.setFlag(CARRY, c.X >= value)
}

func (c *Cpu) CPY(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.Read(addr)
	result := c.Y - value
	c.setZeroAndNegativeFlags(result)
	c.setFlag(CARRY, c.Y >= value)
}

func (c *Cpu) EOR(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.Read(addr)
	result := c.A ^ value
	c.setZeroAndNegativeFlags(result)
	c.A = result
}

func (c *Cpu) LDA(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.Read(addr)
	c.A = value
	c.setZeroAndNegativeFlags(c.A)
}

func (c *Cpu) LDX(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.Read(addr)
	c.X = value
	c.setZeroAndNegativeFlags(c.X)
}

func (c *Cpu) LDY(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.Read(addr)
	c.Y = value
	c.setZeroAndNegativeFlags(c.Y)
}

func (c *Cpu) LSR(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.Read(addr)
	result := value >> 1

	if value&0x01 > 0 {
		c.setFlag(CARRY, true)
	} else {
		c.setFlag(CARRY, false)
	}

	c.setZeroAndNegativeFlags(result)
	c.Write(addr, result)
}

func (c *Cpu) ORA(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.Read(addr)
	result := c.A | value
	c.setZeroAndNegativeFlags(result)
	c.A = result
}

func (c *Cpu) ROL(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.Read(addr)
	carry := c.getFlag(CARRY)
	result := value << 1
	if carry {
		result |= 0x01
	}

	if value&0x80 > 0 {
		c.setFlag(CARRY, true)
	} else {
		c.setFlag(CARRY, false)
	}

	c.setZeroAndNegativeFlags(result)
	c.Write(addr, result)
}

func (c *Cpu) ROR(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.Read(addr)
	carry := c.getFlag(CARRY)
	result := value >> 1
	if carry {
		result |= 0x80
	}

	if value&0x01 > 0 {
		c.setFlag(CARRY, true)
	} else {
		c.setFlag(CARRY, false)
	}

	c.setZeroAndNegativeFlags(result)
	c.Write(addr, result)
}

func (c *Cpu) SBC(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	value := c.Read(addr)
	c.addToA(^value + 1)
	c.setZeroAndNegativeFlags(c.A)
}

func (c *Cpu) STA(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	c.Write(addr, c.A)
}

func (c *Cpu) STX(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	c.Write(addr, c.X)
}

func (c *Cpu) STY(mode AddressingMode) {
	addr := c.getOperandAddress(mode)
	c.Write(addr, c.Y)
}

func (c *Cpu) PHA() {
	c.Write(STACK+uint16(c.SP), c.A)
	c.SP -= 1
}

func (c *Cpu) PLA() {
	c.SP += 1
	c.A = c.Read(STACK + uint16(c.SP))
	c.setZeroAndNegativeFlags(c.A)
}

func (c *Cpu) PHP() {
	c.Write(STACK+uint16(c.SP), c.Status)
	c.SP -= 1
}

func (c *Cpu) PLP() {
	c.SP += 1
	c.Status = c.Read(STACK + uint16(c.SP))
}

func (c *Cpu) RTI() {
	c.Status = c.popStack()
	c.setFlag(UNUSED, true)
	c.setFlag(BREAK, false)
	c.PC = c.popStack16()
}

func (c *Cpu) TAX() {
	c.X = c.A
	c.setZeroAndNegativeFlags(c.X)
}

func (c *Cpu) TAY() {
	c.Y = c.A
	c.setZeroAndNegativeFlags(c.Y)
}

func (c *Cpu) TSX() {
	c.X = c.SP
	c.setZeroAndNegativeFlags(c.X)
}

func (c *Cpu) TXA() {
	c.A = c.X
	c.setZeroAndNegativeFlags(c.A)
}

func (c *Cpu) TXS() {
	c.SP = c.X
}

func (c *Cpu) TYA() {
	c.A = c.Y
	c.setZeroAndNegativeFlags(c.A)
}

func (c *Cpu) DEC() {
	addr := c.getOperandAddress(ZEROPAGE)
	value := c.Read(addr) - 1
	c.Write(addr, value)
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
	value := c.Read(addr) + 1
	c.Write(addr, value)
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
