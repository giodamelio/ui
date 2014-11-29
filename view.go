package ui

import (
	"github.com/giodamelio/ui/widgets"
)

// A View displays a collection of widgets
type View struct {
	Name string

	// All of our widgets
	widgets []widgets.Widget
}

// AddWidget adds a widget to our view
func (view *View) AddWidget(widget widgets.Widget) {
	// Add our widget to the list
	view.widgets = append(view.widgets, widget)
}
