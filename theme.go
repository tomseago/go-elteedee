package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
)

type LTDTheme struct {
	fyne.Theme
}

func NewLTDTheme() *LTDTheme {
	t := &LTDTheme{
		Theme: theme.DarkTheme(),
	}

	return t
}

func (t *LTDTheme) Padding() int {
	return 8
}
