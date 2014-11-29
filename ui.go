package ui

import (
	"os"

	"github.com/giodamelio/ui/widgets"
	"github.com/nsf/termbox-go"
)

// UI type
type UI struct {
	// A list of our views
	views map[string]View

	// Our current view
	currentView string
}

// AddView adds a view to our UI
func (ui *UI) AddView(view View) {
	// If our views list does not exist, create it
	if ui.views == nil {
		ui.views = make(map[string]View)
	}

	// Add view to the map
	ui.views[view.Name] = view
}

// ShowView shows a view by name
func (ui *UI) ShowView(viewName string) {
	ui.currentView = viewName
}

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

	// Draw the current view once at startup
	terminfo := getTermInfo()
	ui.drawCurrentView(termbox.Event{}, terminfo)

	// Run the main loop
	for {
		// Wait for an event
		event := termbox.PollEvent()

		// If C-c is pressed, cleanup the terminal and exit
		if event.Type == termbox.EventKey && event.Key == termbox.KeyCtrlC {
			cleanExit()
		}

		// Get some info about the terminal
		terminfo = getTermInfo()

		// Redraw the current view
		ui.drawCurrentView(event, terminfo)
	}
}

// Draw the current view
func (ui *UI) drawCurrentView(event termbox.Event, terminfo widgets.Terminfo) {
	// Get the current view
	currentView := ui.views[ui.currentView]

	// Loop through the views widgets and draw them
	for _, widget := range currentView.widgets {
		widget.Draw(event, terminfo)
	}
}

// Get some info about the terminal
func getTermInfo() widgets.Terminfo {
	width, height := termbox.Size()
	return widgets.Terminfo{
		Width:  width,
		Height: height,
	}
}

// Exit cleanly
func cleanExit() {
	termbox.Close()
	os.Exit(0)
}
