package ui

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/giodamelio/ui/syscalls"
)

// UI type
type UI struct{}

// Run the app
func (ui *UI) Run() {
	// Listen for interrupt so we can exit cleanly
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		syscalls.MakeSane()
		os.Exit(0)
	}()

	// Prep the terminal
	syscalls.MakeRaw()

	// Run the main loop
	mainLoop()
}

// The main loop
func mainLoop() {
	// Read keyboard input
	keys := make(chan Key)
	go readKeyboard(keys)

	for key := range keys {
		fmt.Println(key)
	}
}

// Read keypresses from stdin
func readKeyboard(keys chan Key) {
	// Buffer for the key
	var b = make([]byte, 3)

	// Loop forever reading keys
	for {
		// Read key from stdin
		os.Stdin.Read(b)

		// Send the key over the channel
		key := GetKey(b)
		keys <- key

		// Empty the buffer
		b[0] = 0
		b[1] = 0
		b[2] = 0
	}
}
