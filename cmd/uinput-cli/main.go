package main

import (
	"fmt"
)

func main() {
	fmt.Printf("uinput test\n")
	fmt.Printf("===========\n")

	fmt.Printf("Alt+Tab Test\n")

	vKeyboard := NewKeyboard()

	vKeyboard.AltTab()

	// TODO: Attempt alt + tab via hotkye

}
