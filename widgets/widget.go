package widgets

import "github.com/nsf/termbox-go"

// A widget interface
type Widget interface {
	Draw(event termbox.Event, terminfo Terminfo)
}

// Terminfo stores information about the terminal
type Terminfo struct {
	Width  int
	Height int
}
