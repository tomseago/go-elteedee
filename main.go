package main

import (
	"flag"
	"github.com/eyethereal/go-config"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

//////////////////////////////////////////////////////////////////////
// Package level logging
// This only needs to be in one file per-package

var log = config.Logger("elteedee")

var gWindow *Window
var gSysMon *SysMon

func main() {
	// Setup flags
	fullscreen := flag.Bool("fullscreen", false, "Be fullscreen or not")

	flag.Parse()

	// Build the app
	a := app.New()

	theme := NewLTDTheme()
	a.Settings().SetTheme(theme)

	win := a.NewWindow("ElTeeDee")

	sysMon := NewSysMon()
	w := NewWindow(win)
	gWindow = w
	gSysMon = sysMon

	win.Resize(fyne.Size{480, 320})
	if *fullscreen {
		k := func() {
			time.Sleep(1 * time.Second)
			win.SetFullScreen(true)
		}
		go k()
	}

	go gSysMon.Run()

	win.ShowAndRun()
}
