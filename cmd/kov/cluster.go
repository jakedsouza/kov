///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/spf13/cobra"
)

type clusterCmd struct {
	cluster          string // cluster
	configFile       string // path to a configuration file
	output           string // output format
	details          bool
	timestamps       bool
	isDefaultCluster bool
}

func registerClusterCmds(cli *Cli) {
	cli.clusterCmd = &clusterCmd{}

	clusterCmd := &cobra.Command{
		Use:     "kov",
		Short:   "KOV cluster related commands",
		Long:    "Commands to deploy, list and delete a cluster on Vsphere",
		Example: `kov create cluster kubernetes --config config.json`,
		RunE:    cli.usageRunner(),
	}

	cli.rootCmd.AddCommand(clusterCmd)
}

func registerVersionCmd(cli *Cli) {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		Long:  "Show version information",
		RunE:  cli.runner(printVersion),
	}
	cli.rootCmd.AddCommand(versionCmd)
}
