package main

import (
	"flag"
	"fyne.io/fyne/app"
)

func main() {
	// Setup flags
	fullscreen := flag.Bool("fullscreen", false, "Be fullscreen or not")

	flag.Parse()

	// Build the app
	a := app.New()

	win := a.NewWindow("ElTeeDee")

	sysMon := NewSysMon()
	w := NewWindow(win)

	_ = sysMon
	_ = w

	if *fullscreen {
		win.SetFullScreen(true)
	}

	win.ShowAndRun()
}
