package ui

import (
	"github.com/giodamelio/ui/widgets"
)

// A View
type View struct {
	Name string

	// All of our widgets
	widgets []widgets.Widget
}

// Add a widget
func (view *View) AddWidget(widget widgets.Widget) {
	// Add our widget to the list
	view.widgets = append(view.widgets, widget)
}
