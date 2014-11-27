package ui

// A single key
type Key struct {
	Name string
	Code []byte
}

// All the keys
var keycodes = []Key{
	// Special keys
	Key{"TAB", []byte{9, 0, 0}},
	Key{"UP", []byte{27, 91, 65}},
	Key{"DOWN", []byte{27, 91, 66}},
	Key{"RIGHT", []byte{27, 91, 67}},
	Key{"LEFT", []byte{27, 91, 68}},
	Key{"SPACE", []byte{32, 0, 0}},
	Key{"ENTER", []byte{10, 0, 0}},

	// Numbers
	Key{"0", []byte{48, 0, 0}},
	Key{"1", []byte{49, 0, 0}},
	Key{"2", []byte{50, 0, 0}},
	Key{"3", []byte{51, 0, 0}},
	Key{"4", []byte{52, 0, 0}},
	Key{"5", []byte{53, 0, 0}},
	Key{"6", []byte{54, 0, 0}},
	Key{"7", []byte{55, 0, 0}},
	Key{"8", []byte{56, 0, 0}},
	Key{"9", []byte{57, 0, 0}},

	// Symbols
	Key{"`", []byte{96, 0, 0}},
	Key{"~", []byte{126, 0, 0}},
	Key{"!", []byte{33, 0, 0}},
	Key{"@", []byte{64, 0, 0}},
	Key{"#", []byte{35, 0, 0}},
	Key{"$", []byte{36, 0, 0}},
	Key{"%", []byte{37, 0, 0}},
	Key{"^", []byte{94, 0, 0}},
	Key{"&", []byte{38, 0, 0}},
	Key{"*", []byte{42, 0, 0}},
	Key{"(", []byte{40, 0, 0}},
	Key{")", []byte{41, 0, 0}},
	Key{"-", []byte{45, 0, 0}},
	Key{"_", []byte{95, 0, 0}},
	Key{"=", []byte{61, 0, 0}},
	Key{"+", []byte{43, 0, 0}},
	Key{"\\", []byte{92, 0, 0}},
	Key{"|", []byte{124, 0, 0}},
	Key{"{", []byte{123, 0, 0}},
	Key{"}", []byte{125, 0, 0}},
	Key{"[", []byte{91, 0, 0}},
	Key{"]", []byte{93, 0, 0}},
	Key{";", []byte{59, 0, 0}},
	Key{":", []byte{58, 0, 0}},
	Key{"\"", []byte{34, 0, 0}},
	Key{"'", []byte{39, 0, 0}},
	Key{"<", []byte{60, 0, 0}},
	Key{">", []byte{62, 0, 0}},
	Key{",", []byte{44, 0, 0}},
	Key{".", []byte{46, 0, 0}},
	Key{"/", []byte{47, 0, 0}},
	Key{"?", []byte{63, 0, 0}},

	// Uppercase letters
	Key{"A", []byte{65, 0, 0}},
	Key{"B", []byte{66, 0, 0}},
	Key{"C", []byte{67, 0, 0}},
	Key{"D", []byte{68, 0, 0}},
	Key{"E", []byte{69, 0, 0}},
	Key{"F", []byte{70, 0, 0}},
	Key{"G", []byte{71, 0, 0}},
	Key{"H", []byte{72, 0, 0}},
	Key{"I", []byte{73, 0, 0}},
	Key{"J", []byte{74, 0, 0}},
	Key{"K", []byte{75, 0, 0}},
	Key{"L", []byte{76, 0, 0}},
	Key{"M", []byte{77, 0, 0}},
	Key{"N", []byte{78, 0, 0}},
	Key{"O", []byte{79, 0, 0}},
	Key{"P", []byte{80, 0, 0}},
	Key{"Q", []byte{81, 0, 0}},
	Key{"R", []byte{82, 0, 0}},
	Key{"S", []byte{83, 0, 0}},
	Key{"T", []byte{84, 0, 0}},
	Key{"U", []byte{85, 0, 0}},
	Key{"V", []byte{86, 0, 0}},
	Key{"W", []byte{87, 0, 0}},
	Key{"X", []byte{88, 0, 0}},
	Key{"Y", []byte{89, 0, 0}},
	Key{"Z", []byte{90, 0, 0}},

	// Lowercase letters
	Key{"a", []byte{97, 0, 0}},
	Key{"b", []byte{98, 0, 0}},
	Key{"c", []byte{99, 0, 0}},
	Key{"d", []byte{100, 0, 0}},
	Key{"e", []byte{101, 0, 0}},
	Key{"f", []byte{102, 0, 0}},
	Key{"g", []byte{103, 0, 0}},
	Key{"h", []byte{104, 0, 0}},
	Key{"i", []byte{105, 0, 0}},
	Key{"j", []byte{106, 0, 0}},
	Key{"k", []byte{107, 0, 0}},
	Key{"l", []byte{108, 0, 0}},
	Key{"m", []byte{109, 0, 0}},
	Key{"n", []byte{110, 0, 0}},
	Key{"o", []byte{111, 0, 0}},
	Key{"p", []byte{112, 0, 0}},
	Key{"q", []byte{113, 0, 0}},
	Key{"r", []byte{114, 0, 0}},
	Key{"s", []byte{115, 0, 0}},
	Key{"t", []byte{116, 0, 0}},
	Key{"u", []byte{117, 0, 0}},
	Key{"v", []byte{118, 0, 0}},
	Key{"w", []byte{119, 0, 0}},
	Key{"x", []byte{120, 0, 0}},
	Key{"y", []byte{121, 0, 0}},
	Key{"z", []byte{122, 0, 0}},
}

// Return a key object from the raw key code
func GetKey(b []byte) Key {
	// Search the keymap
	for index := range keycodes {
		key := keycodes[index]
		if key.Code[0] == b[0] && key.Code[1] == b[1] && key.Code[2] == b[2] {
			return key
		}
	}

	// Unknown key
	newB := make([]byte, 3)
	copy(newB, b)
	key := Key{
		Name: "UNKNOWN",
		Code: newB,
	}
	return key
}
