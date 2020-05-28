package util

import "log"

var mode = true

// Debug generates debug message with text s
func Debug(s string) {
	if mode {
		log.Printf("[DBG] %s\n", s)
	}
}
