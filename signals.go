package uinput

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func NewSignalsHandler() {
	shutdownSignals := make(chan bool, 1)
	go signalsHandler(shutdownSignals)
}

func signalsHandler(shutdownSignals <-chan bool) {
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT, os.Kill)
	// TODO: Is this for needed? the select should hold everything open
	for {
		// TODO: Is there not a more compact way to write thise xact thing?
		s := <-signals
		switch s {
		case syscall.SIGTERM, syscall.SIGINT:
			fmt.Printf("user initiated shutdown (CTRL+C)...\n")
			os.Exit(0)
		case os.Kill:
			fmt.Printf("user killed the program...\n")
			os.Exit(0)
		default:
			fmt.Printf("how often is this being fucking ran?\n")
		}
	}
}
