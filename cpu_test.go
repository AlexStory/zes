package main

import (
	"testing"
)

func Test0xA9LDAImmediate(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0x00})

		testA(t, cpu, 0x05)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("zero is set properly", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA9, 0x00, 0x00})

		testA(t, cpu, 0x00)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xA5LDAZeroPage(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0x55)
		cpu.LoadAndRun([]byte{0xA5, 0x05, 0x00})

		testA(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xB5LDAZeroPageX(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0x55)
		cpu.LoadAndRun([]byte{0xA2, 0x03, 0xB5, 0x02, 0x00})

		testA(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xADLDAAbsolute(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write16(0x0200, 0x55)
		cpu.LoadAndRun([]byte{0xAD, 0x00, 0x02, 0x00})

		testA(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xBDLDAAbsoluteX(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write16(0x0202, 0x55)
		cpu.LoadAndRun([]byte{0xA2, 0x02, 0xBD, 0x00, 0x02, 0x00})

		testA(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xB9LDAAbsoluteY(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write16(0x0202, 0x55)
		cpu.LoadAndRun([]byte{0xA0, 0x02, 0xB9, 0x00, 0x02, 0x00})

		testA(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xA1LDAIndirectX(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0x34)
		cpu.write(0x06, 0x12)
		cpu.write(0x1234, 0x55)

		cpu.LoadAndRun([]byte{0xA2, 0x03, 0xA1, 0x02, 0x00})

		testA(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xB1LDAIndirectY(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0x32)
		cpu.write(0x06, 0x12)
		cpu.write(0x1234, 0x55)
		cpu.LoadAndRun([]byte{0xA0, 0x02, 0xB1, 0x05, 0x00})

		testA(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x85STAZeroPage(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA9, 0x55, 0x85, 0x05, 0x00})

		testMem(t, cpu, 0x05, 0x55)
	})
}

func Test0x95STAZeroPageX(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA2, 0x03, 0xA9, 0x55, 0x95, 0x02, 0x00})

		testMem(t, cpu, 0x05, 0x55)
	})
}

func Test0x8DSTAAbsolute(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA9, 0x55, 0x8D, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0200, 0x55)
	})
}

func Test0x9DSTAAbsoluteX(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA2, 0x02, 0xA9, 0x55, 0x9D, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0202, 0x55)
	})
}

func Test0x81STAIndirectX(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0x34)
		cpu.write(0x06, 0x12)
		cpu.LoadAndRun([]byte{0xA2, 0x03, 0xA9, 0x55, 0x81, 0x02, 0x00})

		testMem(t, cpu, 0x1234, 0x55)
	})
}

func Test0x91STAIndirectY(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0x32)
		cpu.write(0x06, 0x12)
		cpu.LoadAndRun([]byte{0xA0, 0x02, 0xA9, 0x55, 0x91, 0x05, 0x00})

		testMem(t, cpu, 0x1234, 0x55)
	})
}

func Test0x99STAAbsoluteY(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA0, 0x02, 0xA9, 0x55, 0x99, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0202, 0x55)
	})
}

func Test0xAATAX(t *testing.T) {
	t.Run("transefer a to x", func(t *testing.T) {
		cpu := new_cpu()
		program := []byte{0xA9, 0x05, 0xAA, 0x00}

		cpu.LoadAndRun(program)

		testA(t, cpu, 0x05)
		testX(t, cpu, 0x05)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("zero is set properly", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA9, 0x00, 0xAA, 0x00})

		testX(t, cpu, 0x00)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xE8INX(t *testing.T) {
	t.Run("increment x", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0xAA, 0xE8, 0x00})

		testA(t, cpu, 0x05)
		testX(t, cpu, 0x06)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("Test x overflow", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA2, 0xFF, 0xE8, 0xE8, 0x00})

		testX(t, cpu, 0x01)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xA2LDXImmediate(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA2, 0x05, 0x00})

		testX(t, cpu, 0x05)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("zero is set properly", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA2, 0x00, 0x00})

		testX(t, cpu, 0x00)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xA6LDXZeroPage(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0x55)
		cpu.LoadAndRun([]byte{0xA6, 0x05, 0x00})

		testX(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xB6LDXZeroPageY(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0x55)
		cpu.LoadAndRun([]byte{0xA0, 0x01, 0xB6, 0x04, 0x00})

		testX(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xAELDXAbsolute(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write16(0x0200, 0x55)
		cpu.LoadAndRun([]byte{0xAE, 0x00, 0x02, 0x00})

		testX(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xBELDXAbsoluteY(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write16(0x0202, 0x55)
		cpu.LoadAndRun([]byte{0xA0, 0x02, 0xBE, 0x00, 0x02, 0x00})

		testX(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xA0LDYImmediate(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA0, 0x05, 0x00})

		testY(t, cpu, 0x05)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xA4LDYZeroPage(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0x55)
		cpu.LoadAndRun([]byte{0xA4, 0x05, 0x00})

		testY(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xB4LDYZeroPageX(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0x55)
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0xB4, 0x04, 0x00})

		testY(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xACLDYAbsolute(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write16(0x0200, 0x55)
		cpu.LoadAndRun([]byte{0xAC, 0x00, 0x02, 0x00})

		testY(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xBCLDYAbsoluteX(t *testing.T) {
	t.Run("load data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write16(0x0202, 0x55)
		cpu.LoadAndRun([]byte{0xA2, 0x02, 0xBC, 0x00, 0x02, 0x00})

		testY(t, cpu, 0x55)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func TestSimpleProgram(t *testing.T) {
	cpu := new_cpu()
	program := []byte{0xA9, 0xc0, 0xAA, 0xE8, 0x00}
	cpu.LoadAndRun(program)

	testA(t, cpu, 0xC0)
	testX(t, cpu, 0xC1)
}

func Test0x86STXZeroPage(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA2, 0x05, 0x86, 0x05, 0x00})

		testMem(t, cpu, 0x05, 0x05)
	})
}

func Test0x96STXZeroPageY(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA0, 0x01, 0xA2, 0x05, 0x96, 0x04, 0x00})

		testMem(t, cpu, 0x05, 0x05)
	})
}

func Test0x8ESTXAbsolute(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA2, 0x05, 0x8E, 0x00, 0x02, 0x00})
		testMem(t, cpu, 0x0200, 0x05)
	})
}

func Test0x84STYZeroPage(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA0, 0x05, 0x84, 0x05, 0x00})

		testMem(t, cpu, 0x05, 0x05)
	})
}

func Test0x94STYZeroPageX(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0xA0, 0x05, 0x94, 0x04, 0x00})

		testMem(t, cpu, 0x05, 0x05)
	})
}

func Test0x8CTestSTYAbsolute(t *testing.T) {
	t.Run("store data", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA0, 0x05, 0x8C, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0200, 0x05)
	})
}

func Test0xA8TAY(t *testing.T) {
	t.Run("transfer a to y", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0xA8, 0x00})

		testY(t, cpu, 0x05)
	})
}

func Test0xC8INY(t *testing.T) {
	t.Run("increment y", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA0, 0x05, 0xC8, 0x00})

		testY(t, cpu, 0x06)
	})
}

func Test0xE6INCZeroPage(t *testing.T) {
	t.Run("increment memory", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0x55)
		cpu.LoadAndRun([]byte{0xE6, 0x05, 0x00})

		testMem(t, cpu, 0x05, 0x56)
	})
}

func Test0xF6INCZeroPageX(t *testing.T) {
	t.Run("increment memory", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0x55)
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0xF6, 0x04, 0x00})

		testMem(t, cpu, 0x05, 0x56)
	})
}

func Test0xEEINCAbsolute(t *testing.T) {
	t.Run("increment memory", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write16(0x0200, 0x55)
		cpu.LoadAndRun([]byte{0xEE, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0200, 0x56)
	})
}

func Test0xFEINCAbsoluteX(t *testing.T) {
	t.Run("increment memory", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write16(0x01234, 0x55)
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0xFE, 0x33, 0x12, 0x00})

		testMem(t, cpu, 0x1234, 0x56)
	})
}

func Test0xCADEX(t *testing.T) {
	t.Run("decrement x", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA2, 0x05, 0xCA, 0x00})

		testX(t, cpu, 0x04)
	})
}

func Test0x88DEY(t *testing.T) {
	t.Run("decrement y", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA0, 0x05, 0x88, 0x00})

		testY(t, cpu, 0x04)
	})
}

func Test0xC6DECZeroPage(t *testing.T) {
	t.Run("decrement memory", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0x55)
		cpu.LoadAndRun([]byte{0xC6, 0x05, 0x00})

		testMem(t, cpu, 0x05, 0x54)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0xEANOP(t *testing.T) {
	t.Run("do nothing", func(t *testing.T) {
		cpu := new_cpu()
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
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xEA, 0xA9, 0x05, 0x00})

		testA(t, cpu, 0x05)
	})
}

func Test0x48PHA(t *testing.T) {
	t.Run("push a to stack", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0x48, 0x00})

		testStack(t, cpu, 0x05)
		testStackAddress(t, cpu, 0xFC)
	})
}

func Test0x68PLA(t *testing.T) {
	t.Run("pop a from stack", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0x48, 0xA9, 0x00, 0x68, 0x00})

		testA(t, cpu, 0x05)
	})
}

func Test0x08PHP(t *testing.T) {
	t.Run("push status to stack", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA9, 0x00, 0x08, 0x00})

		testStack(t, cpu, 0x02)
		testStackAddress(t, cpu, 0xFC)
	})
}

func Test0x28PLP(t *testing.T) {
	t.Run("pop status from stack", func(t *testing.T) {
		cpu := new_cpu()
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
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0x69, 0x02, 0x00})

		testA(t, cpu, 0x07)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("add with carry carry bit", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA9, 0xFF, 0x69, 0x01, 0x00})

		testA(t, cpu, 0x00)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("add with signed overflow", func(t *testing.T) {
		cpu := new_cpu()
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
		cpu := new_cpu()
		cpu.write(0x05, 0x02)
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0x65, 0x05, 0x00})

		testA(t, cpu, 0x07)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x75ADCZeroPageX(t *testing.T) {
	t.Run("add with carry", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0x02)
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0xA9, 0x05, 0x75, 0x04, 0x00})

		testA(t, cpu, 0x07)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x6DADCAbsolute(t *testing.T) {
	t.Run("add with carry", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write16(0x0200, 0x02)
		cpu.LoadAndRun([]byte{0xA9, 0x05, 0x6D, 0x00, 0x02, 0x00})

		testA(t, cpu, 0x07)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x7DADCAbsoluteX(t *testing.T) {
	t.Run("add with carry", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write16(0x0202, 0x02)
		cpu.LoadAndRun([]byte{0xA2, 0x02, 0xA9, 0x05, 0x7D, 0x00, 0x02, 0x00})

		testA(t, cpu, 0x07)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x79ADCAbsoluteY(t *testing.T) {
	t.Run("add with carry", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write16(0x0202, 0x02)
		cpu.LoadAndRun([]byte{0xA0, 0x02, 0xA9, 0x05, 0x79, 0x00, 0x02, 0x00})

		testA(t, cpu, 0x07)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x61ADCIndirectX(t *testing.T) {
	t.Run("add with carry", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0x34)
		cpu.write(0x06, 0x12)
		cpu.write(0x1234, 0x02)

		cpu.LoadAndRun([]byte{0xA2, 0x03, 0xA9, 0x05, 0x61, 0x02, 0x00})

		testA(t, cpu, 0x07)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x71ADCIndirectY(t *testing.T) {
	t.Run("add with carry", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0x32)
		cpu.write(0x06, 0x12)
		cpu.write(0x1234, 0x02)
		cpu.LoadAndRun([]byte{0xA0, 0x02, 0xA9, 0x05, 0x71, 0x05, 0x00})

		testA(t, cpu, 0x07)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x29ANDImmediate(t *testing.T) {
	t.Run("and", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA9, 0b1111_0000, 0x29, 0b0011_1111, 0x00})

		testA(t, cpu, 0b0011_0000)
	})
}

func Test0x25ANDZeroPage(t *testing.T) {
	t.Run("and", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0b0011_1111)
		cpu.LoadAndRun([]byte{0xA9, 0b1111_0000, 0x25, 0x05, 0x00})

		testA(t, cpu, 0b0011_0000)
	})
}

func Test0x35ANDZeroPageX(t *testing.T) {
	t.Run("and", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0b0011_1111)
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0xA9, 0b1111_0000, 0x35, 0x04, 0x00})

		testA(t, cpu, 0b0011_0000)
	})
}

func Test0x2DANDAbsolute(t *testing.T) {
	t.Run("and", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write16(0x0200, 0b0011_1111)
		cpu.LoadAndRun([]byte{0xA9, 0b1111_0000, 0x2D, 0x00, 0x02, 0x00})

		testA(t, cpu, 0b0011_0000)
	})
}

func Test0x3DANDAbsoluteX(t *testing.T) {
	t.Run("and", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write16(0x0202, 0b0011_1111)
		cpu.LoadAndRun([]byte{0xA2, 0x02, 0xA9, 0b1111_0000, 0x3D, 0x00, 0x02, 0x00})

		testA(t, cpu, 0b0011_0000)
	})
}

func Test0x39ANDAbsoluteY(t *testing.T) {
	t.Run("and", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write16(0x0202, 0b0011_1111)
		cpu.LoadAndRun([]byte{0xA0, 0x02, 0xA9, 0b1111_0000, 0x39, 0x00, 0x02, 0x00})

		testA(t, cpu, 0b0011_0000)
	})
}

func Test0x21ANDIndirectX(t *testing.T) {
	t.Run("and", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0x34)
		cpu.write(0x06, 0x12)
		cpu.write(0x1234, 0b0011_1111)

		cpu.LoadAndRun([]byte{0xA2, 0x03, 0xA9, 0b1111_0000, 0x21, 0x02, 0x00})

		testA(t, cpu, 0b0011_0000)
	})
}

func Test0x31ANDIndirectY(t *testing.T) {
	t.Run("and", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0x32)
		cpu.write(0x06, 0x12)
		cpu.write(0x1234, 0b0011_1111)
		cpu.LoadAndRun([]byte{0xA0, 0x02, 0xA9, 0b1111_0000, 0x31, 0x05, 0x00})

		testA(t, cpu, 0b0011_0000)
	})
}

func Test0x0AASLAccumulator(t *testing.T) {
	t.Run("shift left", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA9, 0b0000_0001, 0x0A, 0x00})

		testA(t, cpu, 0b0000_0010)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("shift left with carry", func(t *testing.T) {
		cpu := new_cpu()
		cpu.LoadAndRun([]byte{0xA9, 0b1000_0000, 0x0A, 0x00})

		testA(t, cpu, 0b0000_0000)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x06ASLZeroPage(t *testing.T) {
	t.Run("shift left", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0b0000_0001)
		cpu.LoadAndRun([]byte{0x06, 0x05, 0x00})

		testMem(t, cpu, 0x05, 0b0000_0010)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("shift left with carry", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0b1000_0000)
		cpu.LoadAndRun([]byte{0x06, 0x05, 0x00})

		testMem(t, cpu, 0x05, 0b0000_0000)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x16ASLZeroPageX(t *testing.T) {
	t.Run("shift left", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0b0000_0001)
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0x16, 0x04, 0x00})

		testMem(t, cpu, 0x05, 0b0000_0010)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("shift left with carry", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x05, 0b1000_0000)
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0x16, 0x04, 0x00})

		testMem(t, cpu, 0x05, 0b0000_0000)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x0EASLAbsolute(t *testing.T) {
	t.Run("shift left", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write16(0x0200, 0b0000_0001)
		cpu.LoadAndRun([]byte{0x0E, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0200, 0b0000_0010)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("shift left with carry", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write16(0x0200, 0b1000_0000)
		cpu.LoadAndRun([]byte{0x0E, 0x00, 0x02, 0x00})

		testMem(t, cpu, 0x0200, 0b0000_0000)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x1EASLAbsoluteX(t *testing.T) {
	t.Run("shift left", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write16(0x0202, 0b0000_0001)
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0x1E, 0x01, 0x02, 0x00})

		testMem(t, cpu, 0x0202, 0b0000_0010)
		testFlag(t, cpu, CARRY, false)
		testFlag(t, cpu, ZERO, false)
		testFlag(t, cpu, NEGATIVE, false)
	})

	t.Run("shift left with carry", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write16(0x0202, 0b1000_0000)
		cpu.LoadAndRun([]byte{0xA2, 0x01, 0x1E, 0x01, 0x02, 0x00})

		testMem(t, cpu, 0x0202, 0b0000_0000)
		testFlag(t, cpu, CARRY, true)
		testFlag(t, cpu, ZERO, true)
		testFlag(t, cpu, NEGATIVE, false)
	})
}

func Test0x4CJMPAbsolute(t *testing.T) {
	t.Run("jump", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x0200, 0xA9)
		cpu.write(0x0201, 0x05)
		cpu.LoadAndRun([]byte{0x4C, 0x00, 0x02})

		testA(t, cpu, 0x05)
	})
}

func Test0x6CJMPIndirect(t *testing.T) {
	t.Run("jump", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x0200, 0x67)
		cpu.write(0x0201, 0x45)
		cpu.write(0x4567, 0xA9)
		cpu.write(0x4568, 0x55)
		cpu.LoadAndRun([]byte{0x6C, 0x00, 0x02})

		testA(t, cpu, 0x55)
	})

	t.Run("jump with page boundary", func(t *testing.T) {
		cpu := new_cpu()
		cpu.write(0x0200, 0x67)
		cpu.write(0x02FF, 0x45)
		cpu.write(0x0300, 0xA9)
		cpu.write(0x6745, 0xA9)
		cpu.write(0x6746, 0x55)
		cpu.write(0xA945, 0xA9)
		cpu.write(0xA946, 0x05)
		cpu.LoadAndRun([]byte{0x6C, 0xFF, 0x02})

		testA(t, cpu, 0x55)
	})
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
	if cpu.read(addr) != eq {
		t.Errorf("Expected memory at 0x%04X to be 0x%02X, got 0x%02X", addr, eq, cpu.read(addr))
	}
}

func testFlag(t *testing.T, cpu *Cpu, flag StatusFlag, eq bool) {
	t.Helper()
	if cpu.getFlag(flag) != eq {
		t.Errorf("Expected flag %v to be %v, got %v", flag, eq, cpu.getFlag(flag))
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
	stackItem := cpu.read(STACK + uint16(cpu.SP) + 1)

	if stackItem != eq {
		t.Errorf("Expected stack item to be 0x%02X, got 0x%02X", eq, stackItem)
	}
}
