package emu

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Context struct {
	Paused  bool
	Running bool
	Ticks   uint64
}

func Delay(ms uint32) {
	sdl.Delay(ms)
}

func RunEmulator(ctx Context, argc int, argv []string) int {
	if argc < 2 {
		fmt.Printf("Usage: gb_emu <rom_file>\n")
		return -1
	}

	if sdlInitVideoErr := sdl.Init(sdl.INIT_VIDEO); sdlInitVideoErr != nil {
		panic(sdlInitVideoErr)
	}
	defer sdl.Quit()
	fmt.Print("SDL init...")

	if ttfInitError := ttf.Init(); ttfInitError != nil {
		panic(ttfInitError)
	}
	fmt.Print("TTF init...")

	ctx.Running = true
	ctx.Paused = false
	ctx.Ticks = 0

	for ctx.Running {
		if ctx.Paused {
			Delay(10)
			continue
		}

		ctx.Ticks++
	}

	return 0
}
