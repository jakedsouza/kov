///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"strings"

	"errors"

	"github.com/supervised-io/kov/pkg/cluster"
	"github.com/supervised-io/kov/pkg/configfile/reader"
)

func createClusterPre(cli *Cli) error {
	if cli.clusterCmd.configFile == "" {
		return errors.New("Configfile path not provided")
	}
	if cli.clusterCmd.url == "" {
		if cli.v.Get(kovEndpoint) == "" {
			return errors.New("KOV endpoint not provided")
		}
		cli.clusterCmd.url = cli.v.Get(kovEndpoint).(string)
	}
	var err error
	cli.cluster, err = cluster.NewClusterClient(strings.TrimPrefix(strings.TrimPrefix(cli.clusterCmd.url, "https://"), "http://"))
	if err != nil {
		return err
	}

	return nil
}

func createClusterRun(cli *Cli) error {
	// read config from config file
	clusterConfig, err := reader.ParseClusterCreateConfig(cli.clusterCmd.configFile)
	if err != nil {
		cli.printer.Error(fmt.Sprintf("Creating cluster Error: error parsing config file"))
		return err
	}
	// send create cluster request
	cli.printer.Println(fmt.Sprintf("Creating cluster %s", *clusterConfig.Name))
	taskID, err := cli.cluster.CreateCluster(clusterConfig)
	if err != nil {
		cli.printer.Error(fmt.Sprintf("Creating cluster Error: %s", err.Error()))
		return err
	}

	// wait for the task to complete
	cli.printer.Println(fmt.Sprintf("Starting task %s", *taskID))
	err = waitForTask(cli, *taskID)
	if err != nil {
		cli.printer.Error(fmt.Sprintf("Creating cluster Error: %s", err.Error()))
		return err
	}
	cli.printer.Println(fmt.Sprintf("Task %s Completed", *taskID))
	cli.printer.Println(fmt.Sprintf("Created cluster %s", *clusterConfig.Name))

	return nil
}
