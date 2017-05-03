///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	"errors"

	"github.com/go-openapi/swag"
	"github.com/supervised-io/kov/gen/client/operations"
	"github.com/supervised-io/kov/gen/models"
	"github.com/supervised-io/kov/pkg/configfile_utils"
	"github.com/supervised-io/kov/pkg/kovclient_utils"
)

func createClusterPre(cli *Cli) error {
	if cli.clusterCmd.configFile == "" {
		return errors.New("Configuration path not provided")
	}
	if cli.clusterCmd.url == "" {
		return errors.New("KOV server endpoint not provided")
	}
	return nil
}

func createClusterRun(cli *Cli) error {
	err := createCluster(cli)

	return err
}

func createCluster(cli *Cli) error {
	cli.printer.VerboseInfo("Creating cluster")

	var clusterConfig *models.ClusterConfig
	if cli.clusterCmd.configFile != "" {
		bytes, err := configfileutils.ReadConfigFile(cli.clusterCmd.configFile)
		if err != nil {
			return err
		}
		config, err := configfileutils.ParseClusterCreateConfig(bytes)
		if err != nil {
			return err
		}
		clusterConfig = config
	} else {
		clusterConfig = &models.ClusterConfig{}
	}

	params := operations.NewCreateClusterParams().WithClusterConfig(clusterConfig)

	kovClient, err := kovclientutils.GetKOVClient(cli.clusterCmd.url)
	if err != nil {
		cli.printer.Error(fmt.Sprintf("Fatal Error Create Cluster : Create KOVClient error ErrorMsg %s", err))
		return err
	}

	// send request
	resp, err := kovClient.CreateCluster(params)

	if err != nil {
		var payload *models.Error
		switch etp := err.(type) {
		case *operations.CreateClusterConflict:
			payload = etp.Payload

			cli.printer.Fatal(fmt.Sprintf("Fatal Error Create Cluster Conflict: Code %d, ErrorMsg %s", payload.Code, *payload.Message))
		default:
			return fmt.Errorf("%s: %s \n", "Failed cluster create", err.Error())
		}

		if swag.StringValue(payload.Message) == "" {
			msg := err.Error()
			payload.Message = swag.String(msg)
		}
		return fmt.Errorf(err.Error(), swag.StringValue(payload.Message), swag.Int64Value(payload.Code))
	}

	// TODO support async
	cli.printer.Println(fmt.Sprintf("Creating cluster %s", *clusterConfig.Name))
	err = waitForTask(cli, string(resp.Payload))
	if err != nil {
		return fmt.Errorf(err.Error(), clusterConfig.Name, err.Error())
	}
	cli.printer.Println(fmt.Sprintf("Created cluster %s", *clusterConfig.Name))
	return nil
}
