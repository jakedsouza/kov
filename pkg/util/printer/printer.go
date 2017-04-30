///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package printer

import (
	"fmt"
	"io"
	"runtime/debug"

	"os"

	"github.com/fatih/color"
)

// Printer for custermized printing
type Printer struct {
	out io.Writer
}

// New return Printer Obj
func New(stdOut, stderr io.Writer) *Printer {
	color.Output = stderr
	return &Printer{out: stdOut}
}

var std = New(os.Stdout, os.Stderr)

// Println Print a message to Stdout
func Println(args ...interface{}) {
	fmt.Fprintln(std.out, args...)
}

// Printf Print a message to Stdout
func Printf(format string, args ...interface{}) {
	fmt.Fprintf(std.out, format, args...)
}

// Warn Print a warning to Stdout
func Warn(format string, args ...interface{}) {
	color.Yellow(format, args...)
}

// Error Print an error message to Stdout
func Error(isDebug bool, format string, args ...interface{}) {
	color.Red(format, args...)
	if isDebug {
		debug.PrintStack()
	}
}

// Fatal Print a error message and and stack trace to Stdout and exit.
// Prints stack if debug flag is passed
func Fatal(isDebug bool, format string, args ...interface{}) {
	Error(isDebug, format, args...)
	os.Exit(1)
}

// Info print an info message
func Info(format string, args ...interface{}) {
	Printf(format, args...)
}

// VerboseInfo print a VerboseInfo message
func VerboseInfo(format string, args ...interface{}) {
	color.New(color.Faint).Printf(format, args...)
}
