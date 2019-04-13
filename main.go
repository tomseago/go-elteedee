package main

import (
	"flag"
	"time"

	"fyne.io/fyne"
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

	win.Resize(fyne.Size{480,320})
	if *fullscreen {
		k := func() {
			time.Sleep(1 * time.Second)
			win.SetFullScreen(true)
		}
		go k()
	}

	win.ShowAndRun()
}
