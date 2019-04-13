package main

import (
	"fmt"
	"fyne.io/fyne"

	"fyne.io/fyne/widget"
)

type Window struct {
	win fyne.Window

	ipLabel     *widget.Label
	statusLabel *widget.Label
}

func NewWindow(win fyne.Window) *Window {
	w := &Window{
		win: win,
	}

	win.SetContent(widget.NewTabContainer(
		widget.NewTabItem("Net", w.netTab()),
		widget.NewTabItem("Sys", w.sysTab()),
	))

	// win.SetContent(widget.NewVBox(
	// 	widget.NewLabel("Hello Fyne!"),
	// 	widget.NewButton("Quit", func() {
	// 		a.Quit()
	// 	})))

	return w
}

func (w *Window) netTab() fyne.CanvasObject {
	w.ipLabel = widget.NewLabel("1.1.1.1")
	w.statusLabel = widget.NewLabel("status")

	return widget.NewVBox(w.ipLabel, w.statusLabel)
}

func (w *Window) sysTab() fyne.CanvasObject {

	hi := widget.NewLabel("Hello")
	btnTap := widget.NewButton("Tap me!", func() {
		size := w.win.Canvas().Size()

		s := fmt.Sprintf("%d x %d", size.Width, size.Height)
		w.statusLabel.SetText(s)
	})

	btnExit := widget.NewButton("Exit", func() {
		w.win.Close()
	})

	return widget.NewVBox(hi, btnTap, btnExit)
}
