package widgets

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

// A Label widget
type Label struct {
	Text string
	X    int
	Y    int
}

// Draw the label
func (label *Label) Draw(event termbox.Event, terminfo Terminfo) {
	fmt.Printf("%s%d;%dH", "\033[", label.X, label.Y)
	fmt.Print(label.Text)
}
