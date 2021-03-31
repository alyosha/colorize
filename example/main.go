package main

import (
	"fmt"

	"github.com/alyosha/colorize"
	"golang.org/x/image/colornames"
)

func main() {
	fgColorizer := colorize.New(colornames.Darkcyan, nil)
	bgColorizer := colorize.New(nil, colornames.Darkgoldenrod)
	fgBgColorizer := colorize.New(colornames.Firebrick, colornames.Mediumaquamarine)

	fgColorizer.Println("println - foreground only, dark cyan")
	bgColorizer.Println("println - background only, dark goldenrod")
	fgBgColorizer.Println("println - foreground + background, firebrick - medium aquamarine")

	fgColorizer.UpdateForeground(colornames.Dodgerblue)
	bgColorizer.UpdateBackground(colornames.Crimson)
	fgBgColorizer.UpdateForeground(colornames.Blueviolet)
	fgBgColorizer.UpdateBackground(colornames.Indianred)

	fgColorizer.Printf("printf - foreground only, dodger blue: %s\n", "arg")
	bgColorizer.Printf("printf - background only, crimson: %s\n", "arg")
	fgBgColorizer.Printf("printf - foreground + background, plum - medium aquamarine: %s\n", "arg")

	fgColorizer.UpdateForeground(colornames.Steelblue)
	bgColorizer.UpdateBackground(colornames.Mediumspringgreen)
	fgBgColorizer.UpdateForeground(colornames.Teal)
	fgBgColorizer.UpdateBackground(colornames.Violet)

	fgColorizer.Print("print - foreground only, steel blue")
	bgColorizer.Print("print - background only, medium spring green")
	fgBgColorizer.Print("print - foreground + background, teal - violet")

	fmt.Printf(fgColorizer.Sprintf("\nsprintf - foreground only: %s\n", "arg"))
	fmt.Println(string(bgColorizer.Bytes([]byte("bytes - background only"))))
}
