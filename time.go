package uinput

import (
	"time"
)

// TODO: COuld and should be like uint8 or uint16 max but int is easier
func WaitMinutes(minutes int) { WaitSeconds(60 * minutes) }

func WaitSeconds(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}
