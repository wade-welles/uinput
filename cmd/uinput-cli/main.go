package main

import (
	"fmt"
	"time"

	uinput "github.com/bendahl/uinput"
)

type Keyboard struct {
	uinput.Keyboard
}

type Mouse struct {
	uinput.Mouse
}

// TODO: WE should be passing in the name atleast and possibly the location
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

func (kb *Keyboard) AltTab() *Keyboard {
	kb.KeyDown(uinput.KeyLeftalt)
	WaitSeconds(1)
	kb.KeyPress(uinput.KeyTab)
	WaitSeconds(1)
	kb.KeyUp(uinput.KeyLeftalt)
	return kb
}

func (kb *Keyboard) Close() error { return kb.Keyboard.Close() }

func NewMouse() *Mouse {
	autopoint, err := uinput.CreateMouse("/dev/uinput", []byte("autopoint-0.1.0"))
	if err != nil {
		panic(err)
	}
	// TODO Again must be better way
	return &Mouse{Mouse: autopoint}
}

func (m *Mouse) Close() error { return m.Mouse.Close() }

func WaitMinutes(minutes int) { WaitSeconds(60 * minutes) }

func WaitSeconds(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Printf("uinput test\n")
	fmt.Printf("===========\n")

	fmt.Printf("Alt+Tab Test\n")

	vKeyboard := NewKeyboard()

	vKeyboard.AltTab()

	// TODO: Attempt alt + tab via hotkye

}
