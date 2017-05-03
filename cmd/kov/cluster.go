///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/spf13/cobra"
)

type clusterCmd struct {
	configFile       string // path to a configuration file
	output           string // output format
	details          bool
	timestamps       bool
	isDefaultCluster bool
	url              string // url of KOV server endpoint
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

	createClusterCmd := &cobra.Command{
		Use:     "kov",
		Short:   "KOV cluster related commands",
		Long:    "Commands to deploy, list and delete a cluster on Vsphere",
		Example: `kov create cluster kubernetes --url http://KOV_ENDPOINT --config config.json`,
		PreRunE: cli.preRunner(createClusterPre),
		RunE:    cli.runner(createCluster),
	}
	cli.rootCmd.AddCommand(createClusterCmd)

	createClusterCmd.Flags().StringVarP(&cli.clusterCmd.configFile, "config", "c", "", "Path to a JSON or YAML configuration file for the cell")
	createClusterCmd.Flags().StringVarP(&cli.clusterCmd.url, "url", "u", "", "url of KOV service")
}
