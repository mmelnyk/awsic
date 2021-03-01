package main

import (
	"errors"

	"github.com/spf13/cobra"
	"go.melnyk.org/awsic/internal/awsec2"
)

var ec2Cmd = &cobra.Command{
	Use:   "ec2",
	Short: "Work with EC2 resources",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var lsEC2Cmd = &cobra.Command{
	Use:   "ls",
	Short: "List EC2 instances",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		return awsec2.List()
	},
}

var inspectEC2Cmd = &cobra.Command{
	Use:   "inspect <Instance-ID>",
	Short: "Inspect EC2 instance",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Requires Instance ID or name")
		}
		return nil
	},
	Long: ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		return awsec2.Inspect(args[0])
	},
}

var startEC2Cmd = &cobra.Command{
	Use:   "start <Instance-ID>",
	Short: "Start EC2 instances",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Requires Instance ID or name")
		}
		return nil
	}, RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		return awsec2.Start(args[0])
	},
}

var stopEC2Cmd = &cobra.Command{
	Use:   "stop <Instance-ID>",
	Short: "Stop EC2 instances",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Requires Instance ID or name")
		}
		return nil
	}, RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		return awsec2.Stop(args[0])
	},
}

func init() {
	ec2Cmd.AddCommand(lsEC2Cmd)
	ec2Cmd.AddCommand(inspectEC2Cmd)
	ec2Cmd.AddCommand(startEC2Cmd)
	ec2Cmd.AddCommand(stopEC2Cmd)
	rootCmd.AddCommand(ec2Cmd)
}
