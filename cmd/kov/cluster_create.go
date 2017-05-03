///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	"errors"

	"github.com/supervised-io/kov/pkg/configfile/reader"
)

func createClusterPre(cli *Cli) error {
	if cli.clusterCmd.configFile == "" {
		return errors.New("Configfile path not provided")
	}
	if cli.clusterCmd.url == "" {
		return errors.New("KOV endpoint not provided")
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
	taskID, err := cli.cluster.CreateCluster(cli.clusterCmd.url, clusterConfig)
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
