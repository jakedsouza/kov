///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

//NO TESTS

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
	out   io.Writer
	debug bool
}

// New return Printer Obj
func New(stdOut, stderr io.Writer) *Printer {
	color.Output = stderr
	return &Printer{
		out:   stdOut,
		debug: false,
	}
}

// Println Print a message to Stdout
func (p *Printer) Println(args ...interface{}) {
	fmt.Fprintln(p.out, args...)
}

// Printf Print a message to Stdout
func (p *Printer) Printf(format string, args ...interface{}) {
	fmt.Fprintf(p.out, format, args...)
}

// Warn Print a warning to Stdout
func (p *Printer) Warn(format string, args ...interface{}) {
	color.Yellow(format, args...)
}

// Error Print an error message to Stdout
func (p *Printer) Error(format string, args ...interface{}) {
	color.Red(format, args...)
	if p.debug {
		debug.PrintStack()
	}
}

// Fatal Print a error message and and stack trace to Stdout and exit.
// Prints stack if debug flag is passed
func (p *Printer) Fatal(format string, args ...interface{}) {
	p.Error(format, args...)
	os.Exit(1)
}

// VerboseInfo print a VerboseInfo message
func (p *Printer) VerboseInfo(format string, args ...interface{}) {
	color.New(color.Faint).Printf(format, args...)
}

// SetOutput set output for printer
func (p *Printer) SetOutput(stdout, stderr io.Writer) {
	p.out = stdout
	color.Output = stderr
}

// SetDebug set debug bool
func (p *Printer) SetDebug(debug bool) {
	p.debug = debug
}
