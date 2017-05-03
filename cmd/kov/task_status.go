///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package main

import taskpoller "github.com/supervised-io/kov/pkg/poller/status/task"

// loop for getting task status
func waitForTask(cli *Cli, taskID string) error {
	return taskpoller.PollWait(func() (bool, error) {
		return cli.cluster.GetTaskStatus(cli.clusterCmd.url, taskID)
	})
}
