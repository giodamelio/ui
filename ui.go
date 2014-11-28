package ui

import (
	"os"

	"github.com/nsf/termbox-go"
)

// UI type
type UI struct{}

// Run the app
func (ui *UI) Run() {
	// Setup termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Put the terminal in raw mode
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	// Run the main loop
	for {
		event := termbox.PollEvent()
		switch event.Type {
		case termbox.EventKey:
			// If C-c is pressed, cleanup the terminal and exit
			if event.Key == termbox.KeyCtrlC {
				cleanExit()
			}
		}
	}
}

// Exit cleanly
func cleanExit() {
	termbox.Close()
	os.Exit(0)
}
