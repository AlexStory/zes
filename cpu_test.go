package main

import (
	"testing"
)

func Test0xA9LDAImmediate(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0x00})

		testA(t, cpu, 0x05)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("zero is set properly", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x00, 0x00})

		testA(t, cpu, 0x00)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xA5LDAZeroPage(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0x55)
		cpu.LoadAndRun([]byte{0xA5, 0x05, 0x00})

		testA(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xB5LDAZeroPageX(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0x55)
		cpu.LoadAndRun([]byte{0xA2, 0x03, 0xB5, 0x02, 0x00})

		testA(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xADLDAAbsolute(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0200, 0x55)
		cpu.LoadAndRun([]byte{0xAD, 0x00, 0x02, 0x00})

		testA(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xBDLDAAbsoluteX(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0x55)
		cpu.LoadAndRun([]byte{0xA2, 0x02, 0xBD, 0x00, 0x02, 0x00})

		testA(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xB9LDAAbsoluteY(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0x55)
		cpu.LoadAndRun([]byte{0xA0, 0x02, 0xB9, 0x00, 0x02, 0x00})

		testA(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xA1LDAIndirectX(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0x34)
		cpu.Write(0x06, 0x12)
		cpu.Write(0x1234, 0x55)

		cpu.LoadAndRun([]byte{0xA2, 0x03, 0xA1, 0x02, 0x00})

		testA(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xB1LDAIndirectY(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0x32)
		cpu.Write(0x06, 0x12)
		cpu.Write(0x1234, 0x55)
		cpu.LoadAndRun([]byte{0xA0, 0x02, 0xB1, 0x05, 0x00})

		testA(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x85STAZeroPage(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x55, 0x85, 0x05, 0x00})

		testMem(t, cpu, 0x05, 0x55)
	})
}

func Test0x95STAZeroPageX(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA2, 0x03, 0xA9, 0x55, 0x95, 0x02, 0x00})

		testMem(t, cpu, 0x05, 0x55)
	})
}

func Test0x8DSTAAbsolute(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x55, 0x8D, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0200, 0x55)
	})
}

func Test0x9DSTAAbsoluteX(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA2, 0x02, 0xA9, 0x55, 0x9D, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0202, 0x55)
	})
}

func Test0x81STAIndirectX(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{
			0xA2, 0x03,
			0xA9, 0x34,
			0x85, 0x05,
			0xA9, 0x12,
			0x85, 0x06,
			0xA9, 0x55,
			0x81, 0x02,
			0x00,
		})

		testMem(t, cpu, 0x1234, 0x55)
	})
}

func Test0x91STAIndirectY(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{
			0xA9, 0x32,
			0x85, 0x80,
			0xA9, 0x12,
			0x85, 0x81,
			0xA0, 0x02,
			0xA9, 0x55,
			0x91, 0x80,
			0x00,
		})

		testMem(t, cpu, 0x1234, 0x55)
	})
}

func Test0x99STAAbsoluteY(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA0, 0x02, 0xA9, 0x55, 0x99, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0202, 0x55)
	})
}

func Test0xAATAX(t *testing.T) {
	t.Run("transefer a to x", func(t *testing.T) {
		cpu := newCpu()
		program := []byte{0xA9, 0x05, 0xAA, 0x00}

		cpu.LoadAndRun(program)

		testA(t, cpu, 0x05)
		testX(t, cpu, 0x05)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("zero is set properly", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x00, 0xAA, 0x00})

		testX(t, cpu, 0x00)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xE8INX(t *testing.T) {
	t.Run("increment x", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0xAA, 0xE8, 0x00})

		testA(t, cpu, 0x05)
		testX(t, cpu, 0x06)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("Test x overflow", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA2, 0xFF, 0xE8, 0xE8, 0x00})

		testX(t, cpu, 0x01)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xA2LDXImmediate(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA2, 0x05, 0x00})

		testX(t, cpu, 0x05)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("zero is set properly", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA2, 0x00, 0x00})

		testX(t, cpu, 0x00)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xA6LDXZeroPage(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0x55)
		cpu.LoadAndRun([]byte{0xA6, 0x05, 0x00})

		testX(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xB6LDXZeroPageY(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0x55)
		cpu.LoadAndRun([]byte{0xA0, 0x01, 0xB6, 0x04, 0x00})

		testX(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xAELDXAbsolute(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0200, 0x55)
		cpu.LoadAndRun([]byte{0xAE, 0x00, 0x02, 0x00})

		testX(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xBELDXAbsoluteY(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0x55)
		cpu.LoadAndRun([]byte{0xA0, 0x02, 0xBE, 0x00, 0x02, 0x00})

		testX(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xA0LDYImmediate(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA0, 0x05, 0x00})

		testY(t, cpu, 0x05)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xA4LDYZeroPage(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0x55)
		cpu.LoadAndRun([]byte{0xA4, 0x05, 0x00})

		testY(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xB4LDYZeroPageX(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0x55)
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0xB4, 0x04, 0x00})

		testY(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xACLDYAbsolute(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0200, 0x55)
		cpu.LoadAndRun([]byte{0xAC, 0x00, 0x02, 0x00})

		testY(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xBCLDYAbsoluteX(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0x55)
		cpu.LoadAndRun([]byte{0xA2, 0x02, 0xBC, 0x00, 0x02, 0x00})

		testY(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func TestSimpleProgram(t *testing.T) {
	cpu := newCpu()
	program := []byte{0xA9, 0xc0, 0xAA, 0xE8, 0x00}
	cpu.LoadAndRun(program)

	testA(t, cpu, 0xC0)
	testX(t, cpu, 0xC1)
}

func Test0x86STXZeroPage(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA2, 0x05, 0x86, 0x05, 0x00})

		testMem(t, cpu, 0x05, 0x05)
	})
}

func Test0x96STXZeroPageY(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA0, 0x01, 0xA2, 0x05, 0x96, 0x04, 0x00})

		testMem(t, cpu, 0x05, 0x05)
	})
}

func Test0x8ESTXAbsolute(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA2, 0x05, 0x8E, 0x00, 0x02, 0x00})
		testMem(t, cpu, 0x0200, 0x05)
	})
}

func Test0x84STYZeroPage(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA0, 0x05, 0x84, 0x05, 0x00})

		testMem(t, cpu, 0x05, 0x05)
	})
}

func Test0x94STYZeroPageX(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0xA0, 0x05, 0x94, 0x04, 0x00})

		testMem(t, cpu, 0x05, 0x05)
	})
}

func Test0x8CTestSTYAbsolute(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA0, 0x05, 0x8C, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0200, 0x05)
	})
}

func Test0xA8TAY(t *testing.T) {
	t.Run("transfer a to y", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0xA8, 0x00})

		testY(t, cpu, 0x05)
	})
}

func Test0xC8INY(t *testing.T) {
	t.Run("increment y", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA0, 0x05, 0xC8, 0x00})

		testY(t, cpu, 0x06)
	})
}

func Test0xE6INCZeroPage(t *testing.T) {
	t.Run("increment memory", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0x55)
		cpu.LoadAndRun([]byte{0xE6, 0x05, 0x00})

		testMem(t, cpu, 0x05, 0x56)
	})
}

func Test0xF6INCZeroPageX(t *testing.T) {
	t.Run("increment memory", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0x55)
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0xF6, 0x04, 0x00})

		testMem(t, cpu, 0x05, 0x56)
	})
}

func Test0xEEINCAbsolute(t *testing.T) {
	t.Run("increment memory", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0200, 0x55)
		cpu.LoadAndRun([]byte{0xEE, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0200, 0x56)
	})
}

func Test0xFEINCAbsoluteX(t *testing.T) {
	t.Run("increment memory", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x01234, 0x55)
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0xFE, 0x33, 0x12, 0x00})

		testMem(t, cpu, 0x1234, 0x56)
	})
}

func Test0xCADEX(t *testing.T) {
	t.Run("decrement x", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA2, 0x05, 0xCA, 0x00})

		testX(t, cpu, 0x04)
	})
}

func Test0x88DEY(t *testing.T) {
	t.Run("decrement y", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA0, 0x05, 0x88, 0x00})

		testY(t, cpu, 0x04)
	})
}

func Test0xC6DECZeroPage(t *testing.T) {
	t.Run("decrement memory", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0x55)
		cpu.LoadAndRun([]byte{0xC6, 0x05, 0x00})

		testMem(t, cpu, 0x05, 0x54)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xEANOP(t *testing.T) {
	t.Run("do nothing", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xEA, 0x00})

		testA(t, cpu, 0x00)
		testX(t, cpu, 0x00)
		testY(t, cpu, 0x00)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, OVERFLOW, false)
	})

	t.Run("NOP doesn't break program", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xEA, 0xA9, 0x05, 0x00})

		testA(t, cpu, 0x05)
	})
}

func Test0x48PHA(t *testing.T) {
	t.Run("push a to stack", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0x48, 0x00})

		testStack(t, cpu, 0x05)
		testStackAddress(t, cpu, 0xFC)
	})
}

func Test0x68PLA(t *testing.T) {
	t.Run("pop a from stack", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0x48, 0xA9, 0x00, 0x68, 0x00})

		testA(t, cpu, 0x05)
	})
}

func Test0x08PHP(t *testing.T) {
	t.Run("push status to stack", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x00, 0x08, 0x00})

		testStack(t, cpu, 0x02)
		testStackAddress(t, cpu, 0xFC)
	})
}

func Test0x28PLP(t *testing.T) {
	t.Run("pop status from stack", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0b10101010, 0x48, 0x28})

		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, INTERRUPT, false)
		testFlag(t, cpu, DECIMAL, true)
		testFlag(t, cpu, BREAK, false)
		testFlag(t, cpu, UNUSED, true)
		testFlag(t, cpu, OVERFLOW, false)
		testFlag(t, cpu, NEGATIVE, true)
	})
}

func Test0x69ADCImmediate(t *testing.T) {
	t.Run("add with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0x69, 0x02, 0x00})

		testA(t, cpu, 0x07)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("add with carry carry bit", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0xFF, 0x69, 0x01, 0x00})

		testA(t, cpu, 0x00)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("add with signed overflow", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x50, 0x69, 0x50, 0x00})

		testA(t, cpu, 0xa0)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, true)
		testFlag(t, cpu, OVERFLOW, true)
	})
}

func Test065ADCZeroPage(t *testing.T) {
	t.Run("add with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0x02)
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0x65, 0x05, 0x00})

		testA(t, cpu, 0x07)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x75ADCZeroPageX(t *testing.T) {
	t.Run("add with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{
			0xA9, 0x02,
			0x85, 0x90,
			0xA2, 0x01,
			0xA9, 0x05,
			0x75, 0x8F,
			0x00,
		})

		testA(t, cpu, 0x07)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x6DADCAbsolute(t *testing.T) {
	t.Run("add with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0200, 0x02)
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0x6D, 0x00, 0x02, 0x00})

		testA(t, cpu, 0x07)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x7DADCAbsoluteX(t *testing.T) {
	t.Run("add with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0x02)
		cpu.LoadAndRun([]byte{0xA2, 0x02, 0xA9, 0x05, 0x7D, 0x00, 0x02, 0x00})

		testA(t, cpu, 0x07)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x79ADCAbsoluteY(t *testing.T) {
	t.Run("add with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0x02)
		cpu.LoadAndRun([]byte{0xA0, 0x02, 0xA9, 0x05, 0x79, 0x00, 0x02, 0x00})

		testA(t, cpu, 0x07)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x61ADCIndirectX(t *testing.T) {
	t.Run("add with carry", func(t *testing.T) {
		cpu := newCpu()

		cpu.LoadAndRun([]byte{
			0xA9, 0x34,
			0x85, 0x90,
			0xA9, 0x12,
			0x85, 0x91,
			0xA9, 0x02,
			0x8D, 0x34, 0x12,
			0xA2, 0x01,
			0xA9, 0x05,
			0x61, 0x8F,
			0x00,
		})

		testA(t, cpu, 0x07)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x71ADCIndirectY(t *testing.T) {
	t.Run("add with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{
			0xA9, 0x32,
			0x85, 0x80,
			0xA9, 0x12,
			0x85, 0x81,
			0xA9, 0x02,
			0x8D, 0x34, 0x12,
			0xA0, 0x02,
			0xA9, 0x05,
			0x71, 0x80,
			0x00,
		})

		testA(t, cpu, 0x07)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x29ANDImmediate(t *testing.T) {
	t.Run("and", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0b1111_0000, 0x29, 0b0011_1111, 0x00})

		testA(t, cpu, 0b0011_0000)
	})
}

func Test0x25ANDZeroPage(t *testing.T) {
	t.Run("and", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0b0011_1111)
		cpu.LoadAndRun([]byte{0xA9, 0b1111_0000, 0x25, 0x05, 0x00})

		testA(t, cpu, 0b0011_0000)
	})
}

func Test0x35ANDZeroPageX(t *testing.T) {
	t.Run("and", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{
			0xA9, 0b0011_1111,
			0x85, 0x90,
			0xA2, 0x01,
			0xA9, 0b1111_0000,
			0x35, 0x8F,
			0x00,
		})

		testA(t, cpu, 0b0011_0000)
	})
}

func Test0x2DANDAbsolute(t *testing.T) {
	t.Run("and", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0200, 0b0011_1111)
		cpu.LoadAndRun([]byte{0xA9, 0b1111_0000, 0x2D, 0x00, 0x02, 0x00})

		testA(t, cpu, 0b0011_0000)
	})
}

func Test0x3DANDAbsoluteX(t *testing.T) {
	t.Run("and", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0b0011_1111)
		cpu.LoadAndRun([]byte{0xA2, 0x02, 0xA9, 0b1111_0000, 0x3D, 0x00, 0x02, 0x00})

		testA(t, cpu, 0b0011_0000)
	})
}

func Test0x39ANDAbsoluteY(t *testing.T) {
	t.Run("and", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0b0011_1111)
		cpu.LoadAndRun([]byte{0xA0, 0x02, 0xA9, 0b1111_0000, 0x39, 0x00, 0x02, 0x00})

		testA(t, cpu, 0b0011_0000)
	})
}

func Test0x21ANDIndirectX(t *testing.T) {
	t.Run("and", func(t *testing.T) {
		cpu := newCpu()

		cpu.LoadAndRun([]byte{
			0xA9, 0x34,
			0x85, 0x90,
			0xA9, 0x12,
			0x85, 0x91,
			0xA9, 0b0011_1111,
			0x8D, 0x34, 0x12,
			0xA2, 0x03,
			0xA9, 0b1111_0000,
			0x21, 0x8D,
			0x00,
		})

		testA(t, cpu, 0b0011_0000)
	})
}

func Test0x31ANDIndirectY(t *testing.T) {
	t.Run("and", func(t *testing.T) {
		cpu := newCpu()

		cpu.LoadAndRun([]byte{
			0xA9, 0x32,
			0x85, 0x90,
			0xA9, 0x12,
			0x85, 0x91,
			0xA9, 0b0011_1111,
			0x8D, 0x34, 0x12,
			0xA0, 0x02,
			0xA9, 0b1111_0000,
			0x31, 0x90,
			0x00,
		})

		testA(t, cpu, 0b0011_0000)
	})
}

func Test0x0AASLAccumulator(t *testing.T) {
	t.Run("shift left", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0b0000_0001, 0x0A, 0x00})

		testA(t, cpu, 0b0000_0010)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("shift left with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0b1000_0000, 0x0A, 0x00})

		testA(t, cpu, 0b0000_0000)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x06ASLZeroPage(t *testing.T) {
	t.Run("shift left", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0b0000_0001)
		cpu.LoadAndRun([]byte{0x06, 0x05, 0x00})

		testMem(t, cpu, 0x05, 0b0000_0010)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("shift left with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0b1000_0000)
		cpu.LoadAndRun([]byte{0x06, 0x05, 0x00})

		testMem(t, cpu, 0x05, 0b0000_0000)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x16ASLZeroPageX(t *testing.T) {
	t.Run("shift left", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0b0000_0001)
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0x16, 0x04, 0x00})

		testMem(t, cpu, 0x05, 0b0000_0010)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("shift left with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0b1000_0000)
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0x16, 0x04, 0x00})

		testMem(t, cpu, 0x05, 0b0000_0000)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x0EASLAbsolute(t *testing.T) {
	t.Run("shift left", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0200, 0b0000_0001)
		cpu.LoadAndRun([]byte{0x0E, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0200, 0b0000_0010)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("shift left with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0200, 0b1000_0000)
		cpu.LoadAndRun([]byte{0x0E, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0200, 0b0000_0000)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x1EASLAbsoluteX(t *testing.T) {
	t.Run("shift left", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0b0000_0001)
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0x1E, 0x01, 0x02, 0x00})

		testMem(t, cpu, 0x0202, 0b0000_0010)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("shift left with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0b1000_0000)
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0x1E, 0x01, 0x02, 0x00})

		testMem(t, cpu, 0x0202, 0b0000_0000)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x4CJMPAbsolute(t *testing.T) {
	t.Run("jump", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x0200, 0xA9)
		cpu.Write(0x0201, 0x05)
		cpu.LoadAndRun([]byte{0x4C, 0x00, 0x02})

		testA(t, cpu, 0x05)
	})
}

func Test0x6CJMPIndirect(t *testing.T) {
	t.Run("jump", func(t *testing.T) {
		cpu := newCpu()

		cpu.LoadAndRun([]byte{
			0xA9, 0x34,
			0x8D, 0x00, 0x02,
			0xA9, 0x12,
			0x8D, 0x01, 0x02,
			0xA9, 0xA9,
			0x8D, 0x34, 0x12,
			0xA9, 0x55,
			0x8D, 0x35, 0x12,
			0x6C, 0x00, 0x02,
		})

		testA(t, cpu, 0x55)
	})

	t.Run("jump with page boundary", func(t *testing.T) {
		cpu := newCpu()

		// Set up the indirect jump address
		cpu.LoadAndRun([]byte{
			0xA9, 0x45, // LDA #$45
			0x85, 0xFF, // STA $FF
			0xA9, 0x02, // LDA #$02
			0x85, 0x00, // STA $00
			0xA9, 0xA9, // LDA #$A9
			0x8D, 0x45, 0x02, // STA $0245
			0xA9, 0x55, // LDA #$55
			0x8D, 0x46, 0x02, // STA $0246
			0xA9, 0xA9, // LDA #$A9
			0x8D, 0x45, 0x03, // STA $0345
			0xA9, 0x05, // LDA #$05
			0x8D, 0x46, 0x03, // STA $0346
			0x6C, 0xFF, 0x00, // JMP ($00FF)
			0x00, // BRK
		})

		testA(t, cpu, 0x55)
	})
}

func Test0x90BCC(t *testing.T) {
	t.Run("branch", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0x90, 0x02, 0x00, 0x00, 0xA9, 0x05, 0x00})

		testA(t, cpu, 0x05)
	})

	t.Run("no branch", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0xFD, 0x69, 0xFF, 0x90, 0x02, 0xA9, 0x05, 0x00})

		testA(t, cpu, 0x05)
	})
}

func Test0xB0BCS(t *testing.T) {
	t.Run("branch", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0xFF, 0x69, 0xFF, 0xB0, 0x02, 0x00, 0x00, 0xA9, 0x05, 0x00})

		testA(t, cpu, 0x05)
	})

	t.Run("no branch", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xB0, 0x02, 0xA9, 0x05, 0x00})

		testA(t, cpu, 0x05)
	})
}

func Test0xF0BEQ(t *testing.T) {
	t.Run("branch", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x00, 0xF0, 0x02, 0x00, 0x00, 0xA9, 0x05, 0x00})

		testA(t, cpu, 0x05)
	})

	t.Run("no branch", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x01, 0xF0, 0x02, 0xA9, 0x05, 0x00})

		testA(t, cpu, 0x05)
	})
}

func Test0x24BITZeroPage(t *testing.T) {
	t.Run("bit test", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0b1110_1010)
		cpu.LoadAndRun([]byte{0xA9, 0b1010_1010, 0x24, 0x05, 0x00})

		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, true)
		testFlag(t, cpu, OVERFLOW, true)
	})
}

func Test0x30BMI(t *testing.T) {
	t.Run("branch", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0xFF, 0x30, 0x02, 0x00, 0x00, 0xA9, 0x05, 0x00})

		testA(t, cpu, 0x05)
	})

	t.Run("no branch", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x01, 0x30, 0x02, 0xA9, 0x05, 0x00})

		testA(t, cpu, 0x05)
	})
}

func Test0xD0BNE(t *testing.T) {
	t.Run("branch", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x01, 0xD0, 0x02, 0x00, 0x00, 0xA9, 0x05, 0x00})

		testA(t, cpu, 0x05)
	})

	t.Run("no branch", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x00, 0xD0, 0x02, 0xA9, 0x05, 0x00})

		testA(t, cpu, 0x05)
	})
}

func Test0x10BPL(t *testing.T) {
	t.Run("branch", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x01, 0x10, 0x02, 0x00, 0x00, 0xA9, 0x05, 0x00})

		testA(t, cpu, 0x05)
	})

	t.Run("no branch", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0xFF, 0x10, 0x02, 0xA9, 0x05, 0x00})

		testA(t, cpu, 0x05)
	})
}

func Test0x50BVC(t *testing.T) {
	t.Run("branch", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x50, 0x69, 0x05, 0x50, 0x02, 0x00, 0x00, 0xA9, 0x05, 0x00})

		testA(t, cpu, 0x05)
	})

	t.Run("no branch", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x50, 0x69, 0x50, 0x50, 0x02, 0xA9, 0x05, 0x00})

		testA(t, cpu, 0x05)
	})
}

func Test0x70BVS(t *testing.T) {
	t.Run("branch", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x50, 0x69, 0x50, 0x70, 0x02, 0x00, 0x00, 0xA9, 0x05, 0x00})

		testA(t, cpu, 0x05)
	})

	t.Run("no branch", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x50, 0x69, 0x01, 0x70, 0x02, 0xA9, 0x05, 0x00})

		testA(t, cpu, 0x05)
	})
}

func Test0x18CLC(t *testing.T) {
	cpu := newCpu()
	cpu.setFlag(CARRY, true)
	cpu.LoadAndRun([]byte{0x18, 0x00})

	testFlag(t, cpu, CARRY, false)
}

func Test0xD8CLD(t *testing.T) {
	cpu := newCpu()
	cpu.setFlag(DECIMAL, true)
	cpu.LoadAndRun([]byte{0xD8, 0x00})

	testFlag(t, cpu, DECIMAL, false)
}

func Test0x58CLI(t *testing.T) {
	cpu := newCpu()
	cpu.setFlag(INTERRUPT, true)
	cpu.LoadAndRun([]byte{0x58, 0x00})

	testFlag(t, cpu, INTERRUPT, false)
}

func Test0xB8CLV(t *testing.T) {
	cpu := newCpu()
	cpu.setFlag(OVERFLOW, true)
	cpu.LoadAndRun([]byte{0xB8, 0x00})

	testFlag(t, cpu, OVERFLOW, false)
}

func Test0xC9CMPImmediate(t *testing.T) {
	t.Run("compare", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0xC9, 0x05, 0x00})

		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
		testFlag(t, cpu, CARRY, true)
	})
}

func Test0xC5ZeroPage(t *testing.T) {
	t.Run("compare", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x01, 0x05)
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0xC5, 0x01, 0x00})

		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
		testFlag(t, cpu, CARRY, true)
	})
}

func Test0xD5CMPZeroPageX(t *testing.T) {
	t.Run("compare", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x02, 0x05)
		cpu.LoadAndRun([]byte{
			0xA9, 0x05,
			0x85, 0x90,
			0xA2, 0x01,
			0xA9, 0x05,
			0xD5, 0x8F,
			0x00,
		})

		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
		testFlag(t, cpu, CARRY, true)
	})
}

func Test0xCDCMPAbsolute(t *testing.T) {
	t.Run("compare", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0200, 0x05)
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0xCD, 0x00, 0x02, 0x00})

		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
		testFlag(t, cpu, CARRY, true)
	})
}

func Test0xDDCMPAbsoluteX(t *testing.T) {
	t.Run("compare", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0x05)
		cpu.LoadAndRun([]byte{0xA2, 0x02, 0xA9, 0x05, 0xDD, 0x00, 0x02, 0x00})

		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
		testFlag(t, cpu, CARRY, true)
	})
}

func Test0xD9CMPAbsoluteY(t *testing.T) {
	t.Run("compare", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0x05)
		cpu.LoadAndRun([]byte{0xA0, 0x02, 0xA9, 0x05, 0xD9, 0x00, 0x02, 0x00})

		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
		testFlag(t, cpu, CARRY, true)
	})
}

func Test0xC1CMPIndirectX(t *testing.T) {
	t.Run("compare", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0x34)
		cpu.Write(0x06, 0x12)
		cpu.Write(0x1234, 0x05)

		cpu.LoadAndRun([]byte{
			0xA9, 0x34,
			0x85, 0x90,
			0xA9, 0x12,
			0x85, 0x91,
			0xA9, 0x05,
			0x8D, 0x34, 0x12,
			0xA2, 0x03,
			0xA9, 0x05,
			0xC1, 0x8D,
			0x00,
		})

		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
		testFlag(t, cpu, CARRY, true)
	})
}

func Test0xD1CMPIndirectY(t *testing.T) {
	t.Run("compare", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{
			0xA9, 0x32,
			0x85, 0x80,
			0xA9, 0x12,
			0x85, 0x81,
			0xA9, 0x05,
			0x8D, 0x34, 0x12,
			0xA0, 0x02,
			0xA9, 0x05,
			0xD1, 0x80,
			0x00,
		})

		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
		testFlag(t, cpu, CARRY, true)
	})
}

func Test0xE0CPXImmediate(t *testing.T) {
	t.Run("compare", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA2, 0x05, 0xE0, 0x05, 0x00})

		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
		testFlag(t, cpu, CARRY, true)
	})
}

func Test0xE4CPXZeroPage(t *testing.T) {
	t.Run("compare", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0x05)
		cpu.LoadAndRun([]byte{0xA2, 0x05, 0xE4, 0x05, 0x00})

		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
		testFlag(t, cpu, CARRY, true)
	})
}

func Test0xECCPXAbsolute(t *testing.T) {
	t.Run("compare", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0200, 0x05)
		cpu.LoadAndRun([]byte{0xA2, 0x05, 0xEC, 0x00, 0x02, 0x00})

		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
		testFlag(t, cpu, CARRY, true)
	})
}

func Test0xC0CPYImmediate(t *testing.T) {
	t.Run("compare", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA0, 0x05, 0xC0, 0x05, 0x00})

		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
		testFlag(t, cpu, CARRY, true)
	})
}

func Test0xC4CPYZeroPage(t *testing.T) {
	t.Run("compare", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0x05)
		cpu.LoadAndRun([]byte{0xA0, 0x05, 0xC4, 0x05, 0x00})

		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
		testFlag(t, cpu, CARRY, true)
	})
}

func Test0xCCCPYAbsolute(t *testing.T) {
	t.Run("compare", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x1234, 0x55)
		cpu.LoadAndRun([]byte{0xA0, 0x55, 0xCC, 0x34, 0x12, 0x00})

		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
		testFlag(t, cpu, CARRY, true)
	})

	t.Run("not equal", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x1234, 0x04)
		cpu.LoadAndRun([]byte{0xA0, 0x05, 0xCC, 0x34, 0x12, 0x00})

		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
		testFlag(t, cpu, CARRY, true)
	})
}

func Test0x49EORImmediate(t *testing.T) {
	t.Run("xor", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0b1010_1010, 0x49, 0b1100_1100, 0x00})

		testA(t, cpu, 0b0110_0110)
	})
}

func Test0x45EORZeroPage(t *testing.T) {
	t.Run("xor", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0b1100_1100)
		cpu.LoadAndRun([]byte{0xA9, 0b1010_1010, 0x45, 0x05, 0x00})

		testA(t, cpu, 0b0110_0110)
	})
}

func Test0x55EORZeroPageX(t *testing.T) {
	t.Run("xor", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{
			0xA9, 0b1100_1100,
			0x85, 0x90,
			0xA2, 0x01,
			0xA9, 0b1010_1010,
			0x55, 0x8F,
			0x00,
		})

		testA(t, cpu, 0b0110_0110)
	})
}

func Test0x4DEORAbsolute(t *testing.T) {
	t.Run("xor", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0200, 0b1100_1100)
		cpu.LoadAndRun([]byte{0xA9, 0b1010_1010, 0x4D, 0x00, 0x02, 0x00})

		testA(t, cpu, 0b0110_0110)
	})
}

func Test0x5DEORAbsoluteX(t *testing.T) {
	t.Run("xor", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0b1100_1100)
		cpu.LoadAndRun([]byte{0xA2, 0x02, 0xA9, 0b1010_1010, 0x5D, 0x00, 0x02, 0x00})

		testA(t, cpu, 0b0110_0110)
	})
}

func Test0x59EORAbsoluteY(t *testing.T) {
	t.Run("xor", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0b1100_1100)
		cpu.LoadAndRun([]byte{0xA0, 0x02, 0xA9, 0b1010_1010, 0x59, 0x00, 0x02, 0x00})

		testA(t, cpu, 0b0110_0110)
	})
}

func Test0x41EORIndirectX(t *testing.T) {
	t.Run("xor", func(t *testing.T) {
		cpu := newCpu()

		cpu.LoadAndRun([]byte{
			0xA9, 0x34,
			0x85, 0x90,
			0xA9, 0x12,
			0x85, 0x91,
			0xA9, 0b1100_1100,
			0x8D, 0x34, 0x12,
			0xA2, 0x03,
			0xA9, 0b1010_1010,
			0x41, 0x8D,
			0x00,
		})

		testA(t, cpu, 0b0110_0110)
	})
}

func Test0x51EORIndirectY(t *testing.T) {
	t.Run("xor", func(t *testing.T) {
		cpu := newCpu()

		cpu.LoadAndRun([]byte{
			0xA9, 0x32,
			0x85, 0x80,
			0xA9, 0x12,
			0x85, 0x81,
			0xA9, 0b1100_1100,
			0x8D, 0x34, 0x12,
			0xA0, 0x02,
			0xA9, 0b1010_1010,
			0x51, 0x80,
			0x00,
		})

		testA(t, cpu, 0b0110_0110)
	})
}

func Test0x20JSRAbsolute(t *testing.T) {
	t.Run("jump", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x0205, 0xA9)
		cpu.Write(0x0206, 0x05)
		cpu.Write(0x0207, 0x00)
		cpu.LoadAndRun([]byte{0x20, 0x05, 0x02})

		testA(t, cpu, 0x05)
		testStackAddress(t, cpu, 0xFB)
		testStack(t, cpu, 0x02)
	})
}

func Test0x60RTS(t *testing.T) {
	t.Run("return", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x0205, 0xA2)
		cpu.Write(0x0206, 0x05)
		cpu.Write(0x0207, 0x60)
		cpu.LoadAndRun([]byte{0x20, 0x05, 0x02, 0xA9, 0x05, 0x00})

		testA(t, cpu, 0x05)
		testX(t, cpu, 0x05)
		testStackAddress(t, cpu, 0xFD)
	})
}

func Test0x4ALSRAccumulator(t *testing.T) {
	t.Run("shift right", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0b0000_0010, 0x4A, 0x00})

		testA(t, cpu, 0b0000_0001)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("shift right with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0b0000_0001, 0x4A, 0x00})

		testA(t, cpu, 0b0000_0000)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x46LSRZeroPage(t *testing.T) {
	t.Run("shift right", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0b0000_0010)
		cpu.LoadAndRun([]byte{0x46, 0x05, 0x00})

		testMem(t, cpu, 0x05, 0b0000_0001)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("shift right with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0b0000_0001)
		cpu.LoadAndRun([]byte{0x46, 0x05, 0x00})

		testMem(t, cpu, 0x05, 0b0000_0000)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x56LSRZeroPageX(t *testing.T) {
	t.Run("shift right", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0b0000_0010)
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0x56, 0x04, 0x00})

		testMem(t, cpu, 0x05, 0b0000_0001)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("shift right with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0b0000_0001)
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0x56, 0x04, 0x00})

		testMem(t, cpu, 0x05, 0b0000_0000)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x4ELSRAbsolute(t *testing.T) {
	t.Run("shift right", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0200, 0b0000_0010)
		cpu.LoadAndRun([]byte{0x4E, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0200, 0b0000_0001)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("shift right with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0200, 0b0000_0001)
		cpu.LoadAndRun([]byte{0x4E, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0200, 0b0000_0000)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x5ELSRAbsoluteX(t *testing.T) {
	t.Run("shift right", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0b0000_0010)
		cpu.LoadAndRun([]byte{0xA2, 0x02, 0x5E, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0202, 0b0000_0001)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("shift right with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0b0000_0001)
		cpu.LoadAndRun([]byte{0xA2, 0x02, 0x5E, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0202, 0b0000_0000)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x09ORAImmediate(t *testing.T) {
	t.Run("or", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0b1010_1010, 0x09, 0b1100_1100, 0x00})

		testA(t, cpu, 0b1110_1110)
	})
}

func Test0x05ORAZeroPage(t *testing.T) {
	t.Run("or", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0b1100_1100)
		cpu.LoadAndRun([]byte{0xA9, 0b1010_1010, 0x05, 0x05, 0x00})

		testA(t, cpu, 0b1110_1110)
	})
}

func Test0x15ORAZeroPageX(t *testing.T) {
	t.Run("or", func(t *testing.T) {
		cpu := newCpu()

		cpu.LoadAndRun([]byte{
			0xA9, 0b1100_1100,
			0x85, 0x90,
			0xA2, 0x01,
			0xA9, 0b1010_1010,
			0x15, 0x8F,
			0x00,
		})

		testA(t, cpu, 0b1110_1110)
	})
}

func Test0x0DORAAbsolute(t *testing.T) {
	t.Run("or", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0200, 0b1100_1100)
		cpu.LoadAndRun([]byte{0xA9, 0b1010_1010, 0x0D, 0x00, 0x02, 0x00})

		testA(t, cpu, 0b1110_1110)
	})
}

func Test0x1DORAAbsoluteX(t *testing.T) {
	t.Run("or", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0b1100_1100)
		cpu.LoadAndRun([]byte{0xA2, 0x02, 0xA9, 0b1010_1010, 0x1D, 0x00, 0x02, 0x00})

		testA(t, cpu, 0b1110_1110)
	})
}

func Test0x19ORAAbsoluteY(t *testing.T) {
	t.Run("or", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0b1100_1100)
		cpu.LoadAndRun([]byte{0xA0, 0x02, 0xA9, 0b1010_1010, 0x19, 0x00, 0x02, 0x00})

		testA(t, cpu, 0b1110_1110)
	})
}

func Test0x01ORAIndirectX(t *testing.T) {
	t.Run("or", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0x34)
		cpu.Write(0x06, 0x12)
		cpu.Write(0x1234, 0b1100_1100)

		cpu.LoadAndRun([]byte{
			0xA9, 0x34,
			0x85, 0x90,
			0xA9, 0x12,
			0x85, 0x91,
			0xA9, 0b1100_1100,
			0x8D, 0x34, 0x12,
			0xA2, 0x03,
			0xA9, 0b1010_1010,
			0x01, 0x8D,
			0x00,
		})

		testA(t, cpu, 0b1110_1110)
	})
}

func Test0x11ORAIndirectY(t *testing.T) {
	t.Run("or", func(t *testing.T) {
		cpu := newCpu()

		cpu.LoadAndRun([]byte{
			0xA9, 0x32,
			0x85, 0x80,
			0xA9, 0x12,
			0x85, 0x81,
			0xA9, 0b1100_1100,
			0x8D, 0x34, 0x12,
			0xA0, 0x02,
			0xA9, 0b1010_1010,
			0x11, 0x80,
			0x00,
		})

		testA(t, cpu, 0b1110_1110)
	})
}

func Test0x2AROLAccumulator(t *testing.T) {
	t.Run("rotate left", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0b1000_0001, 0x2A, 0x00})

		testA(t, cpu, 0b0000_0010)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("rotate left with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0b1000_0000, 0x38, 0x2A, 0x00})

		testA(t, cpu, 0b0000_0001)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x26ROLZeroPage(t *testing.T) {
	t.Run("rotate left", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0b1000_0001)
		cpu.LoadAndRun([]byte{0x26, 0x05, 0x00})

		testMem(t, cpu, 0x05, 0b0000_0010)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("rotate left with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0b1000_0000)
		cpu.LoadAndRun([]byte{0x38, 0x26, 0x05, 0x00})

		testMem(t, cpu, 0x05, 0b0000_0001)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x36ROLZeroPageX(t *testing.T) {
	t.Run("rotate left", func(t *testing.T) {
		cpu := newCpu()

		cpu.LoadAndRun([]byte{
			0xA9, 0b1000_0001,
			0x85, 0x90,
			0xA2, 0x01,
			0x36, 0x8F,
			0x00,
		})

		testMem(t, cpu, 0x90, 0b0000_0010)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("rotate left with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0b1000_0000)
		cpu.LoadAndRun([]byte{
			0xA9, 0b1000_0000,
			0x85, 0x90,
			0xA2, 0x01,
			0x38,
			0x36, 0x8F,
			0x00,
		})

		testMem(t, cpu, 0x05, 0b0000_0001)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x2EROLAbsolute(t *testing.T) {
	t.Run("rotate left", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0200, 0b1000_0001)
		cpu.LoadAndRun([]byte{0x2E, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0200, 0b0000_0010)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("rotate left with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0200, 0b1000_0000)
		cpu.LoadAndRun([]byte{0x38, 0x2E, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0200, 0b0000_0001)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x3EROLAbsoluteX(t *testing.T) {
	t.Run("rotate left", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0b1000_0001)
		cpu.LoadAndRun([]byte{0xA2, 0x02, 0x3E, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0202, 0b0000_0010)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("rotate left with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0b1000_0000)
		cpu.LoadAndRun([]byte{0xA2, 0x02, 0x38, 0x3E, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0202, 0b0000_0001)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x38SEC(t *testing.T) {
	cpu := newCpu()
	cpu.LoadAndRun([]byte{0x38, 0x00})

	testFlag(t, cpu, CARRY, true)
}

func Test0x6ARORAccumulator(t *testing.T) {
	t.Run("rotate right", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0b0000_0010, 0x6A, 0x00})

		testA(t, cpu, 0b0000_0001)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("rotate right with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0b0000_0001, 0x38, 0x6A, 0x00})

		testA(t, cpu, 0b1000_0000)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, true)
	})
}

func Test0x66RORZeroPage(t *testing.T) {
	t.Run("rotate right", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0b0000_0010)
		cpu.LoadAndRun([]byte{0x66, 0x05, 0x00})

		testMem(t, cpu, 0x05, 0b0000_0001)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("rotate right with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0b0000_0001)
		cpu.LoadAndRun([]byte{0x38, 0x66, 0x05, 0x00})

		testMem(t, cpu, 0x05, 0b1000_0000)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, true)
	})
}

func Test0x76RORZeroPageX(t *testing.T) {
	t.Run("rotate right", func(t *testing.T) {
		cpu := newCpu()

		cpu.LoadAndRun([]byte{
			0xA9, 0b0000_0010,
			0x85, 0x90,
			0xA2, 0x01,
			0x76, 0x8F,
			0x00,
		})

		testMem(t, cpu, 0x90, 0b0000_0001)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("rotate right with carry", func(t *testing.T) {
		cpu := newCpu()

		cpu.LoadAndRun([]byte{
			0xA9, 0b0000_0001,
			0x85, 0x90,
			0xA2, 0x01,
			0x38,
			0x76, 0x8F,
			0x00,
		})

		testMem(t, cpu, 0x90, 0b1000_0000)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, true)
	})
}

func Test0x6ERORAbsolute(t *testing.T) {
	t.Run("rotate right", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0200, 0b0000_0010)
		cpu.LoadAndRun([]byte{0x6E, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0200, 0b0000_0001)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("rotate right with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0200, 0b0000_0001)
		cpu.LoadAndRun([]byte{0x38, 0x6E, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0200, 0b1000_0000)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, true)
	})
}

func Test0x7EAbsoluteX(t *testing.T) {
	t.Run("rotate right", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0b0000_0010)
		cpu.LoadAndRun([]byte{0xA2, 0x02, 0x7E, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0202, 0b0000_0001)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("rotate right with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0b0000_0001)
		cpu.LoadAndRun([]byte{0xA2, 0x02, 0x38, 0x7E, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0202, 0b1000_0000)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, true)
	})
}

func Test0x40RTI(t *testing.T) {
	t.Run("return from interrupt", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0200, 0xA9)
		cpu.Write16(0x0201, 0x05)
		cpu.Write16(0x0202, 0x00)
		cpu.LoadAndRun([]byte{
			0xA9, 0x02, // LDA #$02
			0x48,       // PHA
			0xA9, 0x00, // LDA #$00
			0x48,              // PHA
			0xA9, 0b1101_1111, // LDA #$FF
			0x48, // PHA
			0x40, // RTI
		})

		testStackAddress(t, cpu, 0xFD)
		testA(t, cpu, 0x05)
		testFlag(t, cpu, BREAK, false)
		testFlag(t, cpu, UNUSED, true)
	})
}

func Test0xE9SBCImmediate(t *testing.T) {
	t.Run("subtract with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0xE9, 0x02, 0x00})

		testA(t, cpu, 0x03)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("subtract with carry and borrow", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xA9, 0x02, 0xE9, 0x05, 0x00})

		testA(t, cpu, 0xFD)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, true)
	})
}

func Test0xE5SBCZeroPage(t *testing.T) {
	t.Run("subtract with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0x02)
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0xE5, 0x05, 0x00})

		testA(t, cpu, 0x03)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("subtract with carry and borrow", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x05, 0x05)
		cpu.LoadAndRun([]byte{0xA9, 0x02, 0xE5, 0x05, 0x00})

		testA(t, cpu, 0xFD)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, true)
	})
}

func Test0xF5SBCZeroPageX(t *testing.T) {
	t.Run("subtract with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x90, 0x02)
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0xA9, 0x05, 0xF5, 0x8F, 0x00})

		testA(t, cpu, 0x03)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("subtract with carry and borrow", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x90, 0x05)
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0xA9, 0x02, 0xF5, 0x8F, 0x00})

		testA(t, cpu, 0xFD)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, true)
	})
}

func Test0xEDSBCAbsolute(t *testing.T) {
	t.Run("subtract with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0200, 0x02)
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0xED, 0x00, 0x02, 0x00})

		testA(t, cpu, 0x03)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("subtract with carry and borrow", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0200, 0x05)
		cpu.LoadAndRun([]byte{0xA9, 0x02, 0xED, 0x00, 0x02, 0x00})

		testA(t, cpu, 0xFD)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, true)
	})
}

func Test0xFDSBCAbsoluteX(t *testing.T) {
	t.Run("subtract with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0x02)
		cpu.LoadAndRun([]byte{0xA2, 0x02, 0xA9, 0x05, 0xFD, 0x00, 0x02, 0x00})

		testA(t, cpu, 0x03)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("subtract with carry and borrow", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0x05)
		cpu.LoadAndRun([]byte{0xA2, 0x02, 0xA9, 0x02, 0xFD, 0x00, 0x02, 0x00})

		testA(t, cpu, 0xFD)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, true)
	})
}

func Test0xF9SBCAbsoluteY(t *testing.T) {
	t.Run("subtract with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0x02)
		cpu.LoadAndRun([]byte{0xA0, 0x02, 0xA9, 0x05, 0xF9, 0x00, 0x02, 0x00})

		testA(t, cpu, 0x03)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("subtract with carry and borrow", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write16(0x0202, 0x05)
		cpu.LoadAndRun([]byte{0xA0, 0x02, 0xA9, 0x02, 0xF9, 0x00, 0x02, 0x00})

		testA(t, cpu, 0xFD)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, true)
	})
}

func Test0xE1SBCIndirectX(t *testing.T) {
	t.Run("subtract with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x90, 0x34)
		cpu.Write(0x91, 0x12)
		cpu.Write(0x1234, 0x02)
		cpu.LoadAndRun([]byte{0xA2, 0x03, 0xA9, 0x05, 0xE1, 0x8D, 0x00})

		testA(t, cpu, 0x03)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xF1SBCIndirectY(t *testing.T) {
	t.Run("subtract with carry", func(t *testing.T) {
		cpu := newCpu()
		cpu.Write(0x90, 0x32)
		cpu.Write(0x91, 0x12)
		cpu.Write(0x1234, 0x02)
		cpu.LoadAndRun([]byte{0xA0, 0x02, 0xA9, 0x05, 0xF1, 0x90, 0x00})

		testA(t, cpu, 0x03)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xF8SED(t *testing.T) {
	cpu := newCpu()
	cpu.LoadAndRun([]byte{0xF8, 0x00})

	testFlag(t, cpu, DECIMAL, true)
}

func Test0x78SEI(t *testing.T) {
	cpu := newCpu()
	cpu.LoadAndRun([]byte{0x78, 0x00})

	testFlag(t, cpu, INTERRUPT, true)
}

func Test0xBATSX(t *testing.T) {
	t.Run("transfer x to stack pointer", func(t *testing.T) {
		cpu := newCpu()
		cpu.LoadAndRun([]byte{0xBA, 0x00})

		testStackAddress(t, cpu, 0xFD)
		testX(t, cpu, 0xFD)
	})
}

func Test0x8ATXA(t *testing.T) {
	cpu := newCpu()
	cpu.LoadAndRun([]byte{0xA2, 0x05, 0x8A, 0x00})

	testA(t, cpu, 0x05)
}

func Test0x9ATXS(t *testing.T) {
	cpu := newCpu()
	cpu.LoadAndRun([]byte{0xA2, 0x05, 0x9A, 0x00})

	testStackAddress(t, cpu, 0x05)
}

func Test0x98TYA(t *testing.T) {
	cpu := newCpu()
	cpu.LoadAndRun([]byte{0xA0, 0x05, 0x98, 0x00})

	testA(t, cpu, 0x05)
}

// TEST HELPERS

func testA(t *testing.T, cpu *Cpu, eq byte) {
	t.Helper()
	if cpu.A != eq {
		t.Errorf("Expected a register to be 0x%02X, got 0x%02X", eq, cpu.A)
	}
}

func testX(t *testing.T, cpu *Cpu, eq byte) {
	t.Helper()
	if cpu.X != eq {
		t.Errorf("Expected x register to be 0x%02X, got 0x%02X", eq, cpu.X)
	}
}

func testY(t *testing.T, cpu *Cpu, eq byte) {
	t.Helper()
	if cpu.Y != eq {
		t.Errorf("Expected y register to be 0x%02X, got 0x%02X", eq, cpu.Y)
	}
}

func testMem(t *testing.T, cpu *Cpu, addr uint16, eq byte) {
	t.Helper()
	if cpu.Read(addr) != eq {
		t.Errorf("Expected memory at 0x%04X to be 0x%02X, got 0x%02X", addr, eq, cpu.Read(addr))
	}
}

func testFlag(t *testing.T, cpu *Cpu, flag StatusFlag, eq bool) {
	t.Helper()
	if cpu.getFlag(flag) != eq {
		t.Errorf("Expected flag %b to be %v, got %v", flag, eq, cpu.getFlag(flag))
	}
}

func testStackAddress(t *testing.T, cpu *Cpu, addr byte) {
	t.Helper()
	if cpu.SP != addr {
		t.Errorf("Expected stack pointer to be 0x%02X, got 0x%02X", addr, cpu.SP)
	}
}

func testStack(t *testing.T, cpu *Cpu, eq byte) {
	t.Helper()
	stackItem := cpu.Read(STACK + uint16(cpu.SP) + 1)

	if stackItem != eq {
		t.Errorf("Expected stack item to be 0x%02X, got 0x%02X", eq, stackItem)
	}
}
