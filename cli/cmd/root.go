/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "iac-ttl",
	Short: "A CLI tool to manage TTL (Time To Live) for CloudFormation stacks",
	Long: `IaC-TTL is a command-line tool that helps manage the lifecycle of CloudFormation stacks
by automatically scheduling their deletion after a specified time period.

This tool uses AWS EventBridge Scheduler to create and manage deletion schedules for
CloudFormation stacks, making it easy to implement TTL policies for your infrastructure.

Available Commands:
  add     Schedule a stack for deletion after a specified time
  ls      List all scheduled stack deletions
  rm      Remove a scheduled stack deletion

Examples:
  # Schedule a stack for deletion after 24 hours
  iac-ttl add my-stack -H 24

  # List all scheduled deletions
  iac-ttl ls

  # Remove a scheduled deletion
  iac-ttl rm my-stack`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
}
