package main

import (
	"fmt"

	uinput "github.com/wade-welles/uinput"
)

func main() {
	fmt.Printf("uinput test\n")
	fmt.Printf("===========\n")

	vKeyboard := uinput.NewKeyboard()
	defer vKeyboard.Close()

	fmt.Printf("Alt+Tab Test\n")
	vKeyboard.AltTab()

	// TODO: Attempt alt + tab via hotkye

}
