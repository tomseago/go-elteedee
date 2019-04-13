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
	timeLabel   *widget.Label

	cpus []*widget.ProgressBar
}

func NewWindow(win fyne.Window) *Window {
	w := &Window{
		win: win,

		cpus: make([]*widget.ProgressBar, 0),
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
	w.timeLabel = widget.NewLabel("time")

	return widget.NewVBox(w.ipLabel, w.statusLabel, w.timeLabel)
}

func (w *Window) sysTab() fyne.CanvasObject {

	numCPUs := gSysMon.NumCPUs()
	cpuBox := widget.NewHBox()

	if numCPUs > 0 {
		for i := 0; i < numCPUs; i++ {
			pb := widget.NewProgressBar()
			pb.Max = 100
			w.cpus = append(w.cpus, pb)
			cpuBox.Append(pb)
		}
	} else {
		cpuBox.Append(widget.NewLabel("CPU count is 0?"))
	}

	hi := widget.NewLabel("Hello")
	btnTap := widget.NewButton("Tap me!", func() {
		size := w.win.Canvas().Size()

		s := fmt.Sprintf("%d x %d", size.Width, size.Height)
		w.statusLabel.SetText(s)
	})

	btnExit := widget.NewButton("Exit", func() {
		w.win.Close()
	})

	return widget.NewVBox(cpuBox, hi, btnTap, btnExit)
}

func (w *Window) SetTime(s string) {
	w.timeLabel.SetText(s)
}

func (w *Window) SetCPU(p []float64) {
	for ix, cpu := range w.cpus {
		cpu.SetValue(p[ix])
	}
}
