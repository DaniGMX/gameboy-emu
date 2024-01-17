package main

import (
	"gameboy-emu/pkg/emu"
	"os"
)

func main() {
	args := os.Args[1:]
	argc := len(args)
	ctx := emu.Context{
		Running: false,
		Paused:  true,
		Ticks:   0,
	}

	emu.RunEmulator(ctx, argc, args)
}
