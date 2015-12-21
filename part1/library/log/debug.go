package log

import (
	wrapper "github.com/evalphobia/go-log-wrapper/log"
)

// Nothing is dummy variable for import error
var Nothing int

// Dump prints dump variable in console
func Dump(v interface{}) {
	wrapper.Dump(v)
}

// Print prints variable information in console
func Print(v interface{}) {
	wrapper.Print(v)
}

// Header prints separator in console
func Header(v ...interface{}) {
	wrapper.Header(v...)
}

// Mark prints trace info
func Mark() {
	wrapper.Mark(3)
}
