package uinput

import "github.com/bendahl/uinput"

type Mouse struct {
	uinput.Mouse
}

// TODO: WE should be passing in the name atleast and possibly the location
func NewMouse() *Mouse {
	autopoint, err := uinput.CreateMouse("/dev/uinput", []byte("autopoint-0.1.0"))
	if err != nil {
		panic(err)
	}
	// TODO Again must be better way
	return &Mouse{Mouse: autopoint}
}

func (m *Mouse) Close() error { return m.Mouse.Close() }
