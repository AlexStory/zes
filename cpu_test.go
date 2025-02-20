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
