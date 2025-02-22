package main

type Memory interface {
	Read(addr uint16) byte
	Write(addr uint16, data byte)
}
