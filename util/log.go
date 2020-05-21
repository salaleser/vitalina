package util

import "log"

var DebugMode = true

func Debug(s string) {
	if DebugMode {
		log.Printf("[DBG] %s\n", s)
	}
}
