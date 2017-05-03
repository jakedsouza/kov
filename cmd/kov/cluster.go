///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/spf13/cobra"
)

type clusterCmd struct {
	configFile string // path to a configuration file
	url        string // url of KOV endpoint
}

func registerClusterCmds(cli *Cli) {
	cli.clusterCmd = &clusterCmd{}

	clusterCmd := &cobra.Command{
		Use:   "cluster",
		Short: "KOV cluster related commands",
		Long:  "Commands to deploy, list and delete clusters on vSphere",
		RunE:  cli.usageRunner(),
	}

	createClusterCmd := &cobra.Command{
		Use:     "create",
		Short:   "create cluster",
		Long:    "Command to create cluster on vSphere",
		Example: `kov cluster create --url https://KOV_ENDPOINT --config config.json`,
		PreRunE: cli.preRunner(createClusterPre),
		RunE:    cli.runner(createClusterRun),
	}

	clusterCmd.AddCommand(createClusterCmd)
	cli.rootCmd.AddCommand(clusterCmd)

	createClusterCmd.Flags().StringVarP(&cli.clusterCmd.configFile, "config", "c", "", "Path to a JSON or YAML configuration file for the cluster")
	createClusterCmd.Flags().StringVarP(&cli.clusterCmd.url, "url", "u", "", "url of KOV endpoint")
}
