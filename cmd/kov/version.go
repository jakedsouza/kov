///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/supervised-io/kov"
)

const (
	// DevBuild flag if build does not contain a valid version
	DevBuild = "dev"
)

func registerVersionCmd(cli *Cli) {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		Long:  "Show version information",
		RunE:  cli.runner(printVersion),
	}
	cli.rootCmd.AddCommand(versionCmd)
}

func printVersion(cli *Cli) error {
	if cli.Verbose {
		cli.printer.VerboseInfo(fmt.Sprintf("%s Version\n", CliProgram))
	}

	v := getVersion()

	cli.printer.Println("version:", v)
	cli.printer.Println("commit:", strings.Replace(kov.Commit, "'", "", -1))
	return nil
}

func getVersion() string {
	if kov.Version == "" {
		kov.Version = DevBuild
	}
	return strings.Replace(kov.Version, "'", "", -1)
}
