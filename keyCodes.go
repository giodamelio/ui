package ui

import "fmt"

type Key struct {
	Name string
	Code []byte
}

var keycodes = []Key{
	Key{"UP", []byte{27, 91, 65}},
	Key{"DOWN", []byte{27, 91, 66}},
	Key{"RIGHT", []byte{27, 91, 67}},
	Key{"LEFT", []byte{27, 91, 68}},
	Key{"ENTER", []byte{10, 0, 0}},
	Key{"j", []byte{106, 0, 0}},
	Key{"k", []byte{107, 0, 0}},
}

// Return a key object from the raw key code
func GetKey(b []byte) Key {
	for index := range keycodes {
		key := keycodes[index]
		if key.Code[0] == b[0] && key.Code[1] == b[1] && key.Code[2] == b[2] {
			return key
		}
	}
	fmt.Println(b)
	return Key{
		Name: "UNKNOWN",
		Code: b,
	}
}
