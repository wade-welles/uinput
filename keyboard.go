package uinput

import "github.com/bendahl/uinput"

type Keyboard struct {
	uinput.Keyboard
}

func NewKeyboard() *Keyboard {
	// TODO: Should handle permissions issues in the software even if its hackey
	// as mycurrent solution.
	autokey, err := uinput.CreateKeyboard("/dev/uinput", []byte("autokey-0.1.0"))
	if err != nil {
		panic(err)
	}
	// TODO: Almost certain there is a better way than this example code
	return &Keyboard{Keyboard: autokey}
}

func (kb *Keyboard) Close() error { return kb.Keyboard.Close() }

func (kb *Keyboard) AltTab() *Keyboard {
	kb.KeyDown(uinput.KeyLeftalt)
	WaitSeconds(1)
	kb.KeyPress(uinput.KeyTab)
	WaitSeconds(1)
	kb.KeyUp(uinput.KeyLeftalt)
	return kb
}
