package main

import "fmt"

const RAM = 0x0000
const RAM_END = 0x1FFF
const PPU = 0x2000
const PPU_END = 0x3FFF

type Bus struct {
	cpu_vram [2048]byte
}

func newBus() *Bus {
	return &Bus{}
}

func (b *Bus) Read(addr uint16) byte {
	if addr >= RAM && addr <= RAM_END {
		mirrorAddr := addr & 0b00000111_11111111
		return b.cpu_vram[mirrorAddr]
	}
	if addr >= PPU && addr <= PPU_END {
		//mirrorAddr := addr & 0b00100000_00000111
		panic("PPU read not implemented yet")
	}
	fmt.Printf("Ignoring read from address: 0x%04X\n", addr)
	return 0
}

func (b *Bus) Write(addr uint16, data byte) {
	if addr >= RAM && addr <= RAM_END {
		mirrorAddr := addr & 0b00000111_11111111
		b.cpu_vram[mirrorAddr] = data
		return
	}
	if addr >= PPU && addr <= PPU_END {
		//mirrorAddr := addr & 0b00100000_00000111
		fmt.Printf("write to PPU address: 0x%04X\n", addr)
		panic("PPU write not implemented yet")
	}
	fmt.Printf("Ignoring write to address: 0x%04X\n", addr)
}
