package syscalls

import "os/exec"

// Nasty hack to put terminal into raw mode
// Should use syscalls
func MakeRaw() {
	// Disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()

	// Do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

// Nasty hack to make the terminal sane
// Should use syscalls
func MakeSane() {
	// Disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "sane").Run()
}
