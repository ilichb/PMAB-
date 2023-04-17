package main

import (
	"os"

	"gioui.org/app"
)

func main() {
	go func() {
		defer os.Exit(0)
		w := app.NewWindow()
		for range w.Events() {

		}
	}()

	app.Main()
}
