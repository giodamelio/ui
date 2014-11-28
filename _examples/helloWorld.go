package main

import (
	"log"
	"os"

	"github.com/giodamelio/ui"
	"github.com/giodamelio/ui/widgets"
)

func main() {
	// Setup logging
	f, err := os.OpenFile("log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panicf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("- Logger Setup ------------------------------------")

	// Create our app
	app := ui.UI{}

	// Create our view
	view := ui.View{
		Name: "Main",
	}

	// Create a simple text label
	label := widgets.Label{
		X:    10,
		Y:    10,
		Text: "Hello World",
	}
	view.AddWidget(&label)

	// Add our view to the app
	app.AddView(view)

	// Show the menu
	app.ShowView("Main")

	// Run the app
	app.Run()
}
