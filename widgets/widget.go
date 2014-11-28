package widgets

import "github.com/nsf/termbox-go"

// A widget
type Widget interface {
	Draw(event termbox.Event, terminfo Terminfo)
}

// Some information about the terminal
type Terminfo struct {
	Width  int
	Height int
}
