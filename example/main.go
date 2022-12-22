package main

import "github.com/LgoLgo/cogen"

func main() {
	cogen.CoverGen(
		cogen.WithTitle("How to generate an article cover by golang"),
		cogen.WithImagePath("example/gopher.png"),
		cogen.WithFontFilePath("font/zads.ttf"),
		cogen.WithFontRGB([]int{0, 0, 0}),
		cogen.WithCoverRGB([]int{250, 235, 215}),
		cogen.WithFontSize(100),
	)
}
