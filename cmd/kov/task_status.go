///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/go-openapi/swag"
	"github.com/gosuri/uilive"
	"github.com/supervised-io/kov/gen/client/operations"
	"github.com/supervised-io/kov/gen/models"
	"github.com/supervised-io/kov/pkg/kovclient_utils"
	"github.com/supervised-io/kov/pkg/task_utils"
	spin "github.com/tj/go-spin"
)

// Can be a util method if needed
func waitForTask(cli *Cli, taskID string) error {
	var wg sync.WaitGroup
	wg.Add(1)

	writer := uilive.New()
	//writer.RefreshInterval = time.Nanosecond
	writer.Start()
	defer writer.Stop()

	msgChan := make(chan string, 10)
	spinChan := make(chan bool, 10)
	doneChan := make(chan bool, 10)
	msgChan <- "Processing Task"
	spinChan <- true
	go func() {
		defer wg.Done()
		printSpinner(writer, msgChan, spinChan, doneChan)
	}()
	return taskutils.PollWait(func() (bool, error) {
		params := operations.NewGetTaskParams().WithTaskid(taskID)

		kovClient, err := kovclientutils.GetKOVClient(cli.clusterCmd.url)
		if err != nil {
			// cli.printer.Error(fmt.Sprintf("Fatal Error Create Cluster : Create KOVClient error: Code %d, ErrorMsg %s", payload.Code, *payload.Message))
			return false, err
		}

		resp, err := kovClient.GetTask(params)
		if err != nil {
			var payload *models.Error
			switch etp := err.(type) {
			case *operations.GetTaskNotFound:
				payload = etp.Payload
				return false, nil
			case *operations.GetTaskDefault:
				payload = etp.Payload
				code := etp.Code()
				if code/100 != 2 {
					doneChan <- true
					wg.Wait()
					return false, errors.New("Error get task: GetTaskDefault")
				}
			default:
				msgChan <- fmt.Sprintf("Failed get task status for task:  %s\n", taskID)
				spinChan <- false
				doneChan <- false
				wg.Wait()
				return false, err
			}
			if swag.StringValue(payload.Message) == "" {
				msg := err.Error()
				payload.Message = swag.String(msg)
			}

			msgChan <- fmt.Sprintf("Failed get task status for task:  %s\n", taskID)
			spinChan <- false
			doneChan <- false
			wg.Wait()
			return false, fmt.Errorf("Error : %s Code: %d", swag.StringValue(payload.Message), swag.Int64Value(payload.Code))
		}
		// if get response without error
		switch resp.Payload.State {
		case models.TaskStateProcessing:
			{
				msgChan <- fmt.Sprintf("Processing Task(%s)", resp.Payload.Step)
				spinChan <- true
				return false, nil
			}
		case models.TaskStateFailed:
			{
				msgChan <- "Task failed.\n"
				spinChan <- false
				doneChan <- false
				wg.Wait()
				if resp.Payload.Context != nil {
					if resp.Payload.Context.Log == "" {
						return true, errors.New(resp.Payload.Context.Cause)
					}
					return true, errors.New(resp.Payload.Context.Log)
				}
				return true, fmt.Errorf(resp.Error())
			}
		case models.TaskStateCompleted:
			{
				doneChan <- true
				wg.Wait()
				return true, nil
			}
		}
		return false, nil
	})
}

/*
 * printSpinner writes and updates cli writer from msgChan channel and
 * prints a spinner depending upon values from spinChan channel. Upon
 * getting a true value from doneChan returns.
 */
func printSpinner(writer *uilive.Writer, msgChan chan string, spinChan, doneChan chan bool) {
	spinner := spin.New()
	spinner.Set(spin.Spin1)
	msg := ""
	useSpin := false
	shouldReturn := false
	for {
		if shouldReturn {
			fmt.Fprintf(writer.Bypass(), "%s", msg)
			writer.Flush()
			return
		}
		select {
		case done, ok := <-doneChan:
			if done && ok {
				fmt.Fprintf(writer.Bypass(), "")
				writer.Flush()
				return
			}
			select {
			case message := <-msgChan:
				msg = message
			default:
			}
			select {
			case usespin := <-spinChan:
				useSpin = usespin
			default:
			}
			shouldReturn = true
		default:
			select {
			case message := <-msgChan:
				msg = message
			default:
			}
			select {
			case usespin := <-spinChan:
				useSpin = usespin
			default:
			}
		}
		if msg != "" {
			if useSpin {
				fmt.Fprintf(writer, "%s %s\n", msg, spinner.Next())
				writer.Flush()
			} else {
				if shouldReturn {
					fmt.Fprintf(writer.Bypass(), "%s", msg)
					writer.Flush()
				} else {
					fmt.Fprintf(writer, "%s", msg)
					writer.Flush()
				}
			}
			time.Sleep(time.Millisecond * 100)
		}
	}
}
