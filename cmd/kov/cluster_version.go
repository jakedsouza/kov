///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"strings"

	"github.com/supervised-io/kov"
	"github.com/supervised-io/kov/pkg/util/printer"
)

const (
	// DevBuild flag if build does not contain a valid version
	DevBuild = "dev"
)

func printVersion(cli *Cli) error {
	if cli.Verbose {
		printer.VerboseInfo(fmt.Sprintf("%s Version\n", CliProgram))
	}

	v := getVersion()

	printer.Println("version:", v)
	printer.Println("commit:", strings.Replace(kov.Commit, "'", "", -1))
	return nil
}

func getVersion() string {
	if kov.Version == "" {
		kov.Version = DevBuild
	}
	return strings.Replace(kov.Version, "'", "", -1)
}
