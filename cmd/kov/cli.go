///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package main

import (
	"bufio"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/supervised-io/kov/pkg/cluster"
	"github.com/supervised-io/kov/pkg/printer"
)

// Cli struct for KOV Cli
type Cli struct {
	Cmd  *cobra.Command // the current command
	Args []string       // the arguments for current command

	Verbose bool
	Debug   bool

	stdIn      io.Reader
	input      *bufio.Reader
	rootCmd    *cobra.Command // the command hierarchy
	v          *viper.Viper
	clusterCmd *clusterCmd

	cluster cluster.ClusterAPI
	printer *printer.Printer
}

const (
	// CliProgram CLI program name
	CliProgram = "kov"
	// kov ENDPOINT ENV
	kovEndpoint = "ENDPOINT"
)

// NewCli configures a new CLI for KOV.
// Loads environment configuration and registers sub commands
func NewCli() *Cli {
	cli := &Cli{
		Verbose: false,
		Debug:   false,
		stdIn:   os.Stdin,
		v:       viper.New(),
	}
	cli.setDefaultConfig()
	cli.rootCmd = &cobra.Command{
		Use:   CliProgram,
		Short: "Kubernetes On vSphere CLI",
		Long: `Kubernetes On vSphere (KOV) is a command line interface for creating and managing clusters.
To get started, visit https://github.com/supervised-io/kov`,
		RunE: cli.usageRunner(),
	}

	cli.printer = printer.New(os.Stdout, os.Stderr)
	cli.SetOutput(os.Stdout, os.Stderr)

	cli.rootCmd.PersistentFlags().
		BoolVarP(&cli.Verbose, "verbose", "v", cli.Verbose, "Output more information")
	cli.rootCmd.PersistentFlags().
		BoolVar(&cli.Debug, "debug", cli.Debug, "Show debug related information, e.g. stack trace.")

	cli.v.BindPFlag("verbose", cli.rootCmd.PersistentFlags().Lookup("verbose"))
	cli.v.BindPFlag("debug", cli.rootCmd.PersistentFlags().Lookup("debug"))

	// set debug flag for swagger
	if cli.v.GetBool("debug") {
		os.Setenv("DEBUG", "1")
		cli.printer.SetDebug(true)
	}

	// Register sub-commands from here
	registerVersionCmd(cli)
	registerClusterCmds(cli)

	return cli
}

func (cli *Cli) setDefaultConfig() {
	cli.v.SetDefault("verbose", false)
	cli.v.SetDefault("debug", false)
	cli.v.SetDefault(kovEndpoint, "")
	cli.v.SetConfigName("config")
	cli.v.SetEnvPrefix("KOV")
	cli.v.AutomaticEnv()
	cli.v.ReadInConfig()
}

// SetOutput set output for err and out streams
func (cli *Cli) SetOutput(stdout, stderr io.Writer) *Cli {
	if stdout != nil {
		cli.rootCmd.SetOutput(stdout)
	}
	if stderr != nil {
		cli.rootCmd.SetOutput(stderr)
	}

	cli.printer.SetOutput(stdout, stderr)
	return cli
}

// Run runs the CLI
func (cli *Cli) Run() {
	cli.v.ReadInConfig()
	if cli.stdIn != nil {
		cli.input = bufio.NewReader(cli.stdIn)
	}
	if err := cli.rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// SetCluster set the cluster client of cli
func (cli *Cli) SetCluster(cluster cluster.ClusterAPI) {
	cli.cluster = cluster
}

// print usage of current command
func (cli *Cli) usage() error {
	if cli.Cmd != nil {
		return cli.Cmd.Usage()
	}
	return cli.rootCmd.Usage()
}

// runner execute current command with args
func (cli *Cli) runner(runner func(*Cli) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cli.rootCmd.SilenceUsage = true
		cli.Cmd = cmd
		cli.Args = args
		return runner(cli)
	}
}

// preRunner run a preRunner function for current command and args
func (cli *Cli) preRunner(preRunner func(*Cli) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cli.Cmd = cmd
		cli.Args = args
		return preRunner(cli)
	}
}

// postRunner runs a postRunner function for current command and args
func (cli *Cli) postRunner(postRunner func(*Cli) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cli.rootCmd.SilenceUsage = true
		cli.Cmd = cmd
		cli.Args = args
		return postRunner(cli)
	}
}

// usageRunner call usage on current command
func (cli *Cli) usageRunner() func(*cobra.Command, []string) error {
	return cli.runner(func(cli *Cli) error {
		return cli.usage()
	})
}
